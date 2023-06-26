package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	dec := true
	len := len(str)
	for i := 0; i < len-1; i++ {
		curr, _ := strconv.Atoi(string(str[i]))
		next, _ := strconv.Atoi(string(str[i+1]))

		if next != curr-1 {
			dec = false
			break
		}
	}
	if dec {
		fmt.Println("A string é uma sequência numérica decrescente.")
	} else {
		fmt.Println("A string não é uma sequência numérica decrescente.")
	}
}
