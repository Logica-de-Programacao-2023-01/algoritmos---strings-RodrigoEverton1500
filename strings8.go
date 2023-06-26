package main

import "fmt"

func main() {
	var input string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&input)
	r := ""
	for i := len(input) - 1; i >= 0; i-- {
		r += string(input[i])
	}
	fmt.Println("String invertida:", r)
}
