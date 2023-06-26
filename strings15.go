package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	str2 := strings.Map(func(r rune) rune {
		if strings.ContainsRune("AEIOUaeiou", r) {
			return '*'
		}
		return r
	}, str)
	fmt.Println("Nova string:", str2)
}
