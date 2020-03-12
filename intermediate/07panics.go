package main

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())

	// TODO add a deferred recover function that checks the type of the message in the panic and
	//   returns a unique message for string, error, and then any other type
	// HINT value.(type) returns the type of an interface

	fmt.Println(couldPanic())
}

func couldPanic() string {
	switch rand.Intn(3) {
	case 0:
		panic("string")
	case 1:
		panic(errors.New("error"))
	default:
		panic(3.14159265)
	}
	return "didn't panic"
}
