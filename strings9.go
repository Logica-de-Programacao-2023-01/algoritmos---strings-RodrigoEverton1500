package main

import "fmt"

func main() {
	var str string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	var car1, car2 byte
	fmt.Print("Digite a letra a ser substituída: ")
	fmt.Scanf("%c\n", &car1)
	fmt.Print("Digite a letra de substituição: ")
	fmt.Scanf("%c\n", &car2)
	resultado := ""
	for i := 0; i < len(str); i++ {
		if str[i] == car1 {
			resultado += string(car2)
		} else {
			resultado += string(str[i])
		}
	}
	fmt.Println("String resultante:", resultado)
}
