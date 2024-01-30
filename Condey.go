package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var (
		amount         int
		amountPeople   int
		test           string
		temperMinimum  string
		temperMax      string
		checkingString bool
	)
	//
	temperMinimum = "15"
	temperMax = "30"
	fmt.Scanln(&amount)
	for i := 0; i < amount; i++ {
		fmt.Scan(&amountPeople)
		for inner := 0; inner < amountPeople; inner++ {
			fmt.Scan(&test)
			checkingString = strings.EqualFold(test, "<=")
			fmt.Scan(&test)
			if checkingString {
				temperMax = test
			} else {
				temperMinimum = test
			}
			intMax, _ := strconv.Atoi(temperMax)
			intMin, _ := strconv.Atoi(temperMinimum)
			if intMin > intMax {
				fmt.Println("-1")
			} else if temperMinimum == temperMax {
				fmt.Println(temperMinimum)
			} else {
				fmt.Println(intMin)
			}
		}
		temperMinimum = "15"
		temperMax = "30"

	}
}
