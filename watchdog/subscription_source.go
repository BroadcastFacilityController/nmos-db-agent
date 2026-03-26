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

type SubscriptionSource struct {
	ws      *Websocket
	version common.APIVersion
	done    chan bool
}

func NewSubscriptionSource(ws *Websocket, version common.APIVersion) (*SubscriptionSource, error) {
	sub := SubscriptionSource{
		ws:      ws,
		done:    make(chan bool),
		version: version,
	}
	return &sub, nil
}

func (sub *SubscriptionSource) Stop() {
	sub.done <- true
}

func (sub *SubscriptionSource) Start() {
	go func() {
		sub.readWebsocket()
	}()
}

func (sub *SubscriptionSource) readWebsocket() {
	defer close(sub.done)
	for {
		_, msgByte, err := sub.ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Source Websocket: %v", err)
		}
		if sub.version.Equals(common.NewAPIVersion(1, 0)) {
			var msg is04v1_0.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Source Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_0.QUERY_WS_SOURCES {
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
					logrus.Errorf("Source Websocket: %v", err)
				}
				var parsed is04v1_0.Source
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
				err = sub.handleSourceV1_0(parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 1)) {
			var msg is04v1_1.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Source Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_1.QUERY_WS_SOURCES {
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
					logrus.Errorf("Source Websocket: %v", err)
				}
				var parsed is04v1_1.Source
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
				err = sub.handleSourceV1_1(parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 2)) {
			var msg is04v1_2.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Source Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_2.QUERY_WS_SOURCES {
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
					logrus.Errorf("Source Websocket: %v", err)
				}
				var parsed is04v1_2.Source
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
				err = sub.handleSourceV1_2(parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 3)) {
			var msg is04v1_3.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Source Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_3.QUERY_WS_SOURCES {
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
					logrus.Errorf("Source Websocket: %v", err)
				}
				var parsed is04v1_3.Source
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
				err = sub.handleSourceV1_3(parsed)
				if err != nil {
					logrus.Errorf("Source Websocket: %v", err)
				}
			}
			continue
		}
	}
}

func (sub *SubscriptionSource) handleSourceV1_0(source is04v1_0.Source) error {
	line, err := schema.NewSourceFromV1_0(&source)
	if err != nil {
		return err
	}
	// Check if node
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
	logrus.Debugf("Source Websocket V1.0: Updating Source %s", source.ID)
	err = database.UpsertSource(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionSource) handleSourceV1_1(source is04v1_1.Source) error {
	line, err := schema.NewSourceFromV1_1(&source)
	if err != nil {
		return err
	}
	// Check if node
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
	logrus.Debugf("Source Websocket V1.1: Updating Source %s", line.ID)
	err = database.UpsertSource(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionSource) handleSourceV1_2(source is04v1_2.Source) error {
	line, err := schema.NewSourceFromV1_2(&source)
	if err != nil {
		return err
	}
	// Check if node
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
	logrus.Debugf("Source Websocket V1.2: Updating Source %s", line.ID)
	err = database.UpsertSource(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionSource) handleSourceV1_3(source is04v1_3.Source) error {
	line, err := schema.NewSourceFromV1_3(&source)
	if err != nil {
		return err
	}
	// Check if node
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
	logrus.Debugf("Source Websocket V1.3: Updating Source %s", line.ID)
	err = database.UpsertSource(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}
