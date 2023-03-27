package main

import "fmt"

func main() {

	var num1, num2, num3 int

	fmt.Print("Digite um número inteiro:")
	fmt.Scanln(&num1)
	fmt.Print("Digite outro inteiro:")
	fmt.Scanln(&num2)
	fmt.Print("Digite o terceiro inteiro:")
	fmt.Scanln(&num3)

	total := (num1 + num2 + num3)

	fmt.Print("O valor da soma entre os números é :", total)

}
