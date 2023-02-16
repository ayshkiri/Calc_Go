package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
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
		err := errors.New("Некорректное римское число от 1 (I) до 10 (X)")
		return 0, err
	}
}

func IntegerToRoman(number int) string {
	var result string
	romanNumbers := []string{"C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	arabicNumbers := []int{100, 90, 50, 40, 10, 9, 5, 4, 1}
	for i := 0; number > 0; i++ {
		for arabicNumbers[i] <= number {
			result += romanNumbers[i]
			number -= arabicNumbers[i]
		}
	}
	return result
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	inputData, _ := reader.ReadString('\n')
	inputData = strings.Trim(inputData, "\n")
	var inputStrArray []string = strings.Split(inputData, " ")
	if len(inputStrArray) != 3 {
		fmt.Println("Некорректные входные данные")
		return
	}
	operators := map[string]int{"+": 1, "-": 1, "*": 1, "/": 1}
	_, check := operators[inputStrArray[1]]
	if !check {
		fmt.Println("Некорректные входные данные")
		return
	}
	numberOne, errNumberOne := strconv.Atoi(inputStrArray[0])
	numberTwo, errNumberTwo := strconv.Atoi(inputStrArray[2])
	if errNumberOne == nil && errNumberTwo == nil && numberOne >= 1 && numberOne <= 10 && numberTwo >= 1 && numberTwo <= 10 {
		// Возвращаемая функцией Calc ошибка пропускается, т.к. в условии выше уже есть проверка,
		// что передаваемые аргументы больше 0
		arabicResult, _ := Calc(numberOne, numberTwo, inputStrArray[1])
		fmt.Println(arabicResult)
		return
	}
	romanNumberOne, errRomanNumberOne := RomanToInteger(inputStrArray[0])
	romanNumberTwo, errRomanNumberTwo := RomanToInteger(inputStrArray[2])
	if errRomanNumberOne == nil && errRomanNumberTwo == nil && romanNumberOne > romanNumberTwo {
		romanResult, _ := Calc(romanNumberOne, romanNumberTwo, inputStrArray[1])
		fmt.Println(IntegerToRoman(romanResult))
		return
	}
	fmt.Println("Некорректные входные данные")
}

func Calc(numOne, numTwo int, operator string) (int, error) {
	var result int
	switch operator {
	case "+":
		result = numOne + numTwo
	case "-":
		result = numOne - numTwo
	case "*":
		result = numOne * numTwo
	case "/":
		if numTwo != 0 {
			result = numOne / numTwo
		} else {
			err := errors.New("Делить на 0 нельзя")
			return 0, err
		}
	}
	return result, nil
}
