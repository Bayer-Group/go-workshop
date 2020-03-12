Covered in this module:

* packages
* scope
* imports

## packages
Packages allow you to group related code by domain, function, or other logical grouping. 

In practice, a package is a directory with one or more `.go` files. Packages are determined by the `package` declaration at the top of each file. You cannot have different package declarations in two `.go` files in the same directory.

As you may have already noticed, you can call functions from other packages with the syntax `package.Function`.

## scope
Scoping functions, values, and types is controlled simply by the casing on the first letter of the item. 

Capitalized items are visible and usable within the package they are declared in as well as any package that imports them. Lowercase items are only visible within the package they are declared and can be considered __private__.

This can be very useful to control access to certain properties on your struct if you want them to be modifiable only by the struct's methods and not from the invoking package:
```go
type MyCoolStruct struct {
	internal string
	Value string
}
```

A common pattern is to provide a `New()` function that returns your struct that may contain hidden properties that were set as part of the creation. This allows you to encapsulate the logic for initializing your struct to your package rather than relying on the caller to do it properly:
```go
package caller

type Caller struct {
	callCount int
	call func(...interface{}) (interface{}, error)
}

func (d *Caller) Call(args ...interface{}) (interface{}, error) {
	d.callCount++
	return d.call(args)
}

func New(call func(...interface{}) (interface{}, error)) *Caller {
	return &Caller{call: call}
}
```

The above function can be invoked from another package as `caller.New`. Package and function names become important here. When naming your functions, keep the package name in mind and avoid unnecessary stuttering. `New` is preferable to `NewCaller` in this case as `NewCaller` would result in stuttering on invocation: `caller.NewCaller()`. 

## imports
You can import non-native Go libraries by specifying the repo location in the import declaration.
```go
package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
)
```

To make the library available within your codebase, use the command `go get -u`:
```bash
go get -u github.com/sirupsen/logrus
```

Go Modules are Go's built-in dependency management. They are not covered in-depth in this workshop but check out [this tutorial](https://blog.golang.org/using-go-modules) to get started. Projects utilizing Go Modules (like this one) will have a `go.mod` and `go.sum` file at root. For those repos `go get -u` command will automatically add the dependency to your `go.mod` file.

Just like packages within your project only uppercased functions, types, etc from the imported packages are accessible within your code.

## Hands on!
1. In the repo, open the file `./advanced/15packages.go`
2. Complete the TODOs
3. Run `make 15` from project root (alternatively, type `go run ./15packages.go`)
4. Example implementation available on `solutions` branch
