Covered in this module:

* goroutines

## goroutines
So far, all the function calls that have been shown are done synchronously. To call a function so that it runs concurrently to the invoker, use the `go` keyword:
```go
package main

import (
	"fmt"
	"time"
)

func main() {
	go repeat("Pete", 3)

	repeat("Repeat", 5)

	time.Sleep(time.Millisecond * 10)
}

func repeat(str string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(str, i)
		time.Sleep(time.Microsecond * 5)
	}
}
```

??? example occasionally "prints"
    ```text
    Repeat 0
    Repeat 1
    Pete 0
    Pete 1
    Repeat 2
    Pete 2
    Repeat 3
    Repeat 4
    ```

!!! note 
    Your program will not wait for your goroutines to complete before closing. Once the goroutine that **main** is running on (usually refered to as the **main** routine) closes the app will close as well.

Only functions with no return values are eligible to be goroutines. There are no such restrictions on goroutine function parameters.

You will likely see a different pattern of printed strings each time you run the `repeat` example function above.

Go handles all the logic for spreading work and swapping routines on CPU threads for you, so you can spin up many goroutines without necessarily needing to think about how the work might be accomplished. 
```go
func main() {
	go ticker(100)
	go ticker(250)
	go ticker(333)
	go ticker(500)

	time.Sleep(time.Millisecond * 1050)
}

func ticker(waitMs int) {
	for i := 1; i <= 1000/waitMs; i++ {
		time.Sleep(time.Millisecond * time.Duration(waitMs))
		fmt.Printf("%v, ", waitMs * i)
	}
	fmt.Println()
}
```

??? example "prints"
    ```text
    100, 200, 250, 300, 333, 400, 500, 500, 500, 600, 666, 700, 750, 800, 900, 1000, 
    1000, 
    999, 
    1000, 
    ```

!!! note
    Go has a ticker implementation in the native package `time` (created by `time.NewTicker()`). The ticker above is just for teaching purposes. 

You cannot return a value from goroutine directly, but go offers a construct called `channel` for communicating between goroutines that will be covered in the next module.

## Hands on!
1. In the repo, open the file `./intermediate/11goroutines.go`
2. Complete the TODOs
3. Run `make 11` from project root (alternatively, type `go run ./11goroutines.go`)
4. Example implementation available on `solutions` branch
