## Custom Error implementation
- When we want to create our own error system, we may want to implement the `error interface` 
- From the Go's source code, we have this:
```go
type error interface {
  Error() string
}
```
- Meaning that our error `struct` must have _at least_ the function `Error() string`
- CustomErr implementation
```go
const (
	// TAG01
	// There's no Database Connection
	TAG01 = "TAG01"
	// TAG02
	// Parsing error
	TAG02 = "TAG02"
	// TAG03
	// Validation errors
	TAG03 = "TAG03"
)

const (
	// TAG|TIMESTAMP|ERR_MESSAGE
	customErrTemplate = `%s|%s|%s`
)

// CustomErr
// The custom error struct
type CustomErr struct {
	Tag string
	Msg string
}

// Error is the obligatory function
// so it can match the error interface
func (c *CustomErr) Error() string {
	return fmt.Sprintf(customErrTemplate,
		c.Tag,
		time.Now().String(),
		c.Msg)
}
```

- Using the previous built `CustomErr`
```go
type Foo struct {
	IntString string
	IntInt    int
	Error     error
}

// Note that this function doesn't
// return an error
func (f *Foo) process() {
	if f.IntString == "" {
		f.Error = &CustomErr{
			Tag: TAG03,
			Msg: `"IntString" field is required`,
		}
		return
	}

	v, err := strconv.Atoi(f.IntString)
	if err != nil {
		f.Error = &CustomErr{
			Tag: TAG02,
			Msg: err.Error(),
		}
		return
	}
	f.IntInt = v
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
			IntString: "",
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

```
