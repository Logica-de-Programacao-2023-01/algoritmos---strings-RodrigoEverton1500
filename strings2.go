package main

import (
	"fmt"
	"strings"
)

func main() {

	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)

	fmt.Println("Resultado:", strings.ReplaceAll(input, " ", ""))
}
