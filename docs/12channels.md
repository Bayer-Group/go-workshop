Covered in this module:

* channels

## channels
Channels are mono-directional pipes that goroutines can use to communicate.

To create a channel use `make(chan <type>)` where type is the type of the values that will be passed in the channel.

To send a value **to** a channel you use the syntax `channel <- value`.

To receive a value **from** a channel, use `<-channel` and capture the retrieved value into a variable like `variable := <-channel`:
```go
func main() {
	randInts := make(chan int)

	go roll1d6(randInts)

	roll := <-randInts // this pulls new rolls from the channel
	fmt.Println("You rolled a", roll)
}

func roll1d6(c chan int) {
    c <- rand.Intn(6) + 1 // this passes new rolls into the channel
}
```

If you are just using the channel as a _ping_ mechanism (i.e. you don't care about the value in the channel), you can receive with simply `<-channel` without capturing the retrieved value.
```go
func main() {
	ping := make(chan bool)
	startTime := time.Now()

	go waitForPing(ping)

	<-ping
	fmt.Println("Program took", time.Since(startTime))
}

func waitForPing(c chan bool) {
	time.Sleep(time.Second * 1)
	c <- true
}
```

??? example "prints"
    ```text
    Program took 1.00006811s
    ```

Send and receive on channels are both blocking. Go will block the enclosing goroutine until both sides of the channel are ready. 

You can specify whether a channel is to be used for sending or receiving as part of the function parameter list:
```go
func main() {
	randInts := make(chan int)
	done := make(chan bool)

	go roll1d6(randInts)

	go printRoll(randInts, done)
	<-done
}

func roll1d6(c chan<- int) {
	c <- rand.Intn(6) + 1
}

func printRoll(c <-chan int, done chan<- bool) {
	fmt.Println("You rolled a", <-c)
	done <- true
}
```

A channel can have any number of senders and receivers. A value going into a channel is guaranteed to be received by only one receiver.

## Hands on!
1. In the repo, open the file `./intermediate/12channels.go`
2. Complete the TODOs
3. Run `make 12` from project root (alternatively, type `go run ./12channels.go`)
4. Example implementation available on `solutions` branch
