package main

import (
	"fmt"
	"time"

	"github.com/epazote/scheduler"
)

func main() {
	// Create new scheduler
	s := scheduler.New()

	every := time.Second
	// Add a scheduled function
	s.AddScheduler("every second", every, func() {
		fmt.Println("Second passed")
	})

	// Let scheduler run for five seconds
	time.Sleep(5 * time.Second)

	// Stop the scheduled "every second" function
	err := s.Stop("every second")
	if err != nil {
		panic(err)
	}

	// Scheduler has now stopped
	time.Sleep(5 * time.Second)
}
