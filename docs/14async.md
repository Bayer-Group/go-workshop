Covered in this module:

* buffered channels
* select
* non-blocking channels
* close

## buffered channels
As discussed in [Part 12](12channels.md), channels by default are synchronous in that they require that both sender and receiver are ready before the value is sent across the channel. Channels can be buffered to enable values to be sent into the channel without a ready receiver.
```go
func main() {
	buffered := make(chan string, 3)

	buffered <- "first"
	buffered <- "second"
	buffered <- "third"

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
}
```

??? example "prints"
    ```text
    first
    second
    third
    ```

However, once the buffered channel is full, any new send will block until the channel has space.
```go
func main() {
	buffered := make(chan string, 2) // decrease buffer to 2

	buffered <- "first"
	buffered <- "second"
	buffered <- "third" // this is now blocking

	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
	fmt.Println(<-buffered)
}
```

## select
Select allows your routine to listen to several channels simultaneously and shares syntax with `switch` except that each case monitors a channel:
```go
func waitToRespond(waitTimeMs time.Duration, response string, respond chan<- string) {
	time.Sleep(time.Millisecond * waitTimeMs)
	respond <- response
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go waitToRespond(500, "first", chan1)

	go waitToRespond(250, "second", chan2)

	for i := 0; i < 2; i++ {
		select {
		case res := <-chan1:
			fmt.Println(res)
		case res := <-chan2:
			fmt.Println(res)
		}
	}
}
```

??? example "prints"
    ```text
    second
    first
    ```

!!! note 
    If two or more channels have messages ready to be received by a single `select`, one case is chosen at random.

## non-blocking channels
The `select` example above behaves in a synchronous manner as it will block until one of its cases is resolved. Adding a `default` case to the `select` will make it non-blocking:
```go
func waitToRespond(waitTimeMs time.Duration, response string, respond chan<- string) {
	time.Sleep(time.Millisecond * waitTimeMs)
	respond <- response
}

func main() {
	chan1 := make(chan string)
	chan2 := make(chan string)

	go waitToRespond(500, "first", chan1)

	go waitToRespond(250, "second", chan2)

	var i int
	for i < 2 {
		select {
		case res := <-chan1:
			fmt.Println(res)
			i++
		case res := <-chan2:
			fmt.Println(res)
			i++
		default:
			fmt.Println("nothing. waiting 50ms")
			time.Sleep(time.Millisecond * 50)
		}
	}
}
```

??? example "prints"
    ```text
    nothing. waiting 50ms
    nothing. waiting 50ms
    nothing. waiting 50ms
    nothing. waiting 50ms
    nothing. waiting 50ms
    second
    nothing. waiting 50ms
    nothing. waiting 50ms
    nothing. waiting 50ms
    nothing. waiting 50ms
    nothing. waiting 50ms
    first
    ```

!!! note
    For cases where you need a looping `select`, it is recommended to omit a default case unless your default is putting the loop to sleep for a period of time. `select` does a lazy block whereas `for` will continue to hog CPU.  

## close
After your routine is done writing to a channel, it can `close` the channel to indicate to the receiver that no more items will be sent. The convention is that only the writer/sender should close a channel; this is enforced by there being no way to tell whether a channel is closed on the send end.
```go
func ping(tick chan<- time.Time) {
	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()

	for i := 0; i < 4; i++{
		tick <- <-ticker.C
	}
	close(tick)
}

func main() {
	tick := make(chan time.Time, 4)

	go ping(tick)

	time.Sleep(time.Millisecond * 1050)

	for t := range tick {
		fmt.Println("tick", t)
	}
}
```

??? example "prints"
    ```text
    tick 2019-02-22 14:48:28.680445 -0500 EST m=+0.255451520
    tick 2019-02-22 14:48:28.929341 -0500 EST m=+0.504354086
    tick 2019-02-22 14:48:29.180488 -0500 EST m=+0.755507241
    tick 2019-02-22 14:48:29.425965 -0500 EST m=+1.000990704
    ```

You can check if a channel is open on the receiver end by capturing two values on the receive. The first value will be the received item and the second is a boolean which is true if the channel is open.
```go
func ping(tick chan<- time.Time) {
	ticker := time.NewTicker(time.Millisecond * 250)
	defer ticker.Stop()

	for i := 0; i < 4; i++{
		tick <- <-ticker.C
	}
	close(tick)
}

func main() {
	tick := make(chan time.Time, 4)

	go ping(tick)

	for {
		t, open := <- tick
		if open {
			fmt.Println("tick",t)
		} else {
			break
		}
	}
}
```

??? example "prints"
    ```text
    tick 2019-02-22 14:49:24.685916 -0500 EST m=+0.251246062
    tick 2019-02-22 14:49:24.936428 -0500 EST m=+0.501765172
    tick 2019-02-22 14:49:25.188174 -0500 EST m=+0.753517292
    tick 2019-02-22 14:49:25.43667 -0500 EST m=+1.002020070
    ```

!!! note
    The close signal takes up a spot in a buffered channel, so the `close` will block if there is no room in the channel.

`close` works for both buffered and unbuffered channels and is a permanent state. Reading from a closed channel will continue to return false for its second value in perpetuity. You can not 'unclose' a channel. Also, attempting to send a value to a closed channel causes a panic. If you have multiple senders for a single channel, take care to make sure all senders are done before closing the channel. If you close the channel too early, an active sender routine will trigger a panic.

For buffered channels, any values put in the channel prior to the close will still be delivered in the order they were sent before the `close` signal will be received.

It is not required to `close` a channel. Channels can be used indefinitely or reinitialized as needed. `close` is just another way to signal between goroutines. 

## Hands on!
1. In the repo, open the file `./advanced/14async.go`
2. Complete the TODOs
3. Run `make 14` from project root (alternatively, type `go run ./14async.go`)
4. Example implementation available on `solutions` branch
