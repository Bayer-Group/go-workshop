Covered in this module:

* structs
* methods

## structs
Go structs are a way of collecting related items into an object/record:
```go
type character struct {
	name string
	class string
	level int
}

func main() {
	char := character{name: "Helmut Van Kaiser", class: "Fighter", level: 4}
	fmt.Println(char)
}
```

Struct properties are mutable and can be accessed with dot notation:
```go
char := character{name: "Helmut Van Kaiser", class: "Fighter", level: 4}
fmt.Println(char.name)
fmt.Println(char.class)
fmt.Println(char.level)
char.level++
fmt.Println(char.level)
```

??? example "prints"
    ```text
    Helmut Van Kaiser
    Fighter
    4
    5
    ```

Struct properties are not required at initialization and are zero-valued by default:
```go
char := character{name: "Helmut Van Kaiser"}
fmt.Println(char.name)
fmt.Println(char.class)
fmt.Println(char.level)
```

??? example "prints"
    ```text
    Helmut Van Kaiser
    
    0
    ```

Passing a struct as a pointer avoids making another copy:
```go
func main() {
	char := character{name: "Helmut Van Kaiser", class: "Fighter", level: 4}
	fmt.Println(char.name)
	fmt.Println(char.class)
	fmt.Println(char.level)
	levelUp(&char)
	fmt.Println(char.level)
}

func levelUp(char *character) {
	char.level++
}
```

??? example "prints"
    ```text
    Helmut Van Kaiser
    Fighter
    4
    5
    ```

## methods
Structs can have methods that have access to struct properties:
```go
type character struct {
	name string
	class string
	level int
}

func (char *character) levelUp() {
	char.level++
}

func main() {
	char := character{name: "Helmut Van Kaiser", class: "Fighter", level: 4}
	fmt.Println(char.name)
	fmt.Println(char.class)
	fmt.Println(char.level)
	char.levelUp()
	fmt.Println(char.level)
}
```

??? example "prints"
    ```text
    Helmut Van Kaiser
    Fighter
    4
    5
    ```

### pointer vs value receiver
You can pass your struct by value or by reference to your method. If your method needs to modify the struct in some way, it must be passed as a pointer.  
```go
type character struct {
	name string
	class string
	level int
}

func (char *character) levelUp() { // pointer receiver
	char.level++
}

func (char character) card() string { // value receiver
	return fmt.Sprintf("Player %v is a level %v %v", char.name, char.level, char.class)
}

func main() {
	char := character{name: "Helmut Van Kaiser", class: "Fighter", level: 4}
    fmt.Println(char.card())
	char.levelUp()
	fmt.Println(char.card())
}
```

!!! note
    By convention, receiver names should be related to the struct in some way. A receiver for `character` struct might be `c` or `char` but never `self` or `this`. All methods for the same struct should share the same receiver name, regardless of if they are pointer or value receivers.

The compiler does **not** catch cases where your method modifies your struct passed by value:
```go
type character struct {
	name string
	class string
	level int
}

func (char character) levelUp() { // value receiver
	char.level++ // level++ modifies the struct copy
}

func main() {
	char := character{name: "Helmut Van Kaiser", class: "Fighter", level: 4}
	fmt.Println(char.name)
	fmt.Println(char.class)
	fmt.Println(char.level)
	char.levelUp()
	fmt.Println(char.level)
}
```

??? example "prints"
    ```text
    Helmut Van Kaiser
    Fighter
    4
    4
    ```

!!! note
    There are linters such as [revive](https://github.com/mgechev/revive) available to help catch cases like this, which can be tricky to notice manually.

## Hands on!
1. In the repo, open the file `./intermediate/09structs.go`
2. Complete the TODOs
3. Run `make 09` from project root (alternatively, type `go run ./09structs.go`)
4. Example implementation available on `solutions` branch
