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

type SubscriptionNode struct {
	ws      *Websocket
	version common.APIVersion
	done    chan bool
}

func NewSubscriptionNode(ws *Websocket, version common.APIVersion) (*SubscriptionNode, error) {
	sub := SubscriptionNode{
		ws:      ws,
		done:    make(chan bool),
		version: version,
	}
	return &sub, nil
}

func (sub *SubscriptionNode) Stop() {
	sub.done <- true
}

func (sub *SubscriptionNode) Start() {
	go func() {
		sub.readWebsocket()
	}()
}

func (sub *SubscriptionNode) readWebsocket() {
	defer close(sub.done)
	for {
		_, msgByte, err := sub.ws.ReadMessage()
		if err != nil {
			logrus.Errorf("Node Websocket: %v", err)
		}
		if sub.version.Equals(common.NewAPIVersion(1, 0)) {
			var msg is04v1_0.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Node Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_0.QUERY_WS_NODES {
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
					logrus.Errorf("Node Websocket: %v", err)
				}
				var parsed is04v1_0.Node
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
				err = sub.handleNodeV1_0(parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 1)) {
			var msg is04v1_1.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Node Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_1.QUERY_WS_NODES {
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
					logrus.Errorf("Node Websocket: %v", err)
				}
				var parsed is04v1_1.Node
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
				err = sub.handleNodeV1_1(parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 2)) {
			var msg is04v1_2.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Node Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_2.QUERY_WS_NODES {
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
					logrus.Errorf("Node Websocket: %v", err)
				}
				var parsed is04v1_2.Node
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
				err = sub.handleNodeV1_2(parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
			}
			continue
		}
		if sub.version.Equals(common.NewAPIVersion(1, 3)) {
			var msg is04v1_3.QueryAPISubscriptionWSGrain
			err = json.Unmarshal(msgByte, &msg)
			if err != nil {
				logrus.Errorf("Node Websocket: %v", err)
			}
			if msg.Grain.Topic != is04v1_3.QUERY_WS_NODES {
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
					logrus.Errorf("Node Websocket: %v", err)
				}
				var parsed is04v1_3.Node
				err = json.Unmarshal(dataPost, &parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
				err = sub.handleNodeV1_3(parsed)
				if err != nil {
					logrus.Errorf("Node Websocket: %v", err)
				}
			}
			continue
		}
	}
}

func (sub *SubscriptionNode) handleNodeV1_0(node is04v1_0.Node) error {
	line, err := schema.NewNodeFromV1_0(node)
	if err != nil {
		return err
	}
	// Check if node exists
	existingNode, err := database.SelectNodeByID(context.Background(), line.ID)
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
	logrus.Debugf("Node Websocket V1.0: Updating Node %s", node.ID)
	err = database.UpsertNode(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionNode) handleNodeV1_1(node is04v1_1.Node) error {
	line, err := schema.NewNodeFromV1_1(node)
	if err != nil {
		return err
	}
	// Check if node exists
	existingNode, err := database.SelectNodeByID(context.Background(), line.ID)
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
	logrus.Debugf("Node Websocket V1.1: Updating Node %s", node.ID)
	err = database.UpsertNode(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionNode) handleNodeV1_2(node is04v1_2.Node) error {
	line, err := schema.NewNodeFromV1_2(node)
	if err != nil {
		return err
	}
	// Check if node exists
	existingNode, err := database.SelectNodeByID(context.Background(), line.ID)
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
	logrus.Debugf("Node Websocket V1.2: Updating Node %s", node.ID)
	err = database.UpsertNode(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}

func (sub *SubscriptionNode) handleNodeV1_3(node is04v1_3.Node) error {
	line, err := schema.NewNodeFromV1_3(node)
	if err != nil {
		return err
	}
	// Check if node exists
	existingNode, err := database.SelectNodeByID(context.Background(), line.ID)
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
	logrus.Debugf("Node Websocket V1.3: Updating Node %s", node.ID)
	err = database.UpsertNode(context.Background(), line)
	if err != nil {
		return err
	}
	return nil
}
