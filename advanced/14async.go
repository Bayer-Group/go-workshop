package main

import (
	"fmt"
	"math/rand"
	"time"
)

// sends a random value for the given color every 5-25ms
func beam(color string, subpixel chan<- string) {
	for {
		time.Sleep(time.Millisecond * time.Duration(rand.Intn(21)+5))
		subpixel <- fmt.Sprintf("%v%v", color, rand.Intn(257))
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	r := make(chan string)
	g := make(chan string)
	b := make(chan string)

	go beam("R", r)
	go beam("G", g)
	go beam("B", b)

	crossedBeams := make(chan string, 6)

	// TODO write a looping `select` that accepts the first six items to appear across all three channels `r`, `g`, and `b` and
	//   sends them into the `crossedBeams` channel.

	// TODO close the crossedBeams channel

	// TODO print the contents of the crossedBeams channel using `range`
}
