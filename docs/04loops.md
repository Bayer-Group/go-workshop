Covered in this module:

* for
* range

## loops
Go only has a single loop construct - `for`. The Go compiler does allow forever loops:
```go
for {}
```

## for
A condition can be provided to make the loop behave like a `while` loop you would see in other languages:
```go
i := 0
for i < 10 {
	fmt.Println(i)
	i++
}
```

You can also provide before, condition, and after clauses separated by semicolons.
```go
for i := 0; i < 10; i++ {
	fmt.Println(i)
}
```

```go
factors := []int{10, 20, 30, 40, 50}
product := 1

for i := 0; i < len(factors); i++ {
	product *= factors[i]
}
fmt.Println("Product is", product)
```

If you use the before/condition/after loop the semicolons are required but it is *not* required to have something in all three sub-clauses.

__"Wait"__, you say, __"there has to be a more concise way to loop over a slice, right?"__

## range
Yep

```go
factors := []int{10, 20, 30, 40, 50}
product := 1

for _, factor := range factors {
	product *= factor
}
fmt.Println("Product is", product)
```

Range returns two parameters for iterable types: the index, and the value.

The `for` loop will automatically loop over all elements of the provided iterable:
```go
factors := []int{10, 20, 30, 40, 50}
product := 1

for ind, factor := range factors {
	fmt.Println("I factored index", ind)
	product *= factor
}
fmt.Println("Product is", product)
```

??? example "prints"
    ```text
    I factored index 0
    I factored index 1
    I factored index 2
    I factored index 3
    I factored index 4
    Product is 12000000
    ```

As you can see, using `range` with a slice yields index and value. It does the same with a string, but the values aren’t characters. They are runes which is an alias for the `int32` type, each representing a unicode code point:
```go
str := "my cool string"

for ind, r := range str {
	fmt.Printf("Index %v is the rune %v\n", ind, r)
}
```

??? example "prints"
    ```text
    Index 0 is the rune 109
    Index 1 is the rune 121
    Index 2 is the rune 32
    Index 3 is the rune 99
    Index 4 is the rune 111
    Index 5 is the rune 111
    Index 6 is the rune 108
    Index 7 is the rune 32
    Index 8 is the rune 115
    Index 9 is the rune 116
    Index 10 is the rune 114
    Index 11 is the rune 105
    Index 12 is the rune 110
    Index 13 is the rune 103
    ```

You can use `string(<rune value>)` to cast a rune to a string or use the fmt verb `%c` in your format string

!!! note "advanced note"
    The Go blog has a great post that explains the differences between [strings, bytes, runes, and characters in Go](https://blog.golang.org/strings)

You can also `range` over maps. For maps, `range` yields the keys and values:
```go
m := map[string]string{"A": "Alfa", "B": "Bravo", "C": "Charlie"}

for key, value := range m {
	fmt.Printf("Key %v holds value %v\n", key, value)
}
```

??? example occasionally "prints"
    ```text
    Key C holds value Charlie
    Key A holds value Alfa
    Key B holds value Bravo
    ```

If you don’t need the key/index, you can throw it out by replacing with underscore `_`. If you don’t need the value, you can omit it and the comma-space. Its important to remember that maps are unordered, so if you are transforming a map to a slice or printing from a map the keys will be returned in an unpredictable order!

## Hands on!
1. In the repo, open the file `./basics/04loops.go`
2. Complete the TODOs
3. Run `make 04` from project root (alternatively, type `go run ./04loops.go`)
4. Example implementation available on `solutions` branch
