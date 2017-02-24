package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	shutdown := EventListener()

	go func() {
		for {
			log.Println("tick")
			time.Sleep(time.Second * 5)
		}
	}()

	<-shutdown
}

// EventListener listens for SIGINT and SIGTERM signals and notifies the
// shutdown channel if it detects that either one was sent.
func EventListener() <-chan bool {
	shutdown := make(chan bool)

	go func() {
		ch := make(chan os.Signal)
		signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)

		for {
			select {
			case <-ch:
				log.Println("shutdown signal received")
				shutdown <- true
				break
			}
		}
	}()

	return shutdown
}
