package config

import (
	"errors"
	"fmt"
	"net"
	"os"

	"github.com/J-Rivard/blockchain-poc/internal/logging"
)

type Config struct {
	HostPort string

	Host        string
	InitialPeer string
}

const (
	hostPort    = "HostPort"
	initialPeer = "InitialPeer"
)

func New(logger *logging.Log) (*Config, error) {
	if v := validateEnvironment(); v != nil {
		return nil, v
	}

	host := "http://" + getOutboundIP() + ":" + os.Getenv(hostPort)

	logger.LogApplication(logging.FormattedLog{
		"host":        host,
		"initialPeer": os.Getenv(initialPeer),
	})

	return &Config{
		HostPort:    os.Getenv(hostPort),
		Host:        host,
		InitialPeer: os.Getenv(initialPeer),
	}, nil
}

func validateEnvironment() error {
	requiredEnvVars := []string{hostPort}

	missingEnvVars := ""

	for _, v := range requiredEnvVars {
		if os.Getenv(v) == "" {
			missingEnvVars += v + ","
		}
	}

	if missingEnvVars != "" {
		return errors.New(fmt.Sprintf("Missing env vars: %s", missingEnvVars))
	}

	return nil
}

// Get preferred outbound ip of this machine
func getOutboundIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		// check the address type and if it is not a loopback the display it
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
