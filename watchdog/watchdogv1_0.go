package watchdog

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/watchdog/stablews"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	"github.com/gorilla/websocket"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"
)

func (w *Watchdog) startV1_0(secure bool) error {
	w.sockets["v1.0"] = make(map[string]*stablews.Client)
	var err error
	log := logrus.WithFields(logrus.Fields{
		"location": "watchdog",
		"version":  "v1.0",
	})
	log.WithField("path", "nodes").Debug("starting watchdog")
	err = w.startV1_0Nodes(secure)
	if err != nil {
		return fmt.Errorf("v1.0 nodes error: %w", err)
	}
	w.waitForDataPause("v1.0", "/nodes", 150*time.Millisecond)
	log.WithField("path", "devices").Debug("starting watchdog")
	err = w.startV1_0Devices(secure)
	if err != nil {
		return fmt.Errorf("v1.0 devices error: %w", err)
	}
	w.waitForDataPause("v1.0", "/devices", 150*time.Millisecond)
	log.WithField("path", "sources").Debug("starting watchdog")
	err = w.startV1_0Sources(secure)
	if err != nil {
		return fmt.Errorf("v1.0 sources error: %w", err)
	}
	w.waitForDataPause("v1.0", "/sources", 150*time.Millisecond)
	log.WithField("path", "flows").Debug("starting watchdog")
	err = w.startV1_0Flows(secure)
	if err != nil {
		return fmt.Errorf("v1.0 flows error: %w", err)
	}
	w.waitForDataPause("v1.0", "/flows", 150*time.Millisecond)
	log.WithField("path", "senders").Debug("starting watchdog")
	err = w.startV1_0Senders(secure)
	if err != nil {
		return fmt.Errorf("v1.0 senders error: %w", err)
	}
	w.waitForDataPause("v1.0", "/senders", 150*time.Millisecond)
	log.WithField("path", "receivers").Debug("starting watchdog")
	err = w.startV1_0Receivers(secure)
	if err != nil {
		return fmt.Errorf("v1.0 receivers error: %w", err)
	}
	w.waitForDataPause("v1.0", "/receivers", 150*time.Millisecond)
	return nil
}

func (w *Watchdog) startV1_0Nodes(secure bool) error {
	resourcePath := "/nodes"
	socket := stablews.New("v1.0", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_0.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_0.QUERY_WS_NODES {
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
				var node is04v1_0.Node
				err = json.Unmarshal(nodeJson, &node)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_0Node(context.Background(), node)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.0"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_0Devices(secure bool) error {
	resourcePath := "/devices"
	socket := stablews.New("v1.0", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_0.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_0.QUERY_WS_DEVICES {
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
				var device is04v1_0.Device
				err = json.Unmarshal(deviceJson, &device)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_0Device(context.Background(), device)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.0"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_0Sources(secure bool) error {
	resourcePath := "/sources"
	socket := stablews.New("v1.0", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_0.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_0.QUERY_WS_SOURCES {
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
				var source is04v1_0.Source
				err = json.Unmarshal(respJson, &source)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_0Source(context.Background(), source)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.0"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_0Flows(secure bool) error {
	resourcePath := "/flows"
	socket := stablews.New("v1.0", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_0.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_0.QUERY_WS_FLOWS {
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
				var item is04v1_0.Flow
				err = json.Unmarshal(respJson, &item)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_0Flow(context.Background(), item)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.0"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_0Senders(secure bool) error {
	resourcePath := "/senders"
	socket := stablews.New("v1.0", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_0.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_0.QUERY_WS_SENDERS {
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
				var item is04v1_0.Sender
				err = json.Unmarshal(respJson, &item)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_0Sender(context.Background(), item)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.0"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}

func (w *Watchdog) startV1_0Receivers(secure bool) error {
	resourcePath := "/receivers"
	socket := stablews.New("v1.0", w.queryEndpointUrl, resourcePath, false, logrus.StandardLogger())
	callback := func(messageType int, data []byte) {
		switch messageType {
		case websocket.TextMessage:
			var grain is04v1_0.QueryAPISubscriptionWSGrain
			err := json.Unmarshal(data, &grain)
			if err != nil {
				logrus.Error(err)
				return
			}
			if grain.Grain.Topic != is04v1_0.QUERY_WS_RECEIVERS {
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
				var item is04v1_0.Receiver
				err = json.Unmarshal(respJson, &item)
				if err != nil {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
				err = w.handleV1_0Receiver(context.Background(), item)
				if err != nil && err != pgx.ErrNoRows {
					logrus.WithFields(logrus.Fields{
						"grain": data,
					}).Error(err)
					continue
				}
			}
		}
	}
	w.sockets["v1.0"][resourcePath] = socket
	socket.SetReadCallback(callback)
	socket.Start()
	return nil
}
