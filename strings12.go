package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	str = strings.ToLower(strings.ReplaceAll(str, " ", ""))
	palindrome := true
	len := len(str)

	for i := 0; i < len/2; i++ {
		if str[i] != str[len-1-i] {
			palindrome = false
			break
		}
	}
	if palindrome {
		fmt.Println("A string é um palíndromo.")
	} else {
		fmt.Println("A string não é um palíndromo.")
	}
}
