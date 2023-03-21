package main

import "fmt"

func main() {
	var altura float64
	var sexo string

	fmt.Print("Digite a altura (em metros): ")
	fmt.Scan(&altura)

	fmt.Print("Digite o sexo (M para masculino ou F para feminino): ")
	fmt.Scan(&sexo)

	var pesoIdeal float64

	if sexo == "M" {
		pesoIdeal = (72.7 * altura) - 58
	} else if sexo == "F" {
		pesoIdeal = (62.1 * altura) - 44.7
	} else {
		fmt.Println("Sexo inválido. Digite M para masculino ou F para feminino.")
		return
	}

	fmt.Printf("O peso ideal para uma pessoa de %.2fm de altura e sexo %s é %.2fkg.\n", altura, sexo, pesoIdeal)

	var peso float64
	fmt.Print("Digite o peso atual (em kg): ")
	fmt.Scan(&peso)

	if peso < pesoIdeal {
		fmt.Println("A pessoa está abaixo do peso ideal.")
	} else if peso > pesoIdeal {
		fmt.Println("A pessoa está acima do peso ideal.")
	} else {
		fmt.Println("A pessoa está dentro do peso ideal.")
	}
}
