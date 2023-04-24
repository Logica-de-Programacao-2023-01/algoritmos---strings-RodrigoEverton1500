package main

import (
	"fmt"
	"strings"
)

func main() {
	var value string
	fmt.Print("Escreva algo: ")
	fmt.Scan(&value)

	if strings.ToLower(string(value[0])) != string(value[0]) {
		fmt.Println("Não está em camelCase")
	} else {
		palavras := 1
		for _, x := range value {
			if strings.ToUpper(string(x)) == string(x) {
				palavras++
			}
		}
		fmt.Printf("Tem %d palavras na string\n", palavras)
	}
}
