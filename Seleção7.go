package main

import "fmt"

func main() {
	var num float64

	fmt.Println("Digite o seu salário para que seja feito o cálculo de aumento:")
	fmt.Scanln(&num)

	var Salario float64

	if num < 100 {
		Salario = num * 1.1

	} else {
		Salario = num * 1.05
	}
	fmt.Printf("O Novo salário é %.2f", Salario)

}
