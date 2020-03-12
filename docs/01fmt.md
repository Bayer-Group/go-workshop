Covered in this module:

* main
* import
* fmt library
* types
* variables
* constants

## hello world
Entry point to your application is a main function:

```go
package main

import "fmt"

func main() {
	fmt.Println("hello world")
}
```

* Files in the same package share scope
* `import` is the keyword for importing packages
* native packages have no path (no slashes)
* parenthesis can be used to import multiple packages without repeating `import`

```go
package main

import (
	"fmt"
	"math"
)
```

## fmt
Library `fmt`, pronounced __fumpt__, provides many common utilities:

* printing to console
* parameterized strings
* error message creation

### Common `fmt` pattern

| function | behavior |
| --- | --- |
| `Println` | print all provided parameters, newline is automatic |
| `Printf` | first parameter is a format string with placeholders followed by parameters to replace the placeholders.  No auto newline |
| `Print` | print all provided parameters |

### Common `fmt` format placeholders (verbs)

| function | behavior |
| --- | --- |
| `%v` | default format (uses the default for types as below)
| `%d` | for int, base 10 value |
| `%s` | uninterpreted bytes of string |
| `%t` | for bool, the word true or false |

### `fmt.Sprintf`
Returns a string built from the formatting string (first parameter) and a list of operands to replace verbs (e.g. `%v`).

```go
fmt.Sprintf("They're waiting for you %v… in the test chamber.", "Gordon")
```

## Common Primitives
* int
* string
* bool
* byte (alias for uint8)
* rune (alias for int32, represents a Unicode code point)
* float32/float64

## Variables
Keyword `var` can be used to declare one or many variables
```go
var name string
```

You can initialize the variable inline.  Types are inferred.
```go
var name = "Gophee"
```

Go provides the shorthand operator `:=` to declare and initialize a variable
```go
time := "a flat circle"
```

Declaring variables that are not initialized are automatically zero-valued
```go
var (
	name string
	num int
	state bool
)

fmt.Printf("\"%v\" - %v: %v", name, num, state)
```

??? example "prints"
    ```text
    "" - 0: false
    ```

!!! note
    Primitives can not be `nil`, only pointers (pointers are covered in part 8).

## Constants
* can be a character, string, boolean, or numeric
* numeric constants have no inherent type
* can be explicitly cast when used
```go
const a = 9000
fmt.Println(float64(a))
```
* will take a type implicitly when used as a parameter
```go
const (
	a = 9000
	b = 4000
)
fmt.Println(math.Max(a, b))
// math.Max expects two float64 parameters
```

## Hands on!
1. In the repo, open the file `./basics/01fmt.go`
2. Complete the TODOs
3. Run `make 01` from project root (alternatively, type `go run ./01fmt.go`)
4. Example implementation available on `solutions` branch
