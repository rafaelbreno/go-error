## Common way
- In a GoRoutine
- Killing the goroutine when an error occurs

```go
// ch receives random values to be processed
var ch chan int

// flag will receive a bool
// value when an error occurs,
// stopping the application
var flag chan bool

func init() {
	ch = make(chan int, 10)
	flag = make(chan bool)
}

func sendMessages() {
	messages := []string{"1", "2", "3", "foo", "4", "bar", "5", "6"}
	for _, msg := range messages {
		num, err := strconv.Atoi(msg)
		if err != nil {
			log.Println(err.Error())
			flag <- true
		}
		ch <- num
		time.Sleep(1 * time.Second)
	}
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
		case <-flag:
			return
		}
	}
}
```
