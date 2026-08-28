package watchdog

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/watchdog/stablews"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (w *Watchdog) startV1_1(secure bool) error {
	w.sockets["v1.1"] = make(map[string]*stablews.Client)
	var err error
	log := logrus.WithFields(logrus.Fields{
		"location": "watchdog",
		"version":  "v1.1",
	})
	log.WithField("path", "nodes").Debug("starting watchdog")
	err = w.startV1_1Nodes(secure)
	if err != nil {
		return fmt.Errorf("v1.1 nodes error: %w", err)
	}
	w.waitForDataPause("v1.1", "/nodes", 150*time.Millisecond)
	log.WithField("path", "devices").Debug("starting watchdog")
	err = w.startV1_1Devices(secure)
	if err != nil {
		return fmt.Errorf("v1.1 devices error: %w", err)
	}
	w.waitForDataPause("v1.1", "/devices", 150*time.Millisecond)
	log.WithField("path", "sources").Debug("starting watchdog")
	err = w.startV1_1Sources(secure)
	if err != nil {
		return fmt.Errorf("v1.1 sources error: %w", err)
	}
	w.waitForDataPause("v1.1", "/sources", 150*time.Millisecond)
	log.WithField("path", "flows").Debug("starting watchdog")
	err = w.startV1_1Flows(secure)
	if err != nil {
		return fmt.Errorf("v1.1 flows error: %w", err)
	}
	w.waitForDataPause("v1.1", "/flows", 150*time.Millisecond)
	log.WithField("path", "senders").Debug("starting watchdog")
	err = w.startV1_1Senders(secure)
	if err != nil {
		return fmt.Errorf("v1.1 senders error: %w", err)
	}
	w.waitForDataPause("v1.1", "/senders", 150*time.Millisecond)
	log.WithField("path", "receivers").Debug("starting watchdog")
	err = w.startV1_1Receivers(secure)
	if err != nil {
		return fmt.Errorf("v1.1 receivers error: %w", err)
	}
	w.waitForDataPause("v1.1", "/receivers", 150*time.Millisecond)
	return nil
}

func (w *Watchdog) startV1_1Nodes(secure bool) error {
	resourcePath := "/nodes"
	socket := stablews.New("v1.1", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_1.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_1.QUERY_WS_NODES {
				logrus.Error("bad topic")
				return
			}
			for _, data := range grain.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				nodeJson, err := json.Marshal(data.Post)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				var node is04v1_1.Node
				err = json.Unmarshal(nodeJson, &node)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_1Node(context.Background(), node)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.1"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_1Devices(secure bool) error {
	resourcePath := "/devices"
	socket := stablews.New("v1.1", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_1.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_1.QUERY_WS_DEVICES {
				logrus.Error("bad topic")
				return
			}
			for _, data := range grain.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				deviceJson, err := json.Marshal(data.Post)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				var device is04v1_1.Device
				err = json.Unmarshal(deviceJson, &device)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_1Device(context.Background(), device)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.1"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_1Sources(secure bool) error {
	resourcePath := "/sources"
	socket := stablews.New("v1.1", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_1.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_1.QUERY_WS_SOURCES {
				logrus.Error("bad topic")
				return
			}
			for _, data := range grain.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				respJson, err := json.Marshal(data.Post)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				var source is04v1_1.Source
				err = json.Unmarshal(respJson, &source)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_1Source(context.Background(), source)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.1"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_1Flows(secure bool) error {
	resourcePath := "/flows"
	socket := stablews.New("v1.1", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_1.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_1.QUERY_WS_FLOWS {
				logrus.Error("bad topic")
				return
			}
			for _, data := range grain.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				respJson, err := json.Marshal(data.Post)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				var item is04v1_1.Flow
				err = json.Unmarshal(respJson, &item)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_1Flow(context.Background(), item)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.1"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_1Senders(secure bool) error {
	resourcePath := "/senders"
	socket := stablews.New("v1.1", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_1.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_1.QUERY_WS_SENDERS {
				logrus.Error("bad topic")
				return
			}
			for _, data := range grain.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				respJson, err := json.Marshal(data.Post)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				var item is04v1_1.Sender
				err = json.Unmarshal(respJson, &item)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_1Sender(context.Background(), item)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.1"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_1Receivers(secure bool) error {
	resourcePath := "/receivers"
	socket := stablews.New("v1.1", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_1.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_1.QUERY_WS_RECEIVERS {
				logrus.Error("bad topic")
				return
			}
			for _, data := range grain.Grain.Data {
				if data.Post == nil {
					// Skip deletions
					continue
				}
				respJson, err := json.Marshal(data.Post)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				var item is04v1_1.Receiver
				err = json.Unmarshal(respJson, &item)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_1Receiver(context.Background(), item)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.1"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}
