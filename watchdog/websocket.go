package watchdog

import (
	"context"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/BroadcastFacilityController/nmos-go-client"
	"github.com/BroadcastFacilityController/nmos-go-client/common"
	is04v1_0 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.0"
	is04v1_1 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.1"
	is04v1_2 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.2"
	is04v1_3 "github.com/BroadcastFacilityController/nmos-go-client/is04/v1.3"
	"github.com/gorilla/websocket"
)

type Websocket struct {
	path                 string
	conn                 *websocket.Conn
	resp                 *http.Response
	version              string
	endpoint             *nmos.NMOSEndpoint
	lastMessage          time.Time
	firstMessageReceived bool
}

func NewWebsocketV1_0(endpoint *nmos.NMOSEndpoint, path is04v1_0.QueryAPISubscriptionResourcePath) (*Websocket, error) {
	apis, err := endpoint.GetSupportedAPIs()
	if err != nil {
		return nil, err
	}
	if !slices.Contains(apis, common.QUERY) {
		return nil, fmt.Errorf("endpoint does not support query api")
	}
	is04versions, err := endpoint.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return nil, err
	}
	supportsVersion := false
	for _, vers := range is04versions {
		if vers.Equals(common.NewAPIVersion(1, 0)) {
			supportsVersion = true
			break
		}
	}
	if !supportsVersion {
		return nil, fmt.Errorf("endpoint does not support api version")
	}
	api := endpoint.IS04().V1_0()

	requestBody := is04v1_0.QueryAPISubscriptionPost{
		MaxUpdateRateMS: 50,
		Persist:         false,
		ResourcePath:    path,
		Params:          nil,
	}
	subscription, err := api.QueryPostSubscription(requestBody)
	if err != nil {
		return nil, err
	}
	conn, resp, err := websocket.DefaultDialer.Dial(subscription.WSHRef, nil)
	if err != nil {
		return nil, err
	}
	ws := Websocket{
		path:     string(path),
		conn:     conn,
		resp:     resp,
		version:  "v1.0",
		endpoint: endpoint,
	}
	return &ws, nil
}

func NewWebsocketV1_1(endpoint *nmos.NMOSEndpoint, path is04v1_1.QueryAPISubscriptionResourcePath) (*Websocket, error) {
	apis, err := endpoint.GetSupportedAPIs()
	if err != nil {
		return nil, err
	}
	if !slices.Contains(apis, common.QUERY) {
		return nil, fmt.Errorf("endpoint does not support query api")
	}
	is04versions, err := endpoint.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return nil, err
	}
	supportsVersion := false
	for _, vers := range is04versions {
		if vers.Equals(common.NewAPIVersion(1, 1)) {
			supportsVersion = true
			break
		}
	}
	if !supportsVersion {
		return nil, fmt.Errorf("endpoint does not support api version")
	}
	api := endpoint.IS04().V1_1()

	requestBody := is04v1_1.QueryAPISubscriptionPost{
		MaxUpdateRateMS: 50,
		Persist:         false,
		ResourcePath:    path,
		Params:          nil,
	}
	subscription, err := api.QueryPostSubscription(requestBody)
	if err != nil {
		return nil, err
	}
	conn, resp, err := websocket.DefaultDialer.Dial(subscription.WSHRef, nil)
	if err != nil {
		return nil, err
	}
	ws := Websocket{
		path:     string(path),
		conn:     conn,
		resp:     resp,
		version:  "v1.1",
		endpoint: endpoint,
	}
	return &ws, nil
}

func NewWebsocketV1_2(endpoint *nmos.NMOSEndpoint, path is04v1_2.QueryAPISubscriptionResourcePath) (*Websocket, error) {
	apis, err := endpoint.GetSupportedAPIs()
	if err != nil {
		return nil, err
	}
	if !slices.Contains(apis, common.QUERY) {
		return nil, fmt.Errorf("endpoint does not support query api")
	}
	is04versions, err := endpoint.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return nil, err
	}
	supportsVersion := false
	for _, vers := range is04versions {
		if vers.Equals(common.NewAPIVersion(1, 2)) {
			supportsVersion = true
			break
		}
	}
	if !supportsVersion {
		return nil, fmt.Errorf("endpoint does not support api version")
	}
	api := endpoint.IS04().V1_2()

	requestBody := is04v1_2.QueryAPISubscriptionPost{
		MaxUpdateRateMS: 50,
		Persist:         false,
		ResourcePath:    path,
		Params:          nil,
	}
	subscription, err := api.QueryPostSubscription(requestBody)
	if err != nil {
		return nil, err
	}
	conn, resp, err := websocket.DefaultDialer.Dial(subscription.WSHRef, nil)
	if err != nil {
		return nil, err
	}
	ws := Websocket{
		path:     string(path),
		conn:     conn,
		resp:     resp,
		version:  "v1.2",
		endpoint: endpoint,
	}
	return &ws, nil
}

func NewWebsocketV1_3(endpoint *nmos.NMOSEndpoint, path is04v1_3.QueryAPISubscriptionResourcePath) (*Websocket, error) {
	apis, err := endpoint.GetSupportedAPIs()
	if err != nil {
		return nil, err
	}
	if !slices.Contains(apis, common.QUERY) {
		return nil, fmt.Errorf("endpoint does not support query api")
	}
	is04versions, err := endpoint.IS04().GetAPIVersions(common.QUERY)
	if err != nil {
		return nil, err
	}
	supportsVersion := false
	for _, vers := range is04versions {
		if vers.Equals(common.NewAPIVersion(1, 3)) {
			supportsVersion = true
			break
		}
	}
	if !supportsVersion {
		return nil, fmt.Errorf("endpoint does not support api version")
	}
	api := endpoint.IS04().V1_3()

	requestBody := is04v1_3.QueryAPISubscriptionPost{
		MaxUpdateRateMS: 50,
		Persist:         false,
		ResourcePath:    path,
		Params:          nil,
	}
	subscription, err := api.QueryPostSubscription(requestBody)
	if err != nil {
		return nil, err
	}
	conn, resp, err := websocket.DefaultDialer.Dial(subscription.WSHRef, nil)
	if err != nil {
		return nil, err
	}
	ws := Websocket{
		path:     string(path),
		conn:     conn,
		resp:     resp,
		version:  "v1.3",
		endpoint: endpoint,
	}
	return &ws, nil
}

func (ws *Websocket) Reconnect() error {
	// Ensure connection is dropped
	if ws.conn != nil {
		ws.conn.Close()
	}

	switch ws.version {
	case "v1.0":
		temp, err := NewWebsocketV1_0(ws.endpoint, is04v1_0.QueryAPISubscriptionResourcePath(ws.path))
		if err != nil {
			return err
		}
		ws.conn = temp.conn
		return nil
	case "v1.1":
		temp, err := NewWebsocketV1_1(ws.endpoint, is04v1_1.QueryAPISubscriptionResourcePath(ws.path))
		if err != nil {
			return err
		}
		ws.conn = temp.conn
		return nil
	case "v1.2":
		temp, err := NewWebsocketV1_2(ws.endpoint, is04v1_2.QueryAPISubscriptionResourcePath(ws.path))
		if err != nil {
			return err
		}
		ws.conn = temp.conn
		return nil
	case "v1.3":
		temp, err := NewWebsocketV1_3(ws.endpoint, is04v1_3.QueryAPISubscriptionResourcePath(ws.path))
		if err != nil {
			return err
		}
		ws.conn = temp.conn
		return nil
	default:
		return fmt.Errorf("unrecognized version %s", ws.version)
	}
}

func (ws *Websocket) ReconnectUntilConnected() {
	err := ws.Reconnect()
	if err != nil {
		ws.ReconnectUntilConnected()
	}
}

func (ws *Websocket) ReadMessage() (messageType int, p []byte, err error) {
	messageType, p, err = ws.conn.ReadMessage()
	if err != nil {
		// Check if connection was dropped
		if ws.conn == nil {
			// Reconnect
			err2 := ws.Reconnect()
			if err2 != nil {
				// Reconnect failed, pass the error up
				err = err2
				return
			}
			// Try again, pass the error up if it happens to prevent infinite recursion
			messageType, p, err = ws.ReadMessage()
			return
		} else {
			// Return the error
			return
		}
	}
	ws.lastMessage = time.Now()
	if !ws.firstMessageReceived {
		ws.firstMessageReceived = true
	}
	return
}

func (ws *Websocket) AwaitFirstMessage(ctx context.Context) error {
	for {
		if ctx.Err() != nil {
			return ctx.Err()
		}
		if ws.firstMessageReceived {
			return nil
		}
		time.Sleep(100 * time.Millisecond)
	}
}
