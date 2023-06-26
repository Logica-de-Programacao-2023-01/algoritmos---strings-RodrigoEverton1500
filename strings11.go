package main

import (
	"fmt"
	"strings"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	resultado := ""
	vogais := "aeiouAEIOU"

	for i := 0; i < len(str); i++ {
		char := string(str[i])
		if !strings.Contains(vogais, char) {
			resultado += char
		}
	}
	fmt.Println("String resultante:", resultado)
}
