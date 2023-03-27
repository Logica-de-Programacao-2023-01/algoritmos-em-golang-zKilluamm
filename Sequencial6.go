package main

import "fmt"

func main() {
	var salario float64

	fmt.Println("Digite o seu salário em $ para que seja feita o cálculo de aumento :")
	fmt.Scanln(&salario)

	aumento := salario * 0.15
	salarionovo := salario + aumento

	fmt.Printf("O seu novo salário é : $%.3f", salarionovo)
}
