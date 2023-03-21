package main

import "fmt"

func main() {
	var num int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&num)

	if num%2 == 0 {
		fmt.Printf("%d é um número par\n", num)
	} else {
		fmt.Printf("%d é um número ímpar\n", num)
	}
}
