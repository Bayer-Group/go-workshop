package main

import (
	"fmt"
	"strings"
)

func main() {
	sequence := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight"}

	for ind := range sequence {
		addThree(&sequence[ind]) // TODO replace this call with a call to addTwo
	}

	fmt.Println("count eight 3/4 bars:", strings.Join(sequence, ", "))
}

// TODO add a function called addTwo that takes a pointer string and appends ", Two" to the value referenced by
//   the pointer and then calls addThree.
//   expected output of main above will be "count eight 3/4 bars: One, Two, Three, Two, Two, Three..."

func addThree(input *string) {
	if input != nil {
		*input = *input + ", Three"
	}
}
