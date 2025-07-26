package main

import (
	"errors"
	"fmt"
	"math/rand"
	"net"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func main() {
	srvRun := true

	go start(10, &srvRun)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	srvRun = false

	fmt.Println("Shutting down gracefully...")
}

func start(total int, srvRun *bool) {
	var wg sync.WaitGroup

	for i := 0; i < total; i++ {
		fmt.Printf("Starting connection %d...\n", i+1)

		wg.Add(1)

		go connect(srvRun, &wg)

	}
}

func connect(srvRun *bool, wg *sync.WaitGroup) {
	for *srvRun {
		// Simulate connection logic ...
		if err := handleConnection(wg); err != nil {
			*srvRun = false
		}
	}
}

func handleConnection(wg *sync.WaitGroup) error {
	defer wg.Done()
	// Simulate handling connection ...
	for {
		select {
		case <-time.After(100 * time.Millisecond):
		default:
			if err := receive(); err != nil {
				// Handle error ...
				fmt.Printf("Error receiving data: %v\n", err)
				return fmt.Errorf("Error receiving data")
			}
		}
	}
}

func receive() error {
	// Simulate receiving data ...

	time.Sleep(200 * time.Millisecond)

	r := rand.Intn(20-1) + 1

	if r%2 == 0 {
		// Simulate an error condition
		return &net.OpError{Op: "receive", Err: errors.New("simulated error")}
	}

	return nil
}
