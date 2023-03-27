package main

import "fmt"

func main() {
	var num1, num2, num3 float64

	//inserção de dados
	fmt.Println("Digite três números reais:")
	fmt.Scan(&num1, &num2, &num3)

	mediaPonderada := (num1*2 + num2*3 + num3*5) / (2 + 3 + 5)

	fmt.Printf("A média ponderada entre %.2f, %.2f e %.2f é: %.2f", num1, num2, num3, mediaPonderada)
}
