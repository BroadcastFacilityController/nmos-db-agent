package watchdog

import (
	"context"
	"encoding/json"
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

type SubscriptionReceiver struct {
	ws      *Websocket
	version common.APIVersion
	done    chan bool
}

func NewSubscriptionReceiver(ws *Websocket, version common.APIVersion) (*SubscriptionReceiver, error) {
	sub := SubscriptionReceiver{
		ws:      ws,
		done:    make(chan bool),
		version: version,
	}
	return &sub, nil
}

func (sub *SubscriptionReceiver) Stop() {
	sub.done <- true
}

func (sub *SubscriptionReceiver) Start() {
	go func() {
		sub.readWebsocket()
	}()
}

func (sub *SubscriptionReceiver) readWebsocket() {
	defer close(sub.done)
	for {
		_, msgByte, err := sub.ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Receiver Websocket: %v", err)
		}
		if sub.version.Equals(common.NewAPIVersion(1, 0)) {
			var msg is04v1_0.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Receiver Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_0.QUERY_WS_RECEIVERS {
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
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				var parsed is04v1_0.Receiver
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				go func(parsed *is04v1_0.Receiver) {
					err = sub.handleReceiverV1_0(*parsed)
					if err != nil {
						logrus.Errorf("Receiver Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 1)) {
			var msg is04v1_1.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Receiver Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_1.QUERY_WS_RECEIVERS {
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
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				var parsed is04v1_1.Receiver
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				go func(parsed *is04v1_1.Receiver) {
					err = sub.handleReceiverV1_1(*parsed)
					if err != nil {
						logrus.Errorf("Receiver Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 2)) {
			var msg is04v1_2.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Receiver Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_2.QUERY_WS_RECEIVERS {
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
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				var parsed is04v1_2.Receiver
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				go func(parsed *is04v1_2.Receiver) {
					err = sub.handleReceiverV1_2(*parsed)
					if err != nil {
						logrus.Errorf("Receiver Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 3)) {
			var msg is04v1_3.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Receiver Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_3.QUERY_WS_RECEIVERS {
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
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				var parsed is04v1_3.Receiver
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Receiver Websocket: %v", err)
				}
				go func(parsed *is04v1_3.Receiver) {
					err = sub.handleReceiverV1_3(*parsed)
					if err != nil {
						logrus.Errorf("Receiver Websocket: %v", err)
					}
				}(&parsed)
			}
			continue
		}
	}
}

func (sub *SubscriptionReceiver) handleReceiverV1_0(receiver is04v1_0.Receiver) error {
	line, err := schema.NewReceiverFromV1_0(&receiver)
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
	logrus.Debugf("Receiver Websocket V1.0: Updating Receiver %s", line.ID)
	err = database.UpsertReceiver(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionReceiver) handleReceiverV1_1(receiver is04v1_1.Receiver) error {
	line, err := schema.NewReceiverFromV1_1(&receiver)
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
	logrus.Debugf("Receiver Websocket V1.1: Updating Receiver %s", line.ID)
	err = database.UpsertReceiver(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionReceiver) handleReceiverV1_2(receiver is04v1_2.Receiver) error {
	line, err := schema.NewReceiverFromV1_2(&receiver)
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
	logrus.Debugf("Receiver Websocket V1.2: Updating Receiver %s", line.ID)
	err = database.UpsertReceiver(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionReceiver) handleReceiverV1_3(receiver is04v1_3.Receiver) error {
	line, err := schema.NewReceiverFromV1_3(&receiver)
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
	logrus.Debugf("Receiver Websocket V1.3: Updating Receiver %s", line.ID)
	err = database.UpsertReceiver(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}
