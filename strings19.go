package main

import (
	"fmt"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	r := ""
	for i := len(str) - 1; i >= 0; i-- {
		r += string(str[i])
	}
	fmt.Println("String invertida:", r)
}
