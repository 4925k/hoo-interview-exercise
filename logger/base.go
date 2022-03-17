package logger

import (
	"log"
	"net"
	"sync"
)

// logHandler has pointers to different locations
// along with previous message, counter and a sync mutex
type logHandler struct {
	prevLog string
	count   int
	mu      sync.Mutex

	file *log.Logger
	aws  *log.Logger //file for now, appropriate type should be assigned that helps to write to cloud storage
	udp  *net.UDPConn
}

// main config
type config struct {
	FilePath string `json:"filePath"`
	AWS      aws    `json:"AWS"`
	UDP      udp    `json:"UDP"`
}

type aws struct {
	Bucket  string `json:"bucket"`
	Address string `json:"address"`
}

type udp struct {
	IP   string `json:"ip"`
	Port int    `json:"port"`
}
