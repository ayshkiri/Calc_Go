package main

import "fmt"

func main() {
	var numOne, numTwo, result int
	var operator string
	fmt.Scanln(&numOne, &operator, &numTwo)
	switch operator {
	case "+":
		result = numOne + numTwo
		fmt.Println(result)
	case "-":
		result = numOne + numTwo
		fmt.Println(result)
	case "*":
		result = numOne + numTwo
		fmt.Println(result)
	case "/":
		if numTwo != 0 {
			result = numOne + numTwo
			fmt.Println(result)
		} else {
			fmt.Println("Делить на 0 нельзя!")
		}
	}

}
