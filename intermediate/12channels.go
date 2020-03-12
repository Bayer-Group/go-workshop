package main

import (
	"fmt"
	"math/rand"
	"time"
)

// food - a type of food and how long it takes to cook said food
type food struct {
	name       string
	prepTimeMs int
}

// cook - a worker routine that makes random foods from the available list indefinitely
func cook(toCustomer chan<- string, availableFoods []food) {
	for {
		itemToCook := availableFoods[rand.Intn(len(availableFoods))]
		time.Sleep(time.Millisecond * (time.Duration)(itemToCook.prepTimeMs))
		toCustomer <- itemToCook.name
	}
}

// newBottomlessPit - creates quantum singularity that can never be sated
func newBottomlessPit() customer {
	return customer{eatenFoods: make(map[string]int)}
}

// customer - a consumer of food
type customer struct {
	eatenFoods map[string]int
}

// printStomachContents - prints the current value of eatenFoods (note that this is not protected by a mutex)
func (c customer) printStomachContents() {
	fmt.Println("The customer ate:")
	for key, value := range c.eatenFoods {
		fmt.Println(value, key)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	availableFoods := []food{{"Steak", 50}, {"Egg", 7}, {"Toast", 5}, {"Chicken Wing", 30}, {"Fry", 20}, {"Cobb Salad", 25}, {"Ice Cream", 120}, {"Popcorn", 10}, {"Mac and Cheese", 15}, {"Taco", 45}}

	// TODO write a `foodStream` method on the `customer` struct that takes a receive channel of type string
	//   for every food name in the channel, add the food name as the key to the `eatenFoods` map with the value being the number of that type of food eaten
	//   HINT: check if key exists to determine whether to insert or just increment value

	// TODO create a `foodStream` string channel

	// TODO write a loop that spins up 5 `cook` goroutines to do the cooking

	bottomlessPit := newBottomlessPit()

	// TODO invoke the `foodStream` method on `bottomlessPit` with the `foodStream` channel

	// let's see how much food your customer eats in 10 seconds
	fmt.Println("The feast goes on...")
	time.Sleep(time.Second * 10)
	bottomlessPit.printStomachContents()
}
