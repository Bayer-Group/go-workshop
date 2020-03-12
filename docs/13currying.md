Covered in this module:

* anonymous functions
* closures
* currying
* recursion

## anonymous functions
Go supports anonymous functions that can be passed around like any other value.
```go
func main() {
	add := func(a, b int) int {
		return a + b
	}

	fmt.Println(add(3, 400))
	fmt.Println(add(10, 59))
	fmt.Println(add(-5, 252))
	fmt.Println(add(80, -311))
	fmt.Println(add(-123, 42))
	fmt.Println(add(0, -65))
}
```

??? example "prints"
    ```text
    403
    69
    247
    -231
    -81
    -65
    ```
    
## currying
Functions can return a function to allow for currying. The function returned by `sequence()` is considered a closure since it encapsulates the sequence value. The returned function essentially takes ownership of the `seq` variable.
```go
func sequence(startingVal int) func() int {
	seq := startingVal
	return func() int {
		seq++
		return seq
	}
}

func main() {
	nextVal := sequence(25)

	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())
	fmt.Println(nextVal())

	moreVal := sequence(0)

	fmt.Println(moreVal())
	fmt.Println(moreVal())
}
```

??? example "prints"
    ```text
    26
    27
    28
    29
    1
    2
    ```

## recursion
Functions can also invoke themselves, allowing for recursive behavior.

```go
package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func mockCall() (int, error) {
	num := rand.Intn(10)
	if num <= 5 {
		return num, errors.New("call failed")
	}
	return num, nil
}


func attemptCall(call func() (int, error), maxRetries, try int) int {
	result, err := call()
	if err != nil {
		if try >= maxRetries {
			fmt.Printf("call failed after %v attempts. last error is: %v\n", maxRetries, err.Error())
			return 0
		}
		return attemptCall(call, maxRetries, try + 1)
	}
	return result
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	result := attemptCall(mockCall, 3, 1)
	fmt.Println("Result is", result)
}

```

??? example "occasionally prints"
    ```text
    call failed after 3 attempts. last error is call failed
    Result is 0
    ```

## Hands on!
1. In the repo, open the file `./advanced/13currying.go`
2. Complete the TODOs
3. Run `make 13` from project root (alternatively, type `go run ./13currying.go`)
4. Example implementation available on `solutions` branch
