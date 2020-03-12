package main

import "fmt"

func main() {
	max := 10
	product := 1
	// TODO write a `for` loop to calculate the factorial of `max` and store the result in `product`
	// HINT Factorial `n!` is the product of all positive integers less than or equal to `n`

	fmt.Printf("Factorial of %v is %v\n", max, product)
	fmt.Println()

	rules := map[string]bool{"dodge": true, "deflect": false, "duck": true, "detect": false, "dip": true, "divvy": false, "dive": true, "dice": false}
	var validRules []string
	// TODO write a `for range` loop that adds the key to `validRules` if the value is `true`

	fmt.Println("Valid rules are", validRules)
}
