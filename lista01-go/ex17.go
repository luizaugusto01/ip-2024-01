package main

import (
	"fmt"
	"math"
)

func main() {
	var N int

	// Entrada do valor inteiro N
	fmt.Print("Digite um valor inteiro N (5 < N < 2000): ")
	fmt.Scan(&N)

	// Verificação do intervalo de N
	if N <= 5 || N >= 2000 {
		fmt.Println("O valor de N deve estar entre 5 e 2000.")
		return
	}

	// Loop para gerar o quadrado dos valores pares até N
	for i := 2; i <= N; i += 2 {
		quadrado := int(math.Pow(float64(i), 2))
		fmt.Printf("%d^2 = %d\n", i, quadrado)
	}
}