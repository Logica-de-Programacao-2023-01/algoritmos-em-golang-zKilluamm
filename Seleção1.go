package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Digite dois números inteiros :")
	fmt.Scanln(&num1, &num2)

	if num1 > num2 {
		fmt.Printf("%d é maior que %d", num1, num2)
	} else if num1 < num2 {
		fmt.Printf("%d é menor que %d", num1, num2)
	} else {
		fmt.Printf("%d é igual a %d", num1, num2)
	}
}
