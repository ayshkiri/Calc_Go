package main

import (
	"errors"
	"fmt"
	"strconv"
)

func RomanToInteger(romanNumber string) (int, error) {
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
		return romanNumbers[romanNumber], nil
	} else {
		err := errors.New("IncorrectRomanNumber")
		return 0, err
	}
}

func IntegerToRoman(number int) (string, error) {
	var res string
	if number >= 1 && <= 10{
		tenCount := number / 10
		if tenCount == 10 {
			res += "C"
		}
		if tenCount == 9 {
			res += "XC"
		}
		if tenCount >= 5 && tenCount < 9 {
			res += "L"
			for i := 1; i <= tenCount-5; i++ {
				res += "X"
			}
		}
		if tenCount >= 1 && tenCount < 5 {
			for i := 1; i <= tenCount; i++ {
				res += "X"
			}
		}
		if number%10 != 0 {
			ostatok := number % 10
			if ostatok == 9 {
				res += "IX"
			}
			if ostatok >= 5 && ostatok < 9 {
				res += "V"
				for i := 1; i <= ostatok-5; i++ {
					res += "I"
				}
			}
			if ostatok >= 1 && ostatok < 5 {
				for i := 1; i <= ostatok; i++ {
					res += "I"
				}
			}
		}
		return res, nil
	} else {
		err := errors.New("IncorrectArabicNumber")
		return "", err
	}
}

func main() {
	var numOneStr, numTwoStr, operator string
	fmt.Scan(&numOneStr, &operator, &numTwoStr)

	numOneInt, errOne := strconv.Atoi(numOneStr)
	if errOne != nil {
		fmt.Println("Первый аргумент не является числом")
	}

	fmt.Println("Первое число:", numOneStr, numOneInt)

	numTwoInt, errTwo := RomanToInteger(numTwoStr)
	if errTwo != nil {
		fmt.Println("Такого римского числа нет")
	}
	fmt.Println("Второе число: ", numTwoStr, numTwoInt)

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
