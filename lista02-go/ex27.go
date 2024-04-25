package main

import (
	"fmt"
)

func main() {
	var num int
	fmt.Println("Digite um número inteiro maior que 1:")
	fmt.Scan(&num)

	// Verifica se o número é válido
	if num <= 1 {
		fmt.Printf("Fatoracao nao e possivel para o numero %d!\n", num)
		return
	}

	fmt.Printf("Fatoracao de %d: ", num)

	// Fatoração em números primos
	for i := 2; i <= num; i++ {
		for num%i == 0 {
			fmt.Printf("%d ", i)
			num /= i
		}
	}

	fmt.Println()
}
