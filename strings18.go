package main

import (
	"fmt"
	"strconv"
)

func main() {
	var value string
	fmt.Print("Escreva algo: ")
	fmt.Scan(&value)

	_, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("A string não contém apenas números")
	} else {
		fmt.Println("A string contém apenas números")
	}
}
