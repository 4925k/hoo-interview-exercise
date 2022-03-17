package logger

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
)

const (
	logLimit = 10
)

// New() initliazes a logHanlder according to the given config file
func New(configFile string) logHandler {
	f, err := os.ReadFile(configFile)
	if err != nil {
		log.Fatalf("read config file at %s: %v", configFile, err)
	}

	var cfg config
	err = json.Unmarshal(f, &cfg)
	if err != nil {
		log.Fatalf("unmarshal config file content: %v", err)
	}

	return cfg.Initialize()
}

// Initialize() creates connections with the given config parameters
func (cfg *config) Initialize() (lh logHandler) {
	var err error

	if cfg.FilePath != "" {
		file, err := os.OpenFile(cfg.FilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Printf("open log file: %v", err)
		}

		lh.file = log.New(file, "", log.LstdFlags)
	}

	if cfg.UDP.IP != "" {
		addr := net.UDPAddr{IP: net.ParseIP(cfg.UDP.IP), Port: cfg.UDP.Port}
		lh.udp, err = net.DialUDP("udp", nil, &addr)
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	// logging aws to a file for mocking
	if cfg.AWS.Bucket != "" {
		//do aws connection stuff
		//call aws connection initialization
		file, err := os.OpenFile(cfg.AWS.Bucket, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
		if err != nil {
			log.Printf("aws connection: %v", err)
		}

		lh.aws = log.New(file, "", log.LstdFlags)
	}

	return
}
