package main

import "fmt"

func main() {
	// TODO write a function `factorial()` that uses recursion to return the factorial of the provided integers below
	// HINT Factorial `n!` is the product of all positive integers less than or equal to `n`

	fmt.Println(factorial(1)) // expect 1
	fmt.Println(factorial(5)) // expect 120
	fmt.Println(factorial(8)) // expect 40320
}
