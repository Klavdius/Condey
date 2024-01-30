package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	var (
		amount         int
		amountPeople   int
		temperMinimum  string
		temperMax      string
		checkingString bool
	)
	//
	temperMinimum = "15"
	temperMax = "30"
	fmt.Scanln(&amount)
	for i := 0; i < amount; i++ {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Scanln(&amountPeople)
		for inner := 0; inner < amountPeople; inner++ {
			scanner.Scan()
			inputTemp := strings.Split(scanner.Text(), " ")
			checkingString = strings.EqualFold(inputTemp[0], "≤")
			if checkingString {
				temperMax = inputTemp[1]
			} else {
				temperMinimum = inputTemp[1]
			}
			intMax, _ := strconv.Atoi(temperMax)
			intMin, _ := strconv.Atoi(temperMinimum)
			if intMin > intMax {
				fmt.Println("-1")
			} else if temperMinimum == temperMax {
				fmt.Println(temperMinimum)
			} else {
				numberOutput := RandomIntInRange(intMin, intMax)
				fmt.Println(numberOutput)
			}
		}
		temperMinimum = "15"
		temperMax = "30"

	}
}

func RandomIntInRange(min int, max int) int {
	number := rand.Intn(max) + min

	if number > max {
		number = number - min
	}
	for i := number; i <= min; i++ {
		number = i
	}

	return number
}
