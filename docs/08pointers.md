Covered in this module:

* pointers
* dereferencing

## pointers
Go defaults to passing parameters by value:
```go
func main() {
	star := "Sol"
	fmt.Println(star)
	
	toUpper(star)
	fmt.Println(star)
}

func toUpper(input string) {
	input = strings.ToUpper(input)
}
```

??? example "prints"
    ```text
    Sol
    Sol
    ```

You may have noticed that `toUpper` didn't have an effect on the output, because it modified its `input` parameter, which was a copy of the string passed in by `main`. `toUpper` would only be able to copy the string initialized in `main` if we pass by reference, via a pointer.

### referencing
You can get the address of a value with the reference syntax:
```go
func main() {
	star := "Sol"
	fmt.Println(star)

	toUpper(&star) // pass reference to star
	fmt.Println(star)
}

func toUpper(input *string) {
	*input = strings.ToUpper(*input) // modify pointed-to value
}
```

??? example "prints"
    ```text
    Sol
    SOL
    ```

The example above passed a reference to `star`, which `toUpper` then dereferences so it can modify the underlying string.

| operator | term        | behavior                            |
| -------- | ----------- | ----------------------------------- |
| `&`      | reference   | get the address of the given value  |
| `*`      | dereference | get the value from the given address|

!!! note
    `*` is used to indicate a type as a pointer (e.g. `*string`) and is also used to dereference pointers.

You can get the address of a value with the reference syntax:
```go
func main() {
	moon := "Luna"
	fmt.Println(moon)
	fmt.Println(&moon)
}
```

??? example "prints"
    ```text
    Luna
    0xc00000e1e0
    ```

!!! note
    Pointer types are referred to verbally by prepending `pointer` to the type. For example, `*string` is a pointer string and `*int` is a pointer int.

### nil
Unlike primitives, a pointer can be nil.
```go
func main() {
	toUpper(nil)
}

func toUpper(input *string) {
	*input = strings.ToUpper(*input)
}
```
This example panics because `strings.ToUpper` is attempting to deference a nil pointer.

It's a good practice to always check `nil`-able values before trying to use them (unless you know for sure that the pointer can never be nil):
```go
func main() {
	toUpper(nil)
}

func toUpper(input *string) {
	if input != nil {
		*input = strings.ToUpper(*input)
    }
}
```

## Hands on!
1. In the repo, open the file `./intermediate/08pointers.go`
2. Complete the TODOs
3. Run `make 08` from project root (alternatively, type `go run ./08pointers.go`)
4. Example implementation available on `solutions` branch
