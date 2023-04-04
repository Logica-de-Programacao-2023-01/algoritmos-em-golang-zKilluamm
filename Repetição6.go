package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um número para que seja feita a multiplicação seguindo a tabuada de 0 a 10")
	fmt.Scanln(&num)

	for i := 1; i <= 10; i++ {
		fmt.Printf("%d vezes %d é igual a %d\n", num, i, num*i)

	}

}
