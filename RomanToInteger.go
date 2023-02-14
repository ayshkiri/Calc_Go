package main

import "fmt"

func main() {
	var res string

	inputNumber := 1

	tenCount := inputNumber / 10
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
	if inputNumber%10 != 0 {
		ostatok := inputNumber % 10
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

	fmt.Println(res)
}
