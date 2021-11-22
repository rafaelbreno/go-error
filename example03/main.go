package main

import (
	"fmt"
	"log"
	"strconv"
	"time"
)

type Foo struct {
	IntString string
	IntInt    int
	Error     error
}

// Note that this function doesn't
// return an error
func (f *Foo) process() {
	v, err := strconv.Atoi(f.IntString)
	f.IntInt = v
	f.Error = err
}

var foos []Foo
var fooCh chan Foo
var stop chan bool

func init() {
	foos = []Foo{
		{
			IntString: "1",
		},
		{
			IntString: "2",
		},
		{
			IntString: "3",
		},
		{
			IntString: "w",
		},
		{
			IntString: "4",
		},
		{
			IntString: "5",
		},
		{
			IntString: "e",
		},
		{
			IntString: "6",
		},
	}
	fooCh = make(chan Foo)
	stop = make(chan bool)
}

// process each item from `foos`
// and send it to fooCh
func processAll() {
	for _, f := range foos {
		f.process()
		fooCh <- f
		time.Sleep(time.Second)
	}

	stop <- true
}

func main() {
	go processAll()

	for {
		select {
		case f := <-fooCh:
			if f.Error != nil {
				log.Println(f.Error.Error())
				continue
			}
			fmt.Println("Received! ", f.IntInt)
		case <-stop:
			return
		}
	}
}
