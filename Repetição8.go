package main

import "fmt"

func main() {
	var num int

	fmt.Println("Digite um número inteiro")
	fmt.Scanln(&num)

	fmt.Printf("Os divisores de %d são :\n", num)
	for i := 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println(i)

		}

	}
}
