Covered in this module:

* functions
* variadic functions

## functions
We've been using functions for almost the whole workshop so far, but let's look take a closer look. Go Functions require an explicit parameter list and return type. The function name comes immediately after the `func` keyword and then a comma-space separated list of parameters wrapped in parentheses. Each parameter has a name and then a type separated by a space. The return type is specified after the parameter list followed by an enclosure `{}`. You must also explicitly return from your function if it has a return type. The returned value must be the same type as the declared return type.
```go
func max(a int, b int) int {
	if a >= b {
		return a
    }
	return b
}
```

Functions without a return type do not need a return and will run through all statements:
```go
func print(message string) {
	fmt.Println(message)
}
```

You invoke a function with the parameters wrapped in parentheses:
```go
package main

import "fmt"

func main() {
	prnt("something important")
}

func prnt(message string) {
	fmt.Println(message)
}
```

Functions can have multiple returned values, comma separated and wrapped in parentheses. Return types are still enforced and `return` must have the same number of values:
```go
func negativePositive(i float64) (float64, float64) {
	return -math.Abs(i), math.Abs(i)
}
```

`return` immediately stops the function. This function will return `goodbye` and never print `"unreachable code"`.
```go
func signOff() string {
	return "goodbye"
	
	fmt.Println("unreachable code")
	
	return "farewell"
}
```

### named return values
Return values can be named. The values are initialized as zero-valued when the function is invoked.  A _naked_ return will return the named return values in their current state:
```go
func concat(strs ...string) (output string) {
	for _, str := range strs {
		output += str
    }
	return
}
```

The named returned values can be overridden by explicitly returning values of the correct type:
```go
func concat(strs ...string) (output string) {
	for _, str := range strs {
		output += str
    }
	if len(output) >= 21 {
		return
	}
    return "real output too short"
}
```

### variadic functions
The last parameter of a function can take any number of arguments of the same type:
```go
func concat(strs ...string) string {
	var output string
	for _, str := range strs {
		output += str
    }
	return output
}
```

### capture returned values
When calling a function, you capture the returned values into variables:
```go
package main

import "fmt"

func main() {
	strings1 := []string{"A","E","I","O","U","sometimes","Y"}
	strings2 := []string{"E","B","N","/","O","Z","N"}
	
	title := concat(strings1...)
	
	var artist string
	artist = concat(strings2...)
	
	fmt.Println(title, artist)
}

func concat(strs ...string) (output string) {
	for _, str := range strs {
		output += str
    }
	return
}
```

All values returned from a function must be captured or explicitly ignored using `_`
```go
func main() {
    first, second := someStrings()

    third, _ := someStrings()

    fourth := someStrings() // compiler throws error here
}

func someStrings() (string, string) {
    // snippet
}
```

Functions can return many values as a tuple, but you sacrifice readability especially in a case where the returned values are not named:
```go
func main() {
    year, _, day, hour, min, _ := getDate()
    
    fmt.Println(year, day, hour, min)
}

func getDate() (int, int, int, int, int, int) {
    // snippet
    return year, month, day, hour, min, sec
}
```

!!! note
    There may be some instances where it is warranted, but a good rule of thumb is try to avoid return value lists over three values. Part 9 will introduce the `struct` type for collecting data into logical groupings.

## Hands on!
1. In the repo, open the file `./basics/05functions.go`
2. Complete the TODOs
3. Run `make 05` from project root (alternatively, type `go run ./05functions.go`)
4. Example implementation available on `solutions` branch
