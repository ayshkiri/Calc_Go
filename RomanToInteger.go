package main

import "fmt"

func main() {
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
	var res = romanNumbers["X"]
	fmt.Println(res)
}
