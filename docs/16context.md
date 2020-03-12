Covered in this module:

* context
* strconv

## context
A common first parameter you will see for many functions is `context` often abbreviated to `ctx`. Context is a mechanism to pass cancellation signals and deadlines, and can also pass values scoped to some portion of the application (such as a module, individual request, etc).

A context will either be `context.Background()` which is the base level context, or will be a child of the `context.Background()` created via functions like `context.WithCancel()` or `context.WithDeadline()`.
```go
package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	reportCount := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())
	go worker(ctx, reportCount)

	time.Sleep(time.Millisecond * 100)
	cancel()
	fmt.Println(<-reportCount)
}

func worker(ctx context.Context, reportCount chan<- int) {
	var count int
	for ctx.Err() == nil {
		count++
	}
	reportCount <- count
}
```

??? "prints"
    ```text
    4393918
    ```

You can use the cancel function to bail out of some unit of work, based on an error condition or another trigger. 

Deadlines and timeouts enable you can set a timeout on requests to dependencies scoped to an incoming request or some portion of your application.

Context state can be access synchronously with `ctx.Err()` which returns an error, or asynchronously with `ctx.Done()` which returns a closed channel, when the context is cancelled or the deadline is reached.
```go
func worker (ctx context.Context, increment <-chan bool) (count int) {
	for {
		select {
		case <-ctx.Done():
			return
		case <-increment:
			count++
		}
	}
}
```

Context signals are unidirectional. A cancelled context will also cancel any child context. You can use context to gracefully stop your application or prematurely stop calls to dependencies. However, you must use something like a channel or a waitgroup to receive done signals from your goroutines if you want to make sure all parts of your application are properly closed before stopping the app.

##strconv
In a data-heavy application there are occasions where you will need to make conversions between various units. Go has several build in libraries built for converting and manipulating certain types.

`strconv` contains functions like `strconv.Atoi()` and `strconv.Itoa()` for converting strings to integers and vice versa:
```go
package main

import (
	"fmt"
	"strconv"

	"github.com/pkg/errors"
)

func main() {
	var str string
	str = strconv.Itoa(256)
	fmt.Println(str)

	var backsies int
	var err error
	backsies, err = strconv.Atoi(str)
	if err != nil {
		fmt.Println(errors.Wrap(err, "Atoi conversion failed"))
	}

	fmt.Println(backsies)
}
```
