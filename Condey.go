package main

import (
	"bufio"
	"fmt"
	"os"
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
			checkingString = strings.EqualFold(inputTemp[0], "<=")
			if checkingString {
				temperMax = inputTemp[1]
			} else {
				temperMinimum = inputTemp[1]
			}
			//
			//	//if inputTemp[0] == "1" {
			//	//	inputNum, _ := strconv.Atoi(inputTemp[1])
			//	//	temperMax = inputNum
			//	//} else {
			//	//	inputNum, _ := strconv.Atoi(inputTemp[1])
			//	//	temperMinimum = inputNum
			//	//}
			//	fmt.Println(inputTemp[1] + " " + inputTemp[0])
			//	//	//if temperMinimum > temperMax {
			//	//	//	fmt.Println("-1")
			//	//	//} else if temperMinimum == temperMax {
			//	//	//	fmt.Println(temperMinimum)
			//	//	//} else {
			//	//	//	number := RandomIntInRange(temperMinimum, temperMax)
			//	//	//	fmt.Println(number)
			//	//}
			//
			fmt.Println(temperMinimum + " " + temperMax)
			fmt.Println(inputTemp[0] + " ")
		}
		fmt.Println(temperMinimum + " " + temperMax)
		temperMinimum = "15"
		temperMax = "30"

	}
}

//func RandomIntInRange(min int, max int) int {
//	number := rand.Intn(max) + min
//
//	if number > max {
//		number = number - min
//	}
//	for i := number; i <= min; i++ {
//		number = i
//	}
//
//	return number
//}
