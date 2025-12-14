package main

import (
	r "canopy/internal/repository"
	s "canopy/internal/service"
	"time"
)

func main() {
	for {
		time.Sleep(5 * time.Second)
		payload := s.Transciever()
		r.Reciever(payload)
	}
}
