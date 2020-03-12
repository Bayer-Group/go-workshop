Covered in this module:

* defer
* panic
* recover

## defer
Any function passed to defer will be run after the containing function completes:
```go
package main

import "fmt"

func main() {
	defer fmt.Println("Clubs")
	defer fmt.Println("Diamonds")
	
	fmt.Println("Hearts")
	fmt.Println("Spades")
}
```

??? example "prints"
    ```text
    Hearts
    Spades
    Diamonds
    Clubs
    ```

Defers are invoked in reverse order (LIFO, last in first out). One common use-case is to defer teardown of resources so they are cleaned up early rather than having to wait on the garbage collector **or** for cases where the GC may not properly clean up a resource:
```go
db, err := sql.Open("postgres", connectionString)
if err != nil {
    return fmt.Errorf("sql.Open: %w", err)
}
defer func() {
    _ = db.Close()
}()

rows, err := db.Query("select * from users")
if err != nil {
    return fmt.Errorf("db.Query: %w", err)
}
defer func() {
    _ = rows.Close()
}()
```

!!! note
    We don't defer close on the `db` or `rows` until after the error check. There is no need to close a resource that didn't get initialized (indicated by the not-nil error). The LIFO behavior of defers here makes it so `rows` is closed first and then `db`. You wouldn't want to close the `db` until the `rows` resource is properly closed.

Return values can be modified within the deferred function:
```go
func main() {
    _, err := convert("two")
    fmt.Println(err)
}

func convert(input string) (output int, errs error) {
	defer func() {
		if errs != nil {
			errs = fmt.Errorf("ya burnt! %v", errs.Error())
        }
    }()
	
	return strconv.Atoi(input) // attempts to convert string to int
}
```

??? example "prints"
    ```text
    ya burnt! strconv.Atoi: parsing "two": invalid syntax
    ```

## panic
Unlike errors which are handled like any other type, `panic` triggers an exception:
```go
func main() {
	for {
		doNotPanic()
    }
}

func doNotPanic() {
	panic("I told you not to panic")
}
```

The panic function accepts a value of any type.

## recover
Panics can be recovered by a deferred function:
```go
func main() {
	defer fmt.Println("End of line, man...")
	
	defer func() {
		recovered := recover()
		if recovered != nil {
			fmt.Println("Recovered from panic:", recovered)
		}
	}()
	
	result := mustConvert("two")
	fmt.Println("Result of conversion:", result)
}

func mustConvert(input string) int {
	output, err := strconv.Atoi(input)
	if err != nil {
		panic(err)
	}
	return output
}
```

??? example "prints"
    ```text
    Recovered from panic: strconv.Atoi: parsing "two": invalid syntax
    End of line, man...
    ```

Avoid using panic in your production code except for specialized use-cases. Panics are meant for unrecoverable errors. As with exception states in other languages, if you do not recover a panic it will kill your app. One valid use-case would be on app start-up if a component fails to initialize.

If you do have a function that could panic, the recommendation is to prepend the name with `must` (e.g. `mustConvert`). That indicates to anyone who uses that function that it could panic.

Panics will occur for various reasons in compiled Go code, such as trying to index a value out of bounds in a slice or if you dereference a nil pointer. Pointers are covered in the next section.

## Hands on!
1. In the repo, open the file `./intermediate/07panics.go`
2. Complete the TODOs
3. Run `make 07` from project root (alternatively, type `go run ./07panics.go`)
4. Example implementation available on `solutions` branch
