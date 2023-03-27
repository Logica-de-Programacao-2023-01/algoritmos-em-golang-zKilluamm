package main

import "fmt"

func main() {
	var num1 int
	fmt.Println("Digite um número inteiro :")
	fmt.Scanln(&num1)

	num_sucessor := num1 + 1
	num_antecessor := num1 - 1

	fmt.Printf("O seu sucessor de %d é %d, e o seu antecessor é %d", num1, num_sucessor, num_antecessor)
}
