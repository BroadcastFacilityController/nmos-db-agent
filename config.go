package main

import (
	"context"
	"os"
	"strconv"
	"time"

	"github.com/grandcat/zeroconf"
	"github.com/sirupsen/logrus"
)

var _LOG_LEVEL logrus.Level
var _DB_URL string
var _DB_PORT string
var _DB_USER string
var _DB_PASSWORD string
var _NMOS_REGISTRY_ADDRESS string
var _NMOS_REGISTRY_PORT int

// Gets the config from ENV variables and loads them into global variables
func parseEnvironmentVariables() {
	// Log level
	debugLvl, debugLvl_ok := os.LookupEnv("DEBUG_LEVEL")
	if !debugLvl_ok || debugLvl == "" {
		logrus.Warn("No DEBUG_LEVEL environment variable supplied. Using level \"info\"")
		debugLvl = "info"
	}
	logLvl, err := logrus.ParseLevel(debugLvl)
	if err != nil {
		logrus.Fatalf("Unable to parse level: %s: %v", debugLvl, err)
	}
	_LOG_LEVEL = logLvl
	logrus.SetLevel(_LOG_LEVEL)

	// Database Connection
	dbUrl, dbUrl_ok := os.LookupEnv("DB_URL")
	if !dbUrl_ok {
		logrus.Fatal("No DB_URL environment variable supplied")
	}
	dbPort, dbPort_ok := os.LookupEnv("DB_PORT")
	if !dbPort_ok {
		logrus.Fatal("No DB_PORT environment variable supplied")
	}
	dbUser, dbUser_ok := os.LookupEnv("DB_USER")
	if !dbUser_ok {
		logrus.Fatal("No DB_USER environment variable supplied")
	}
	dbPassword, dbPassword_ok := os.LookupEnv("DB_PASSWORD")
	if !dbPassword_ok {
		logrus.Fatal("No DB_PASSWORD environment variable supplied")
	}
	_DB_URL = dbUrl
	_DB_PORT = dbPort
	_DB_USER = dbUser
	_DB_PASSWORD = dbPassword

	// Registry Connection
	registryAddress, registryAddress_ok := os.LookupEnv("NMOS_REGISTRY_ADDRESS")
	registryAutoPort := -1
	if !registryAddress_ok {
		logrus.Warn("No NMOS_REGISTRY_ADDRESS environment variable supplied")
		// Attempt to find via mDNS
		logrus.Info("Attempting to find NMOS registry via mDNS / DNS-SD")
		resolver, err := zeroconf.NewResolver(nil)
		if err != nil {
			logrus.Fatal(err)
		}
		registrySearchDomain, registrySearchDomain_ok := os.LookupEnv("NMOS_REGISTRY_DOMAIN")
		if !registrySearchDomain_ok {
			logrus.Warn("No NMOS_REGISTRY_DOMAIN environment variable supplied. Using local. instead")
			registrySearchDomain = "local."
		}
		registryEntriesChan := make(chan *zeroconf.ServiceEntry, 2)
		registryEntriesChanStop := make(chan bool)
		registryEntries := make([]*zeroconf.ServiceEntry, 0)
		ctx, ctxCancel := context.WithTimeout(context.Background(), 10*time.Second)
		go func() {
			for {
				select {
				case entry := <-registryEntriesChan:
					logrus.Infof("Found NMOS registry via mDNS / DNS-SD: %v", entry)
					registryEntries = append(registryEntries, entry)
				case <-registryEntriesChanStop:
					close(registryEntriesChan)
					close(registryEntriesChanStop)
					ctxCancel()
					return
				}
			}
		}()

		err = resolver.Browse(ctx, "_nmos-query._tcp", registrySearchDomain, registryEntriesChan)
		if err != nil {
			logrus.Fatalf("Unable to search for NMOS registry via mDNS / DNS-SD: %v", err)
		}
		// Wait for timeout
		<-ctx.Done()
		registryEntriesChanStop <- true
		if len(registryEntries) < 1 {
			logrus.Fatal("Unable to discover NMOS registry via mDNS / DNS-SD")
		}
		registryAutoPort = registryEntries[0].Port
		registryAddress = registryEntries[0].HostName
		logrus.Infof("Using NMOS registry: %s", registryAddress)
	}
	registryPort, registryPort_ok := os.LookupEnv("NMOS_REGISTRY_PORT")
	if !registryPort_ok {
		if registryAutoPort != -1 {
			logrus.Infof("No NMOS_REGISTRY_PORT enviornment variable supplied. Using auto-discovered value of %d", registryAutoPort)
			registryPort = strconv.Itoa(registryAutoPort)
		} else {
			logrus.Warn("No NMOS_REGISTRY_PORT environment variable supplied. Defaulting to 80.")
			registryPort = "80"
		}
	}
	registryPortParsed, err := strconv.Atoi(registryPort)
	if err != nil {
		logrus.Fatalf("Unable to parse port %s: %v", registryPort, err)
	}
	_NMOS_REGISTRY_ADDRESS = registryAddress
	_NMOS_REGISTRY_PORT = registryPortParsed

}
