package main

import "fmt"

func main() {
	var anos float64
	fmt.Println("Digite sua idade:")
	fmt.Scanln(&anos)

	anos_dias := (anos * 365.25)

	fmt.Printf("A sua idade em dias é : %.2f", anos_dias)

}
