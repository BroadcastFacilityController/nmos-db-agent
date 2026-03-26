package watchdog

import (
	"context"
	"encoding/json"
	"io"
	"net/http"
	"slices"
	"strings"

	"github.com/BroadcastFacilityController/nmos-db-agent/database"
	"github.com/BroadcastFacilityController/nmos-db-agent/database/schema"
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/sirupsen/logrus"
)

type SubscriptionSender struct {
	ws      *Websocket
	version common.APIVersion
	done    chan bool
}

func NewSubscriptionSender(ws *Websocket, version common.APIVersion) (*SubscriptionSender, error) {
	sub := SubscriptionSender{
		ws:      ws,
		done:    make(chan bool),
		version: version,
	}
	return &sub, nil
}

func (sub *SubscriptionSender) Stop() {
	sub.done <- true
}

func (sub *SubscriptionSender) Start() {
	go func() {
		sub.readWebsocket()
	}()
}

func (sub *SubscriptionSender) readWebsocket() {
	defer close(sub.done)
	for {
		_, msgByte, err := sub.ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Sender Websocket: %v", err)
		}
		if sub.version.Equals(common.NewAPIVersion(1, 0)) {
			var msg is04v1_0.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Sender Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_0.QUERY_WS_SENDERS {
				// message is being sent down wrong websocket, ignore it.
				continue
			}
			// Parse data
			for _, data := range msg.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				dataPost, err := json.Marshal(data.Post)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				var parsed is04v1_0.Sender
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				go func(parsed *is04v1_0.Sender) {
					err = sub.handleSenderV1_0(*parsed)
					if err != nil {
						logrus.Errorf("Sender Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 1)) {
			var msg is04v1_1.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Sender Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_1.QUERY_WS_SENDERS {
				// message is being sent down wrong websocket, ignore it.
				continue
			}
			// Parse data
			for _, data := range msg.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				dataPost, err := json.Marshal(data.Post)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				var parsed is04v1_1.Sender
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				go func(parsed *is04v1_1.Sender) {
					err = sub.handleSenderV1_1(*parsed)
					if err != nil {
						logrus.Errorf("Sender Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 2)) {
			var msg is04v1_2.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Sender Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_2.QUERY_WS_SENDERS {
				// message is being sent down wrong websocket, ignore it.
				continue
			}
			// Parse data
			for _, data := range msg.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				dataPost, err := json.Marshal(data.Post)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				var parsed is04v1_2.Sender
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				go func(parsed *is04v1_2.Sender) {
					err = sub.handleSenderV1_2(*parsed)
					if err != nil {
						logrus.Errorf("Sender Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 3)) {
			var msg is04v1_3.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Sender Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_3.QUERY_WS_SENDERS {
				// message is being sent down wrong websocket, ignore it.
				continue
			}
			// Parse data
			for _, data := range msg.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				dataPost, err := json.Marshal(data.Post)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				var parsed is04v1_3.Sender
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Sender Websocket: %v", err)
				}
				go func(parsed *is04v1_3.Sender) {
					err = sub.handleSenderV1_3(*parsed)
					if err != nil {
						logrus.Errorf("Sender Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
	}
}

func (sub *SubscriptionSender) handleSenderV1_0(sender is04v1_0.Sender) error {
	// Fetch transport file
	var transportFile []byte
	if sender.ManifestHRef != "" {
		req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Accept", "application/sdp,application/json,application/xml")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = make([]byte, len(respBody))
		copy(transportFile, respBody)
		if resp.StatusCode >= 400 {
			transportFile = nil
		}
	} else {
		transportFile = nil
	}
	line, err := schema.NewSenderFromV1_0(&sender, transportFile)
	if err != nil {
		return err
	}
	// Check if node's max version is v1.0
	parentDevice, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the source yet as the foreign key will error anyways
		return nil
	}
	existingNode, err := database.SelectNodeByID(context.Background(), parentDevice.NodeID)
	if existingNode != nil {
		// Check if this is the max version
		existingVersions := make([]string, len(existingNode.APIVersions))
		copy(existingVersions, existingNode.SupportedVersions)
		slices.SortFunc(existingVersions, func(a string, b string) int {
			return strings.Compare(a, b)
		})
		if existingVersions[len(existingVersions)-1] != "v1.0" {
			return nil
		}
	}
	// Do insert / update
	logrus.Debugf("Sender Websocket V1.0: Updating Sender %s", line.ID)
	err = database.UpsertSender(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionSender) handleSenderV1_1(sender is04v1_1.Sender) error {
	// Fetch transport file
	var transportFile []byte
	if sender.ManifestHRef != "" {
		req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Accept", "application/sdp,application/json,application/xml")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = make([]byte, len(respBody))
		copy(transportFile, respBody)
		if resp.StatusCode >= 400 {
			transportFile = nil
		}
	} else {
		transportFile = nil
	}
	line, err := schema.NewSenderFromV1_1(&sender, transportFile)
	if err != nil {
		return err
	}
	// Check if node's max version is v1.1
	parentDevice, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the source yet as the foreign key will error anyways
		return nil
	}
	existingNode, err := database.SelectNodeByID(context.Background(), parentDevice.NodeID)
	if existingNode != nil {
		// Check if this is the max version
		existingVersions := make([]string, len(existingNode.APIVersions))
		copy(existingVersions, existingNode.SupportedVersions)
		slices.SortFunc(existingVersions, func(a string, b string) int {
			return strings.Compare(a, b)
		})
		if existingVersions[len(existingVersions)-1] != "v1.1" {
			return nil
		}
	}
	// Do insert / update
	logrus.Debugf("Sender Websocket V1.1: Updating Sender %s", line.ID)
	err = database.UpsertSender(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionSender) handleSenderV1_2(sender is04v1_2.Sender) error {
	// Fetch transport file
	var transportFile []byte
	if sender.ManifestHRef != "" {
		req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef, nil)
		if err != nil {
			return err
		}
		req.Header.Add("Accept", "application/sdp,application/json,application/xml")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = make([]byte, len(respBody))
		copy(transportFile, respBody)
		if resp.StatusCode >= 400 {
			transportFile = nil
		}
	} else {
		transportFile = nil
	}
	line, err := schema.NewSenderFromV1_2(&sender, transportFile)
	if err != nil {
		return err
	}
	// Check if node's max version is v1.2
	parentDevice, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the source yet as the foreign key will error anyways
		return nil
	}
	existingNode, err := database.SelectNodeByID(context.Background(), parentDevice.NodeID)
	if existingNode != nil {
		// Check if this is the max version
		existingVersions := make([]string, len(existingNode.APIVersions))
		copy(existingVersions, existingNode.SupportedVersions)
		slices.SortFunc(existingVersions, func(a string, b string) int {
			return strings.Compare(a, b)
		})
		if existingVersions[len(existingVersions)-1] != "v1.2" {
			return nil
		}
	}
	// Do insert / update
	logrus.Debugf("Sender Websocket V1.2: Updating Sender %s", line.ID)
	err = database.UpsertSender(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionSender) handleSenderV1_3(sender is04v1_3.Sender) error {
	// Fetch transport file
	var transportFile []byte
	if sender.ManifestHRef.Valid && sender.ManifestHRef.ValueOrZero() != "" {
		req, err := http.NewRequest(http.MethodGet, sender.ManifestHRef.ValueOrZero(), nil)
		if err != nil {
			return err
		}
		req.Header.Add("Accept", "application/sdp,application/json,application/xml")
		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		transportFile = make([]byte, len(respBody))
		copy(transportFile, respBody)
		if resp.StatusCode >= 400 {
			transportFile = nil
		}
	} else {
		transportFile = nil
	}
	line, err := schema.NewSenderFromV1_3(&sender, transportFile)
	if err != nil {
		return err
	}
	// Check if node's max version is v1.3
	parentDevice, err := database.SelectDeviceByID(context.Background(), line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the source yet as the foreign key will error anyways
		return nil
	}
	existingNode, err := database.SelectNodeByID(context.Background(), parentDevice.NodeID)
	if existingNode != nil {
		// Check if this is the max version
		existingVersions := make([]string, len(existingNode.APIVersions))
		copy(existingVersions, existingNode.SupportedVersions)
		slices.SortFunc(existingVersions, func(a string, b string) int {
			return strings.Compare(a, b)
		})
		if existingVersions[len(existingVersions)-1] != "v1.3" {
			return nil
		}
	}
	// Do insert / update
	logrus.Debugf("Sender Websocket V1.3: Updating Sender %s", line.ID)
	err = database.UpsertSender(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}
