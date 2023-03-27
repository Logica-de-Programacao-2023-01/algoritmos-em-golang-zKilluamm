package main

import "fmt"

func main() {

	var produto float64

	fmt.Println("Digite o valor do produto:")
	fmt.Scanln(&produto)

	desconto := produto * 0.10
	precofinal := produto - desconto

	fmt.Printf("O preço do produto com desconto é : $%.2f", precofinal)
}
