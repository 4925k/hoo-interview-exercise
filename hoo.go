package main

import (
	"hoo/logger"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup

	configPath := "config.json"

	// initialize log handling
	loggerA := logger.New(configPath)

	// check for 10 same logs only
	for i := 0; i < 15; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			loggerA.Log("only 10 logs of these should appear")
		}()
	}

	time.Sleep(1 * time.Millisecond)

	//check for concurrent writes
	for i := 0; i < 5; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			loggerA.Log("10 of concurrent log writes from A")
		}()

		go func() {
			defer wg.Done()
			loggerA.Log("10 of concurrent log writes from B")
		}()

	}

	time.Sleep(1 * time.Millisecond)

	//check for logging with parameters sent
	for i := 0; i < 5; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			loggerA.Log("log to file only", "file")
		}()

		go func() {
			defer wg.Done()
			loggerA.Log("log to file and aws only", "aws", "file")
		}()

	}

	//check for multiple users
	loggerB := logger.New(configPath)

	for i := 0; i < 5; i++ {
		wg.Add(2)

		go func() {
			defer wg.Done()
			loggerA.Log("5 logs from A")
		}()

		go func() {
			defer wg.Done()
			loggerB.Log("5 logs from B")
		}()
	}

	time.Sleep(1 * time.Millisecond)

	// log to udp
	for i := 0; i < 15; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()
			loggerA.Log("15 but 10 from user A to udp", "udp", "file")
		}()
	}

	wg.Wait()
}
