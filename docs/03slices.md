Covered in this module:

* arrays
* slices
* make
* append
* maps
* delete

## arrays
Arrays in Go are an ordered list of elements with a fixed length:
```go
var numbers [6]int

numbers[1] = 800

fmt.Println(numbers[1])
```

Arrays are initialized as zero-valued. The above prints:
```text
[0 800 0 0 0 0]
```

`len()` returns the length of array:
```go
var numbers [6]int

numbers[1] = 800

fmt.Println(len(numbers))
```

??? example "prints"
    ```text
    6
    ```

Arrays can be initialized in-line
```go
numbers := [6]int{4, 8, 15, 16, 23, 42}
```

However, arrays are typically not used in Go, at least not directly...

## slices
On the surface, slices look just like arrays except they are missing a length declaration.
```go
numbers := []int{4, 8, 15, 16, 23, 42}

numbers[1] = 800

fmt.Println(numbers)
fmt.Println(len(numbers))
```

??? example "prints"
    ```text
    [4 800 15 16 23 42]
    6
    ```

Slices are a reference to all or part of an underlying array, hence the name.

### make
The built-in function `make(type, length)` can be used to initialize a slice. The `length` parameter is required:
```go
numbers := make([]int, 3)

numbers[1] = 2
numbers[2] = 3

fmt.Println(numbers)
```

??? example "prints"
    ```text
    [0 2 3]
    ```
    
If you don't know how large your slice will become, you can make a slice with an initial length of 0.

### append
Unlike arrays, slices can change size after initialization. You can `append` elements to the end of a slice.
```go
numbers := make([]int, 3)

numbers[1] = 2
numbers[2] = 3

fmt.Println(numbers) // prints [0 2 3]
fmt.Println(len(numbers)) // prints 3

numbers = append(numbers, 5, 8)

fmt.Println(numbers) // prints [0 2 3 5 8]
fmt.Println(len(numbers)) // prints 5
```

!!! note "advanced note"
    if the existing underlying array is not large enough to accept the appended values, Go will create a new larger array and re-point the slice. Append is safe to use but there is a performance cost.

### slice a slice
You can re-slice a slice using the syntax `slice[begin:end]`

* Begin and end are optional. If either is excluded the slice boundary is used instead
* The value at index `end` is excluded

```go
days := []string{"Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}

fmt.Println(days[:2], "happy days")
fmt.Println(days[2:4], "happy days")
fmt.Println(days[4:], "happy days")
fmt.Println("the weekend comes")
fmt.Println("my cycle hums")
fmt.Println("ready to race to you")
```

??? example "prints"
    ```text
    [Sunday Monday] happy days
    [Tuesday Wednesday] happy days
    [Thursday Friday] happy days
    the weekend comes
    my cycle hums
    ready to race to you
    ```

## maps
Maps are structures of unordered key-value pairs. They are declared with the syntax `map[key-type]value-type`. `make` for maps doesn't need a length parameter
```go
bases := make(map[string]string)

bases["G"] = "Guanine"
bases["T"] = "Thymine"
bases["A"] = "Adenine"
bases["C"] = "Cytosine"

fmt.Println(bases)
```

```go
bases := map[string]string{
	"G": "Guanine",
	"T": "Thymine",
	"A": "Adenine",
	"C": "Cytosine",
}

fmt.Println(bases)
```

When invoking `make` on a map, the optional second parameter is a capacity hint.
```go
bases := make(map[string]string, 4) // capacity hint set at 4

bases["G"] = "Guanine"
bases["T"] = "Thymine"
bases["A"] = "Adenine"
bases["C"] = "Cytosine"

fmt.Println(bases)
```

!!! note "advanced note"
    You can use the capacity hint to avoid extra memory allocations on the heap if you have a good idea what capacity the map will need during its lifetime.

If it's possible for your key to not exist in the map, you can verify that the key exists by checking the bool returned as the second value when accessing the map by key:
```go
	bases := map[string]string{
		"G": "Guanine",
		"T": "Thymine",
		"A": "Adenine",
		"C": "Cytosine",
	}

	maybeBase, exists := bases["U"]

	fmt.Println(exists) // prints false
	fmt.Println(maybeBase) // prints "" (zero-valued string since key didn't exist)
```
You can use this to gracefully react to scenarios where a key doesn't exist in the map.

### delete
Values can be deleted from maps by their key with `delete(mapName, key)`.

```go
bases := map[string]string{
	"G": "Guanine",
	"T": "Thymine",
	"A": "Adenine",
	"C": "Cytosine",
}

delete(bases, "G")

fmt.Println(bases)
```

??? example "prints"
    ```text
    map[T:Thymine A:Adenine C:Cytosine]
    ```

## Hands on!
1. In the repo, open the file `./basics/03slices.go`
2. Complete the TODOs
3. Run `make 03` from project root (alternatively, type `go run ./03slices.go`)
4. Example implementation available on `solutions` branch
