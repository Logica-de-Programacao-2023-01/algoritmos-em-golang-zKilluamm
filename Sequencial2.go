package main

import "fmt"

func main() {
	var num1 int
	fmt.Print("Digite um número para que seja feita a operação:2")
	fmt.Scanln(&num1)
	var dobro int = (num1 * 2)
	var triplo int = (num1 * 3)
	var quad int = (num1 * 4)
	fmt.Println("O dobro do número é: ", dobro, ", O triplo do número é : ", triplo, " e o quadruplo : ", quad)

}
