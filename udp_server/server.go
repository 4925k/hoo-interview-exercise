package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

var (
	host = "0.0.0.0"
	port = 6969
)

func main() {
	addr := net.UDPAddr{IP: net.ParseIP(host), Port: port}
	handleServer(addr)
}

func handleServer(addr net.UDPAddr) {
	lstnr, err := net.ListenUDP("udp", &addr)
	if err != nil {
		log.Fatalf("[FATAL] set up listener: %v", err)
	}
	defer lstnr.Close()

	reader := bufio.NewReaderSize(lstnr, 10000)

	for {
		content, _, err := reader.ReadLine()
		if err != nil {
			log.Printf("[FATAL] read udp: %v", err)
		}

		fmt.Printf("[INFO] udp content: %v\n", string(content))
	}

}
