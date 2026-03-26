package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/BroadcastFacilityController/nmos-db-agent/database"
	"github.com/BroadcastFacilityController/nmos-db-agent/watchdog"
)

func main() {
	// Load from environment
	parseEnvironmentVariables()

	// Connect to database
	logrus.Info("Connecting to database nmos")
	_, err := database.NewConnection(context.Background(), _DB_USER, _DB_PASSWORD, _DB_URL, _DB_PORT)
	if err != nil {
		logrus.Fatal(err)
	}
	defer database.Close()
	logrus.Info("Connected to database nmos")

	// Get an NMOS watchdog
	watchdog, err := watchdog.NewWatchdog(_NMOS_REGISTRY_ADDRESS, _NMOS_REGISTRY_PORT, false)
	if err != nil {
		logrus.Fatal(err)
	}

	// Subscribe to everything via websockets to watch for changes
	err = watchdog.SubscribeNodes()
	if err != nil {
		logrus.Fatal(err)
	}
	// Wait for messages to be stable
	for {
		if watchdog.TimeSinceLastMessage() > (500 * time.Millisecond) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	err = watchdog.SubscribeDevices()
	if err != nil {
		logrus.Fatal(err)
	}
	// Wait for messages to be stable
	for {
		if watchdog.TimeSinceLastMessage() > (500 * time.Millisecond) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	err = watchdog.SubscribeSources()
	if err != nil {
		logrus.Fatal(err)
	}
	// Wait for messages to be stable
	for {
		if watchdog.TimeSinceLastMessage() > (500 * time.Millisecond) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	err = watchdog.SubscribeFlows()
	if err != nil {
		logrus.Fatal(err)
	}
	// Wait for messages to be stable
	for {
		if watchdog.TimeSinceLastMessage() > (500 * time.Millisecond) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	err = watchdog.SubscribeSenders()
	if err != nil {
		logrus.Fatal(err)
	}
	// Wait for messages to be stable
	for {
		if watchdog.TimeSinceLastMessage() > (500 * time.Millisecond) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}
	err = watchdog.SubscribeReceivers()
	if err != nil {
		logrus.Fatal(err)
	}
	// Wait for messages to be stable
	for {
		if watchdog.TimeSinceLastMessage() > (500 * time.Millisecond) {
			break
		}
		time.Sleep(200 * time.Millisecond)
	}

	// Hold until closed by a signal interrupt
	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGTERM)
	<-done
}
