# go-workshop ([https://MonsantoCo.github.io/go-workshop/](https://MonsantoCo.github.io/go-workshop/))
Structured hands-on workshop for learning Go. 

## Workshop Outline
The 'Basics' module starts with the components that you will see across most commonly-used languages then delves into 'Intermediate' with the concepts and structures that set Go apart.  The 'Advanced' module builds on those concepts and introduces several new built-in components.

The intended flow of the workshop is to go through each module's presentation on the [GitHub pages](https://MonsantoCo.github.io/go-workshop/) one part at a time and complete the TODOs in the associated hands-on exercises.  Each part has its own make command to simplify running the code.

### Basics - common language structures, the Go way

* 01fmt - fmt-y Dumpty
    * main
    * import
    * fmt library
    * types
    * variables
    * constants
    
* 02conditionals - On one conditional
    * if/else
    * switch
    
* 03slices - A healthy slice of arrays
    * arrays
    * slices
    * make
    * append
    * maps
    * delete

* 04loops - Stuck in the loop with you
    * for
    * range
    
* 05functions - Becoming a functioning Gopher
    * functions
    * variadic functions
    
* 06errors - I am Error
    * errors
    * error syntax
    
### Intermediate - what sets Go apart
    
* 07panics - Don't panic
    * defer
    * panic
    * recover
    
* 08pointers - Get to the point
    * pointers
    * dereferencing
    
* 09structs - Hardwired to self de-struct
    * structs
    * methods
    
* 10interfaces - Unlike a user interface, these take a bit of explanation
    * interfaces
    
* 11goroutines - Getting into the routine
    * goroutines
    
* 12channels - Now you're thinking with channels
    * channels

### Advanced - Keep 'em coming

* 13currying - Staying anonymous
    * anonymous functions
    * closures
    * currying
    * recursion

* 14async - Getting out of sync
    * buffered channels
    * select
    * non-blocking channels
    * close

* 15packages - Here's your package
    * packages
    * scope
    * imports
    
* 16context - Taken out of context
    * context
    * strconv
    
### Topics not covered (but can be very useful)
__Disregard the order__

* In-depth Type Conversion and casting (strings, numbers, interfaces, etc)
* Mutexes
* SQL
* Unit Testing
* HTTP server
* HTTP client
* Regex
* File manipulation
* Flags and Environment Variables
* Time

## Mac Setup
To install Go for Mac, you have several options:
 
* manually download the installer from [Go Downloads](https://golang.org/dl/)
    * typically installs at `/usr/local/go` on Mac
    
* do `brew install go`.
    * typically installs at `/usr/local/Cellar/go`

As of Go 1.11, Go has built-in dependency management via [Go Modules](https://blog.golang.org/modules2019) so the **GOPATH** is no longer required. You can now place your Go projects anywhere on your file system as long as the project contains a `go.mod` file.

### go mod
[This tutorial](https://blog.golang.org/using-go-modules) is recommended to get going with Go Modules, but here is a quick way to get started:

When spinning up a new Go project or converting an older project (from `dep` for example), you can simply run `go mod init <path>` and Go will initialize a `go.mod` for your project. A `<path>` structure should be origin/org/repo (e.g. `github.com/MonsantoCo/go-workshop`). That path allows other projects to find yours on the web and is especially important in cases where your repo is intended to be a module/import for another project. 

Now running any Go command (e.g. `go build` or `go test`) from within your project will trigger a go module check that will download and attempt to resolve any dependencies in your project. Dependencies both direct and indirect will be noted in `go.mod` and checksums will be saved in `go.sum`. Both files should be committed to your central repo.

For previous `dep` projects, once `go.mod` and `go.sum` are created you can safely delete `Gopkg.toml` and `Gopkg.lock` files as well as `vendor` and `.vendor-new` folders from your project. 

## MkDocs Site
### To preview local changes to the site:
 - Install Python and pip (pip3 comes by default on mac.  just create a new link from pip to pip3 `ln -fs /usr/local/bin/pip3 /usr/local/bin/pip`) 
 - Install `mkdocs` using: `pip install mkdocs`
    - if you are getting certificate validation issues try the instructions below
 - Install the material theme using: `pip install mkdocs-material`
 - From the project root run: `make local-docs`
 
  
### to deploy github pages
- NOTE: you must have PIP installed (try `pip` command.  You can install with `brew install python`)
1. run `make publish`
