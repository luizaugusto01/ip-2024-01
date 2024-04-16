package main

import (
	"fmt"
)

func main() {
	var n int

	// Entrada do valor inteiro positivo
	fmt.Print("Digite um número inteiro positivo maior que 1: ")
	_, err := fmt.Scan(&n)
	if err != nil || n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	// Cálculo da soma
	soma := 0.0
	for k := 1; k <= n; k++ {
		soma += 1.0 / float64(k)
	}

	// Saída da soma com 6 casas decimais
	fmt.Printf("%.6f\n", soma)
}