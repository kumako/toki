package main

import (
	"os"
	"os/signal"
	"syscall"
	"time"

	"toki/api"
	"toki/config"

	"github.com/sirupsen/logrus"
)

func main() {
	startTime := time.Now()

	// Start Web Server
	server, _ := api.Start()

	// Started
	logrus.Printf("Toki %s started on %s - Ready on http://127.0.0.1:%s/", config.Version, time.Since(startTime), config.ListeningPort())

	// Wait for int/term signal
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)
	<-exit

	// Stop Application
	logrus.Print("Closing HTTP Server...")
	err := server.Shutdown()
	if err != nil {
		logrus.Fatal("Cannot gracefully shutdown webserver", err)
	}
}
