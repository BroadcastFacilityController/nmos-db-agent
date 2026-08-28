package stablews

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type Client struct {
	nmosVersion  string
	nmosRootUrl  string // Ex: http://localhost:8080/x-nmos
	nmosPath     string
	nmosSecure   bool
	log          *logrus.Entry
	readFunc     func(messageType int, data []byte)
	mu           sync.RWMutex
	conn         *websocket.Conn
	ctx          context.Context
	cancel       context.CancelFunc
	minBackoff   time.Duration
	maxBackoff   time.Duration
	lastDataTime time.Time
	lastPingTime time.Time
}

func New(nmosVersion string, nmosRootUrl string, nmosPath string, nmosSecure bool, log *logrus.Logger) *Client {
	c := Client{
		nmosVersion: nmosVersion,
		nmosRootUrl: nmosRootUrl,
		nmosPath:    nmosPath,
		nmosSecure:  nmosSecure,
		log:         log.WithField("component", "websocket"),
		minBackoff:  100 * time.Millisecond,
		maxBackoff:  30 * time.Second,
	}
	return &c
}

func (c *Client) connectOnce() error {
	// Build NMOS post request
	var postReq *http.Request
	url := c.nmosRootUrl + "/query/" + c.nmosVersion + "/subscriptions"
	switch c.nmosVersion {
	case "v1.0":
		postBody, err := json.Marshal(struct {
			Max_update_rate_ms int      `json:"max_update_rate_ms"`
			Persist            bool     `json:"persist"`
			Resource_path      string   `json:"resource_path"`
			Params             struct{} `json:"params"`
		}{
			Max_update_rate_ms: 50,
			Persist:            false,
			Resource_path:      c.nmosPath,
			Params:             struct{}{},
		})
		if err != nil {
			c.log.WithError(err).Error("unable to create post request")
			return err
		}
		postReq, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
		if err != nil {
			return fmt.Errorf("unable to create post request: %w", err)
		}
		postReq.Header.Add("Accept", "application/json")
		postReq.Header.Add("Content-Type", "application/json")
	case "v1.1":
		postBody, err := json.Marshal(struct {
			Max_update_rate_ms int      `json:"max_update_rate_ms"`
			Persist            bool     `json:"persist"`
			Resource_path      string   `json:"resource_path"`
			Params             struct{} `json:"params"`
		}{
			Max_update_rate_ms: 50,
			Persist:            false,
			Resource_path:      c.nmosPath,
			Params:             struct{}{},
		})
		if err != nil {
			c.log.WithError(err).Error("unable to create post request")
			return err
		}
		postReq, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
		if err != nil {
			return fmt.Errorf("unable to create post request: %w", err)
		}
		postReq.Header.Add("Accept", "application/json")
		postReq.Header.Add("Content-Type", "application/json")
	case "v1.2":
		postBody, err := json.Marshal(struct {
			Max_update_rate_ms int      `json:"max_update_rate_ms"`
			Persist            bool     `json:"persist"`
			Resource_path      string   `json:"resource_path"`
			Params             struct{} `json:"params"`
		}{
			Max_update_rate_ms: 50,
			Persist:            false,
			Resource_path:      c.nmosPath,
			Params:             struct{}{},
		})
		if err != nil {
			c.log.WithError(err).Error("unable to create post request")
			return err
		}
		postReq, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
		if err != nil {
			return fmt.Errorf("unable to create post request: %w", err)
		}
		postReq.Header.Add("Accept", "application/json")
		postReq.Header.Add("Content-Type", "application/json")
	case "v1.3":
		postBody, err := json.Marshal(struct {
			Max_update_rate_ms int      `json:"max_update_rate_ms"`
			Persist            bool     `json:"persist"`
			Resource_path      string   `json:"resource_path"`
			Params             struct{} `json:"params"`
		}{
			Max_update_rate_ms: 50,
			Persist:            false,
			Resource_path:      c.nmosPath,
			Params:             struct{}{},
		})
		if err != nil {
			c.log.WithError(err).Error("unable to create post request")
			return err
		}
		postReq, err = http.NewRequest(http.MethodPost, url, bytes.NewReader(postBody))
		if err != nil {
			return fmt.Errorf("unable to create post request: %w", err)
		}
		postReq.Header.Add("Accept", "application/json")
		postReq.Header.Add("Content-Type", "application/json")
	default:
		return errors.New("unknown nmos version: " + c.nmosVersion)
	}
	resp, err := http.DefaultClient.Do(postReq)
	if err != nil {
		return fmt.Errorf("unable to post websocket subscription: %w", err)
	}
	if resp.StatusCode != http.StatusCreated && resp.StatusCode != http.StatusAccepted && resp.StatusCode != http.StatusOK {
		return errors.New("unable to post websocket subscription")
	}
	var body struct {
		ID      string `json:"id"`
		WS_href string `json:"ws_href"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return err
	}
	conn, _, err := websocket.DefaultDialer.Dial(body.WS_href, nil)
	if err != nil {
		return err
	}
	c.mu.Lock()
	c.conn = conn
	c.mu.Unlock()

	return nil
}

func (c *Client) pingHandler() {
	pingInterval := 100 * time.Millisecond
	ticker := time.NewTicker(pingInterval)
	defer func() {
		ticker.Stop()
	}()

	pingMsg, err := websocket.NewPreparedMessage(websocket.PingMessage, nil)
	if err != nil {
		c.log.WithError(err).Error("unable to create ping message")
		return
	}

	for range ticker.C {
		err := c.conn.WritePreparedMessage(pingMsg)
		if err != nil {
			c.log.WithError(err).Error("unable to send ping message. Reconnecting")
			c.connectSustained()
		}
		c.lastPingTime = time.Now()
	}
}

func (c *Client) readHandler() {
	for {
		if c.conn == nil {
			c.log.Error("connection closed; awaiting ping to reopen")
			return
		}
		msgType, p, err := c.conn.ReadMessage()
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
			c.log.WithError(err).Error("connection closed unexpectedly; awaiting ping to reopen")
			return
		}
		if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNormalClosure) {
			c.log.WithError(err).Error("connection closed; awaiting ping to reopen")
			return
		}
		if err != nil {
			c.log.WithError(err).Error("unable to parse received message")
			continue
		}
		c.lastDataTime = time.Now()
		if c.readFunc != nil {
			c.readFunc(msgType, p)
		}
	}
}

func (c *Client) connectSustained() {
	backoff := c.minBackoff
	for {
		err := c.connectOnce()
		if err == nil {
			go c.readHandler()
			c.log.WithFields(logrus.Fields{
				"nmos_version":  c.nmosVersion,
				"nmos_root_url": c.nmosRootUrl,
				"nmos_secure":   c.nmosSecure,
				"nmos_path":     c.nmosPath,
			}).Info("websocket connected")
			return
		}
		delay := c.jitter(backoff)
		c.log.WithField("url", c.nmosRootUrl).WithField("version", c.nmosVersion).WithError(err).Errorf("websocket failed to connect; retrying in %s", delay)
		backoff *= 2
		if backoff > c.maxBackoff {
			backoff = c.maxBackoff
		}
		time.Sleep(delay)
	}

}

func (c *Client) SetReadCallback(fun func(messageType int, data []byte)) error {
	c.readFunc = fun
	return nil
}

func (c *Client) Start() {
	c.connectSustained()
	go c.pingHandler()
}

func (c *Client) jitter(backoff time.Duration) time.Duration {
	return time.Duration((0.75 + rand.Float64()*0.5) * float64(backoff))
}

// Returns 0 if it has never received data
func (c *Client) GetTimeSinceLastPing() time.Duration {
	if c.lastPingTime.IsZero() {
		return 0
	}
	return time.Since(c.lastPingTime)
}

// Returns 0 if it has never received data
func (c *Client) GetTimeSinceLastData() time.Duration {
	if c.lastDataTime.IsZero() {
		return 0
	}
	return time.Since(c.lastDataTime)
}
