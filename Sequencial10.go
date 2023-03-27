package main

import "fmt"

func main() {
	var peso float64


	//inserção de dados
	fmt.Print("Digite seu peso em quilos para a conversão :")
	fmt.Scanln(&peso)

	 pesoLb := peso * 2.2046

	fmt.Printf("O seu peso em libra é : %.2f", pesoLb)

