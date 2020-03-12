package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// TODO write a struct `rectangle` that has two int properties: height and width
	rectangle := rectangle{height: 40, width: 100}

	// TODO add `area` and `circumference` methods to your `rectangle` struct that return said values
	fmt.Printf("Area of rectangle is %v units (expected 4000)\n", rectangle.area())
	fmt.Printf("Circumference of rectangle is %v units (expected 280)\n", rectangle.circumference())

	// TODO calculate the final position after a randomized dance routine on a grid with your dancer
	//   starting at 0,0 facing `north` (facing positive Y axis):
	//     1. in `main()` create an instance of dance routine with at least 20 random steps (loop over `newRandomStep()`)
	//     2. add a method to the danceRoutine struct for printing the final position (currentX and currentY)
	//	   3. call `perform()` on your dance routine
	//     4. call the method you added in step 2 to print the final position
}

// danceRoutine - a routine made up of steps that starts from `currentX` and `currentY` coordinate pair.
type danceRoutine struct {
	steps    []step
	currentX int
	currentY int
}

// perform - perform all steps in the dance routine, updating currentX and currentY to the resulting location
func (dr *danceRoutine) perform() {
	for _, s := range dr.steps {
		dr.currentX, dr.currentY = s.do(dr.currentX, dr.currentY)
	}
}

// newRandomStep - creates a new `step` with a random direction and random distance up to 25 inches
func newRandomStep() step {
	return step{dir: (direction)(rand.Intn(4)), distanceInches: rand.Intn(25)}
}

// step - contains the direction of the movement in relation to facing north (positive Y axis)
type step struct {
	dir            direction
	distanceInches int
}

// do - perform the step from a current set of coordinates, outputting the resulting coordinate pair
func (s step) do(currentX int, currentY int) (int, int) {
	switch s.dir {
	case left:
		return currentX - s.distanceInches, currentY
	case right:
		return currentX + s.distanceInches, currentY
	case forward:
		return currentX, currentY + s.distanceInches
	case back:
		return currentX, currentY - s.distanceInches
	}
	return currentX, currentY
}

// direction - enum containing the four cardinal directions for movement
type direction int

const (
	left direction = iota
	right
	forward
	back
)
