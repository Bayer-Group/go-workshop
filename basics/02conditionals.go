package main

import "fmt"

func main() {
	// TODO modify `a` so the `if` statement results in printing the "success" message
	a := ""
	b := "pneumonoultramicroscopicsilicovolcanoconiosis"
	if len(a) > len(b) {
		fmt.Printf("Success! %v bows down to the mighty %v\n", b, a)
	} else {
		fmt.Printf("Failure! %v isn't quite as long as %v\n", a, b)
	}

	sumA := sumRunes(a) // returns some int
	// TODO add an `if/else` that prints a "success" message if `sumA` is greater than `sumRunes(b)`
	//  	and otherwise prints a "failure" message
	fmt.Println(sumA)

	// TODO rewrite the `if/else` statement from line 9 as a `switch` statement

	var whatAmI interface{} = "I am a string but I don't know it"
	// TODO add a case to the below switch statement to print `whatAmI is a string`
	switch t := whatAmI.(type) {
	default:
		fmt.Printf("Unknown type: %T\n", t)
	}
}

// this function can be ignored for this exercise
func sumRunes(str string) (output int) {
	for _, r := range str {
		output += int(r)
	}
	return
}
