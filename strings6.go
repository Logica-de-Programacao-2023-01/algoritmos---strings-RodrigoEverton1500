package main

import "fmt"

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)
	conta := 0
	valor := false

	for _, e := range input {
		if e == ' ' || e == '\t' || e == '\n' {
			valor = false
		} else if !valor {
			conta++
			valor = true
		}
	}
	fmt.Println("Número de palavras:", conta)
}
