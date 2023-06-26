package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	cres := true
	len := len(str)
	for i := 0; i < len-1; i++ {
		curr, _ := strconv.Atoi(string(str[i]))
		next, _ := strconv.Atoi(string(str[i+1]))

		if next != curr+1 {
			cres = false
			break
		}
	}
	if cres {
		fmt.Println("A string é uma sequência numérica crescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica crescente.")
	}
}
