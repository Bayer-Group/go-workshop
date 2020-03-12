package main

import (
	"fmt"
	"time"
)

func main() {
	// TODO call `repeat()` as a goroutine and synchronously as many times as you like
	// TODO experiment with removing the `time.Sleep` from `repeat()`

	fmt.Scanln() // this will wait for a keypress before bailing out of the program
}

func repeat(str string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(str, i)
		time.Sleep(time.Microsecond * 5)
	}
}
