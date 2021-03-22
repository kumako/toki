package config

import "os"

// Name Application name
const Name = "toki"

// Application version
const Version = "0.0.1"

// DefaultTZ default timezone
const DefaultTZ = "Asia/Tokyo"

// DefaultPort http default port
const DefaultPort = "4242"

// ListeningPort Listening TCP port
func ListeningPort() string {
	port, exists := os.LookupEnv("PORT")
	if !exists {
		port = DefaultPort
	}
	return port
}
