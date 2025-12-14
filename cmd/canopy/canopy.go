package main

import (
	l "canopy/internal/logger"
	m "canopy/internal/messages"
	s "canopy/internal/service"
	"time"
)

func main() {
	var ch = make(chan []m.Sorter)

	go l.Logman()

	for i := 1; i < 5; i++ {
		go s.Transciever(ch, i)
		go s.Reciever(ch, i)
	}

	for {
		time.Sleep(200 * time.Millisecond)
	}

}
