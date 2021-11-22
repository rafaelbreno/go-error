package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

// ch receives random values to be processed
var ch chan int

// errCh will receive an error
var errCh chan error

// stop is a flag to stop
// the application
var stop chan bool

func init() {
	ch = make(chan int)
	errCh = make(chan error)
	stop = make(chan bool)
}

func sendMessages() {
	messages := []string{"1", "2", "3", "foo", "4", "bar", "5", "6"}
	for _, msg := range messages {
		num, err := strconv.Atoi(msg)
		if err != nil {
			log.Println(err.Error())
			errCh <- err
			continue
		}
		ch <- num
		time.Sleep(1 * time.Second)
	}
	stop <- true
}

func main() {
	// Start sending values to `ch`
	go sendMessages()
	for {
		select {
		// handles the value
		// when received
		case n := <-ch:
			fmt.Println("Received: ", n)

		// breaks loop when an
		// error occurs
		case err := <-errCh:
			// handles the error
			fmt.Println(err.Error())

		case <-stop:
			return
		}
	}
}
