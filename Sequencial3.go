package main

import "fmt"

func main() {
	var peso, altura float64

	//inserção de dados
	fmt.Println("Digite seu peso:")
	fmt.Scanln(&peso)
	fmt.Println("Digite sua altura:")
	fmt.Scanln(&altura)

	imc := peso / (altura * altura)

	fmt.Printf("O seu imc é: %.1f  ", imc)
}
