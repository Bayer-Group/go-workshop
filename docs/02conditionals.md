Covered in this module:

* if/else
* switch

## if/else
These have a very simple structure in Go.

* parentheses are omitted
* curly braces are required
* related `if/else` clauses continue on the same line as the last closing brace
```go
i := 5
if i < 0 {
	fmt.Println("i is negative")
} else if i > 0 {
	fmt.Println("i is positive")
} else {
	fmt.Println("i is 0")
}
``` 

!!! note
    Go does not have an analog for the ternary operator!

A statement can precede the condition of an `if` clause. Variables declared in-line are only available within the `if/else` block
```go
if i := 5; i < 0 {
	fmt.Println("i is negative")
} else if i > 0 {
	fmt.Println("i is positive")
} else {
	fmt.Println("i is 0")
}

// referencing `i` here would fail (this would be caught at compilation)
```

## switch
Switch can be used as an alternative to an `if/else`:
```go
i := 5
switch {
case i < 0:
	fmt.Println("i is negative")
case i > 0:
	fmt.Println("i is positive")
default:
	fmt.Println("i is 0")
}
```

Switch can be used for pattern matching. Cases that use the same result can be comma separated. Cases without a result are explicitly skipped.
```go
i := "Monday"
switch i {
case "Monday":
	fmt.Println("Wake up, it’s a beautiful Monday")
case "Tuesday":
    fmt.Println("Hey everybody it’s Tuesday!")
case "Wednesday", "Thursday": // these two cases get the same result
    fmt.Println("Almost Friday")
case "Friday": // skipped
default:
	fmt.Printf("%v is not the best day\n", i)
}
```

Default is not required but is recommended in cases where your pattern match is not exhaustive. If the `switch` matches no cases the code will move forward.

Switch can be used to check the type of an interface (interfaces are covered in part 10):
```go
var i interface{}

switch i.(type) {
case string:
	fmt.Println("i is a string")
case bool:
	fmt.Println("i is a boolean")
case int:
	fmt.Println("i is an int")
default:
	fmt.Println("unexpected type")
}
```

## Hands on!
1. In the repo, open the file `./basics/02conditionals.go`
2. Complete the TODOs
3. Run `make 02` from project root (alternatively, type `go run ./02conditionals.go`)
4. Example implementation available on `solutions` branch
