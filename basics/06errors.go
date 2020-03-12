package main

import (
	"errors"
	"fmt"
)

func main() {
	pairs := [][2]float64{{81, 7}, {47, 9}, {81, 8}, {25, 0}, {56, 0}, {94, 1}, {62, 9}, {28, 4}, {11, 5}, {37, 6}, {95, 6}, {28, 8}, {47, 7}, {87, 8}, {90, 5}, {41, 8}, {87, 1}, {29, 6}, {37, 1}, {85, 6}, {13, 0}, {94, 3}, {33, 7}, {78, 4}, {59, 3}, {57, 1}, {89, 9}, {0, 5}, {88, 8}, {3, 5}, {51, 0}, {5, 6}, {66, 8}, {61, 2}, {83, 6}, {63, 6}, {2, 8}, {47, 4}, {77, 3}, {96, 0}, {23, 3}, {37, 3}, {41, 9}, {33, 3}, {91, 2}, {78, 6}, {46, 7}, {40, 3}, {52, 3}, {5, 8}}
	// TODO call the `divide` function for each pair in `pairs` and print the error if there is one, otherwise print the division result.

	words := []string{"xvlbzgbaic", "mrajwwhthc", "tcuaxhxkqf", "dafplsjfbc", "xoeffrswxp", "ldnjobcsnv", "lgtemapezq", "leqyhyzr", "jjpjzpfrfe", "gmotafeths", "bzrjxawnwe", "krbemfdzdc", "ekxbakjqzl", "cttmttcoan", "atyyinkare", "kjyixjrscc", "tnswynsgru", "ssvmaozfzb", "sbojifqgzs", "nwtksmvoig", "lopbuopedk", "updomervja", "rzlntxyeuc", "wksxbgyrao", "mbtvksjfjz", "albtzsymge", "udtrzqmdq", "ycohghovgs", "eycjpjhynu", "fnjjhhjuvr", "usqfgqvmkp", "yvkurupifv", "izrgbmyark", "ctzkjkziva", "bjmkxvbwgv", "bqzgexyalb", "sdjsgpngcw", "fkdifibuuf", "fmowditskz", "oqjmqrtict", "ojiyxyesxz", "yfrorodmbn", "drznpnrwcj", "pmhdtjmhay", "orsufumaps", "vgzhblmyyt", "ejvgwffbbg", "gcnqbaere", "nuzjqxmzot", "arlutmygms"}
	for ind, word := range words {
		ok, err := hasLengthTen(word)
		if err != nil {
			fmt.Printf("Error encountered in word ind %v: %v\n", ind, err.Error())
		} else if ok {
			fmt.Printf("Word %v is ok!\n", ind)
		} else {
			fmt.Printf("Word %v has the wrong length!\n", ind)
		}
	}
	// TODO write a function called hasLengthTen that takes in a string
	//    * If the string contains the letter `a`, return an error (HINT the rune value for `a` is `97`)
	//    * Otherwise return a bool that is true if the length of the word is 10 and false if not (HINT `len()` returns length of its argument)
}

func divide(dividend float64, divisor float64) (float64, error) {
	if divisor == 0 {
		return 0, errors.New("you can't divide by zero")
	}
	return dividend / divisor, nil
}
