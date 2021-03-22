package api

import (
	"time"

	"github.com/sirupsen/logrus"
	"github.com/valyala/fasthttp"

	"toki/config"
)

// Start spawns a new web server - Will fatal if error
func Start() (*fasthttp.Server, string) {

	// Listening interface
	addr := "0.0.0.0:" + config.ListeningPort()

	// Init Server
	server := fasthttp.Server{
		Name:        "Toki " + config.Version,
		Handler:     routes(),
		IdleTimeout: time.Second,
	}

	// Start Server
	go func() {
		err := server.ListenAndServe(addr)
		if err != nil {
			logrus.Fatal(err)
		}
	}()

	return &server, addr
}
