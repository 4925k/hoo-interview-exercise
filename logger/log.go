package logger

import (
	"log"
)

// Log() will write given msg to all available locations
func (l *logHandler) Log(msg string, location ...string) {
	// checking whether the msg has been logged multiple times before
	l.mu.Lock()
	if msg == l.prevLog {
		l.count++
	} else {
		l.prevLog = msg
		l.count = 0
	}
	l.mu.Unlock()
	if l.count >= logLimit {
		return
	}

	// logs to all the locations if nothing is specified in the parameters
	if len(location) == 0 {
		l.file.Println(msg)
		l.aws.Println(msg)
		l.udp.Write([]byte(msg + "\n")) //error handling omitted
		return
	}

	// log to the locations provided in the parameters
	for _, v := range location {
		switch v {
		case "file":
			l.file.Println(msg)
		case "aws":
			l.aws.Println(msg)
		case "udp":
			l.udp.Write([]byte(msg + "\n")) //error handling omitted
		default:
			log.Printf("invalid location: %v", v)
		}

	}
}
