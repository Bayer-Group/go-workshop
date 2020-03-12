Covered in this module:

* errors
* error syntax

## errors
Go provides a built-in interface called `error` that has a single method `Error()` that returns the error message string. (interfaces will be covered later).

Errors can be created with `errors.New()`:
```go
err := errors.New("I am Error")

fmt.Printf("%v - Press any key to restart", err.Error())
```

??? example "prints"
    ```text
    I am Error - Press any key to restart
    ```

`fmt.Errorf()` can be used to create error text at runtime (as opposed to `errors.New(fmt.Sprintf())`:
```go
number := -10
var err error

if number < 0 {
	err = fmt.Errorf("%v is not positive", number)
}
```

### error syntax
By convention, any function that can return an error has error as its last return value.

`nil` can be returned in place of an error:
```go
func bePositive(number int) (string, error) {
	if number < 0 {
		return "", fmt.Errorf("%v is not positive", number)
    }
	return "I'm so happy", nil
}
```

A common structure that arises in Go code is:

* Invoke a function that can return an error
* React to error if its not nil
* Continue normal operation

```go
number := -10

tidbit, err := bePositive(number)
if err != nil {
	fmt.Println("Error encountered:", err.Error())
}

fmt.Println(tidbit)
```

### error wrapping
`fmt.Errorf` has a special verb `%w` (since 1.13) that enables generation of a stack-trace-like structure and saves manually calling `err.Error()` to properly print error strings. `%w` is only available for `fmt.Errorf()`.

```go
func getDescription(id int) (string, error) {
    if desc, err := getDescriptionByID(id); err != nil {
        return "", fmt.Errorf("getDescriptionByID: %w", err)
    } else {
        return desc, nil
    }
}
```

Printing the returned error gives:
```text
getDescriptionByID: id doesn't exist in database
```
    
A common pattern is to indicate which function was invoked that yielded the error, so that you can pinpoint the issue from the logged errors. By convention, errors shouldn't begin with a capital letter or end in punctuation. Steps in the 'stack-trace' are indicated by `:`

If a function only returns an `error`, the compiler does allow you to skip capturing the returned value, but by convention you should explicitly ignore it with `_`.
```go
func main() {
    mightErrDontCare() // allowed by the compiler
    
    _ = mightErrDontCare() // functionally the same, but increases readability
}

func mightErrDontCare() error {
    // snippet
}
```

### errors as a construct
Errors don’t have any special behavior:

* they are just like any other return value
* It’s up to the function caller to handle them how they like
* Go does have an exception state called `panic` that will be covered in the next section.

## Hands on!
1. In the repo, open the file `./basics/06errors.go`
2. Complete the TODOs
3. Run `make 06` from project root (alternatively, type `go run ./06errors.go`)
4. Example implementation available on `solutions` branch
