package main

import (
	"context"
	"os"
	"os/signal"
	"strconv"
	"syscall"

	"github.com/BroadcastFacilityController/nmos-db-agent/nmosdb"
	"github.com/BroadcastFacilityController/nmos-db-agent/watchdog"
	"github.com/sirupsen/logrus"
)

func main() {
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	// Load environment variables
	logrus.Info("Loading environment variables")
	// Logging
	log_level_str := os.Getenv("LOG_LEVEL") // Logging level. Ex: "debug" "error"
	if log_level_str != "" {
		logLvl, err := logrus.ParseLevel(log_level_str)
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"current_level": logrus.GetLevel(),
				"env":           log_level_str,
				"error":         err,
			}).Error("cannot parse LOG_LEVEL, maintaining current level")
		}
		logrus.WithFields(logrus.Fields{
			"old_level": logrus.GetLevel(),
			"new_level": logLvl,
		}).Info("changing logging level")
		logrus.SetLevel(logLvl)
	}
	// DB URL
	nmos_db_url := os.Getenv("NMOS_DB_URL") // Connection URL. Ex: "postgres://user:pass@localhost:5432/nmos?sslmode=disable"
	// DB Migrations
	nmos_db_migrations_str := os.Getenv("NMOS_DB_RUN_MIGRATIONS") // Run migrations. Ex: "true"
	nmos_db_migrations, err := strconv.ParseBool(nmos_db_migrations_str)
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"env":   nmos_db_migrations_str,
			"error": err,
		}).Warn("cannot parse NMOS_DB_RUN_MIGRATIONS; falling back to false")
		nmos_db_migrations = false
	}
	// NMOS Query API
	nmos_query_url := os.Getenv("NMOS_QUERY_URL") // NMOS Query URL. Ex: "http://localhost:8080/x-nmos"
	if nmos_query_url == "" {
		logrus.Error("cannot parse NMOS_QUERY_URL; exiting")
		return
	}
	logrus.Info("Finished loading environment variables")

	logrus.Info("Connecting to database")
	db, err := nmosdb.Open(context.Background(), nmosdb.Config{
		DatabaseURL:   nmos_db_url,
		RunMigrations: nmos_db_migrations,
	})
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("error connecting to database; exiting")
		return
	}
	defer db.Close()
	logrus.Info("Connected to database")

	logrus.Info("Starting Watchdog")
	wd := watchdog.NewWatchdog(db, nmos_query_url)
	err = wd.Start()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"error": err,
		}).Error("could not start watchdog; exiting")
		return
	}
	logrus.Info("Watchdog running")

	logrus.Info("Waiting for SIGSEV")
	quitCh := make(chan os.Signal, 10)
	signal.Notify(quitCh, os.Interrupt, syscall.SIGTERM)

	<-quitCh
}
