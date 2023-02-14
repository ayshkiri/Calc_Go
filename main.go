package main

import (
	"fmt"
	"strconv"
)

func RomanToInteger(romanNumber string) int {
	var romanNumbers = make(map[string]int)
	romanNumbers["I"] = 1
	romanNumbers["II"] = 2
	romanNumbers["III"] = 3
	romanNumbers["IV"] = 4
	romanNumbers["V"] = 5
	romanNumbers["VI"] = 6
	romanNumbers["VII"] = 7
	romanNumbers["VIII"] = 8
	romanNumbers["IX"] = 9
	romanNumbers["X"] = 10
	if romanNumbers[romanNumber] != 0 {
		return romanNumbers[romanNumber]
	} else {
		// Twrow exception!!!!!!!!!!!!!!!!!!!!
		return 0
	}
}

func IntegerToRoman(number int) string {
	// Add realisation
	return "XX"
}

func main() {
	var numOneStr, numTwoStr, operator string
	fmt.Scan(&numOneStr, &operator, &numTwoStr)

	numOneInt, err := strconv.Atoi(numOneStr)
	if err != nil {
		panic(err)
	}

	fmt.Println(numOneStr, numOneInt)

	//var result int
	//switch operator {
	//case "+":
	//	result = numOne + numTwo
	//	fmt.Println(result)
	//case "-":
	//	result = numOne - numTwo
	//	fmt.Println(result)
	//case "*":
	//	result = numOne * numTwo
	//	fmt.Println(result)
	//case "/":
	//	if numTwo != 0 {
	//		result = numOne / numTwo
	//		fmt.Println(result, len(operator))
	//	} else {
	//		fmt.Println("Делить на 0 нельзя!")
	//	}
	//}
}
