package main

import "fmt"

func main() {
	fmt.Println("slices and arrays:")
	axes := [2][]int{
		{0, 1, 2, 3, 4, 5}, // X axis
		{0, 1, 2, 3, 4, 5}, // Y axis
	}

	// TODO use append to add a 7th row and 7th column to the axes.  The numbers you add are your choice

	renderMultiplicationTable(axes)

	fmt.Println()

	// TODO create a map `rainbow` that assigns names to each letter in the color acronym ROYGBIV

	renderRainbow(rainbow)
}

// this function can be ignored for this exercise
func renderMultiplicationTable(axes [2][]int) {
	for _, x := range axes[0] {
		for _, y := range axes[1] {
			fmt.Printf("%v ", x*y)
		}
		fmt.Println()
	}
}

// this function can be ignored for this exercise
func renderRainbow(rainbow map[string]string) {
	const fullBlock = "\u2588"
	const ROYGBIV = "ROYGBIV"

	fmt.Println("The colors Duke! The colors!")

	for _, r := range ROYGBIV {
		fmt.Println(string(r), rainbow[string(r)], fullBlock)
	}

	fmt.Println("I'm colorblind kid")
}
