Covered in this module:

* interfaces

## interface

An interface in Go is a collection of methods.  Interface declaration is quite similar to structs except you list the methods that the interface has rather than properties. The interface method parameters and return types can be named. The names are optional but may add clarity to the interface. The following two interfaces are equivalent.
```go
type animal interface {
	speak() string
	eat(int) string
	poo(int) string
}
```

```go
type animal interface {
	speak() (message string)
	eat(ounces int) (message string)
	poo(ounces int) (message string)
}
```


Any struct that implements all the methods of an interface is then considered as implementing that interface:
```go
type animal interface {
	speak() string
	eat(int) string
	poo(int) string
}

type dog struct {
	name string
	digestingFoodOunces int
}

func (d dog) speak() string {
	return "rrrr RUFF!"
}

func (d *dog) eat(ounces int) string {
	d.digestingFoodOunces += ounces
	return fmt.Sprintf("%v ate and is now digesting %v ounces of food", d.name, d.digestingFoodOunces)
}

func (d *dog) poo(ounces int) string {
	d.digestingFoodOunces -= ounces
	return fmt.Sprintf("%v poo'd and is now digesting %v ounces of food", d.name, d.digestingFoodOunces)
}
```

Using an interface as a parameter enables a function to accept any struct that implements the interface:
```go
func main() {
	d := dog{name: "Ein"}
	feedAnimal(&d)
}

func feedAnimal(a animal) {
	fmt.Println(a.eat(10))
	fmt.Println(a.speak())
}
```

Because `eat()` and `poo()` have a pointer receiver, only a pointer dog (the type, not the breed) satisfies the interface. The `blackHole` struct below only uses value receivers so it can be passed by value:
```go
type animal interface {
	speak() string
	eat(int) string
	poo(int) string
}

type blackHole struct {
	name string
}

func (b blackHole) speak() string {
	return "..........."
}

func (b blackHole) eat(ounces int) string {
	return fmt.Sprintf("%v ate %v ounces, but it made no difference", b.name, ounces)
}

func (b blackHole) poo(ounces int) string {
	return fmt.Sprintf("%v tried to poo %v but nothing can escape", b.name, ounces)
}

func feedAnimal(a animal) {
	fmt.Println(a.eat(10))
	fmt.Println(a.speak())
}

func main() {
	b := blackHole{name: "Gargantua"}
	feedAnimal(b)
}
```

A struct can have extra methods that aren't part of the interface and a struct can implement any number of interfaces.

## empty interface

The empty interface type `interface{}` can be used to enable a function to take any value. All Go types satisfy the empty interface. `fmt.Println` is a variadic function that takes empty interfaces:
```go
fmt.Println(5 * 5)
fmt.Println("some string")
fmt.Println(func() string { return "hi" })
fmt.Println([]string{"tree", "shrub", "bush", "flower", "leaf", "root"})
```

??? example "prints"
    ```text
    25
    some string
    0x10910c0
    [tree shrub bush flower leaf root]
    ```

Empty interface values must be cast before they can be used as a specific type:
```go
func add(a, b interface{}) interface{} {
	aval, aok := (a).(int)
	bval, bok := (b).(int)

	if aok && bok {
		return aval + bval
	}
	fmt.Printf("%v ok: %v, %v ok: %v\n", a, aok, b, bok)
	return nil
}

func main() {
	fmt.Println(add(5, 4))
	fmt.Println(add("5","4"))
	fmt.Println(add(5.5, 4.4))
	fmt.Println(add("four", "five"))
}
```

??? example "prints"
    ```text
    9
    5 ok: false, 4 ok: false
    <nil>
    5.5 ok: false, 4.4 ok: false
    <nil>
    four ok: false, five ok: false
    <nil>
    ```

If you recall from part 2, you can use `switch` to check the type of an interface and cast it to that type:
```go
var a interface{}
b := 0

switch value := a.(type) {
case string:
    i, _ := strconv.Atoi(value)
    b += i
case bool:
    b = -b
case int:
    b += value
default:
    fmt.Println("unexpected type, b is unchanged")
}
```

## Hands on!
1. In the repo, open the file `./intermediate/10interfaces.go`
2. Complete the TODOs
3. Run `make 10` from project root (alternatively, type `go run ./10interfaces.go`)
4. Example implementation available on `solutions` branch
