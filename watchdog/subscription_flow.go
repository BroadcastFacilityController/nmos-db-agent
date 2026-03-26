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

type SubscriptionFlow struct {
	ws      *Websocket
	version common.APIVersion
	done    chan bool
}

func NewSubscriptionFlow(ws *Websocket, version common.APIVersion) (*SubscriptionFlow, error) {
	sub := SubscriptionFlow{
		ws:      ws,
		done:    make(chan bool),
		version: version,
	}
	return &sub, nil
}

func (sub *SubscriptionFlow) Stop() {
	sub.done <- true
}

func (sub *SubscriptionFlow) Start() {
	go func() {
		sub.readWebsocket()
	}()
}

func (sub *SubscriptionFlow) readWebsocket() {
	defer close(sub.done)
	for {
		_, msgByte, err := sub.ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Flow Websocket: %v", err)
		}
		if sub.version.Equals(common.NewAPIVersion(1, 0)) {
			var msg is04v1_0.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Flow Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_0.QUERY_WS_FLOWS {
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
					logrus.Errorf("Flow Websocket: %v", err)
				}
				var parsed is04v1_0.Flow
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
				err = sub.handleFlowV1_0(parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 1)) {
			var msg is04v1_1.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Flow Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_1.QUERY_WS_FLOWS {
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
					logrus.Errorf("Flow Websocket: %v", err)
				}
				var parsed is04v1_1.Flow
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
				err = sub.handleFlowV1_1(parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 2)) {
			var msg is04v1_2.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Flow Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_2.QUERY_WS_FLOWS {
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
					logrus.Errorf("Flow Websocket: %v", err)
				}
				var parsed is04v1_2.Flow
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
				err = sub.handleFlowV1_2(parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 3)) {
			var msg is04v1_3.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Flow Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_3.QUERY_WS_FLOWS {
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
					logrus.Errorf("Flow Websocket: %v", err)
				}
				var parsed is04v1_3.Flow
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
				err = sub.handleFlowV1_3(parsed)
				if err != nil {
					logrus.Errorf("Flow Websocket: %v", err)
				}
			}
			continue
		}
	}
}

func (sub *SubscriptionFlow) handleFlowV1_0(flow is04v1_0.Flow) error {
	line, err := schema.NewFlowFromV1_0(&flow)
	if err != nil {
		return err
	}
	// Check if node's version is max 1.0, or if this is a downgraded copy
	parentSource, err := database.SelectSourceByID(context.Background(), *line.SourceID)
	if err != nil {
		// If source doesn't exist, don't add the flow yet as the foreign key will error anyways
		return nil
	}
	parentDevice, err := database.SelectDeviceByID(context.Background(), parentSource.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the flow yet as the foreign key will error anyways
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
	logrus.Debugf("Flow Websocket V1.0: Updating Flow %s", flow.ID)
	err = database.UpsertFlow(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionFlow) handleFlowV1_1(flow is04v1_1.Flow) error {
	line, err := schema.NewFlowFromV1_1(&flow)
	if err != nil {
		return err
	}
	// Check if node's version is max 1.1, or if this is a downgraded copy
	parentDevice, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the flow yet as the foreign key will error anyways
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
	logrus.Debugf("Flow Websocket V1.1: Updating Flow %s", line.ID)
	err = database.UpsertFlow(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionFlow) handleFlowV1_2(flow is04v1_2.Flow) error {
	line, err := schema.NewFlowFromV1_2(&flow)
	if err != nil {
		return err
	}
	// Check if node's version is max 1.2, or if this is a downgraded copy
	parentDevice, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the flow yet as the foreign key will error anyways
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
	logrus.Debugf("Flow Websocket V1.2: Updating Flow %s", line.ID)
	err = database.UpsertFlow(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionFlow) handleFlowV1_3(flow is04v1_3.Flow) error {
	line, err := schema.NewFlowFromV1_3(&flow)
	if err != nil {
		return err
	}
	// Check if node's version is max 1.3, or if this is a downgraded copy
	parentDevice, err := database.SelectDeviceByID(context.Background(), *line.DeviceID)
	if err != nil {
		// If device doesn't exist, don't add the flow yet as the foreign key will error anyways
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
	logrus.Debugf("Flow Websocket V1.3: Updating Flow %s", line.ID)
	err = database.UpsertFlow(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}
