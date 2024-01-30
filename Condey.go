package main

import (
	"fmt"
	"strconv"
	"strings"
)

var band = []int{15, 30}

func main() {
	var (
		amount       int
		amountPeople int
		inputChar    string
		inputNumber  string
	)

	checkErrorBand := false
	fmt.Scanln(&amount)
	for i := 0; i < amount; i++ {
		fmt.Scan(&amountPeople)
		for inner := 0; inner < amountPeople; inner++ {
			fmt.Scan(&inputChar)
			checkChar := strings.EqualFold(inputChar, "<=")
			fmt.Scan(&inputNumber)
			if checkErrorBand {
				fmt.Println("-1")
			} else {
				checkErrorBand = CheckingNumberInBand(inputNumber)
				if checkErrorBand {
					fmt.Println("-1")
				} else {
					if checkChar {
						band[1], _ = strconv.Atoi(inputNumber)
					} else {
						band[0], _ = strconv.Atoi(inputNumber)
					}
					fmt.Println(strconv.Itoa(band[0]))
				}
			}
		}
		band[0] = 15
		band[1] = 30
		checkErrorBand = false
	}
}

func CheckingNumberInBand(inputText string) bool {
	inputNumber, _ := strconv.Atoi(inputText)
	var result bool
	if inputNumber >= band[0] && inputNumber <= band[1] {
		result = false
	} else {
		result = true
	}
	return result
}
