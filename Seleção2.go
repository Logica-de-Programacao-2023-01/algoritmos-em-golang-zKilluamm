package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("Digite três números inteiros :")
	fmt.Scanln(&num1, &num2, &num3)

	if num1 < num2 && num1 < num3 {
		fmt.Printf("%d é menor que todos os números digitados", num1)

	} else if num2 < num1 && num3 < num3 {
		fmt.Printf("%d é menos que todos os números digitados", num2)

	} else {
		fmt.Printf("%d é menor que todos os números digitados", num3)
	}

}
