package watchdog

import (
	"errors"
	"time"

	"github.com/BroadcastFacilityController/nmos-db-agent/nmosdb"
	"github.com/BroadcastFacilityController/nmos-db-agent/watchdog/stablews"
	"github.com/BroadcastFacilityController/nmos-go-client"
	"github.com/sirupsen/logrus"
)

type Watchdog struct {
	db               *nmosdb.DB
	queryEndpoint    *nmos.NMOSEndpoint
	queryEndpointUrl string
	sockets          map[string]map[string]*stablews.Client
}

func NewWatchdog(db *nmosdb.DB, queryEndpointUrl string) *Watchdog {
	endpoint, err := nmos.NewNMOSEndpoint(queryEndpointUrl)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"location": "watchdog",
			"url":      queryEndpointUrl,
		}).WithError(err).Error("unable to create nmos endpoint")
		return nil
	}
	w := Watchdog{
		db:               db,
		queryEndpoint:    endpoint,
		queryEndpointUrl: queryEndpointUrl,
		sockets:          make(map[string]map[string]*stablews.Client),
	}
	return &w
}

func (w *Watchdog) Start() error {
	err := w.startV1_3(false)
	if err != nil {
		return err
	}
	err = w.startV1_2(false)
	if err != nil {
		return err
	}
	err = w.startV1_1(false)
	if err != nil {
		return err
	}
	err = w.startV1_0(false)
	if err != nil {
		return err
	}
	return nil
}

// This is a wait. You should not run this from the main thread / things you don't want to wait for
func (w *Watchdog) waitForDataPause(version string, path string, pauseTime time.Duration) error {
	socketsVersion, socketsVersion_ok := w.sockets[version]
	if !socketsVersion_ok {
		return errors.New("sockets version not created")
	}
	socket, socket_ok := socketsVersion[path]
	if !socket_ok {
		return errors.New("socket at path not found")
	}
	for {
		timeSince := socket.GetTimeSinceLastData()
		if timeSince > pauseTime {
			return nil
		}
		waitTime := pauseTime - timeSince
		time.Sleep(waitTime)
	}

}
