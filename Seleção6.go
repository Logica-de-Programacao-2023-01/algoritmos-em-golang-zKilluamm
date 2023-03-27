package main

import "fmt"

func main() {

	var num1, num2 int

	fmt.Println("Digite um número inteiro:")
	fmt.Scanln(&num1)
	fmt.Println("Digite outro número inteiro:")
	fmt.Scanln(&num2)

	if num1 >= 0 && num2 >= 0 {
		fmt.Printf("O resultado da multiplicação é: %d\n", num1*num2)
	} else {
		fmt.Printf("O resultado da soma é: %d\n", num1+num2)
	}
}
