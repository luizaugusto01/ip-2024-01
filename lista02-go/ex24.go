package main

import (
	"fmt"
)

func main() {
	var x float64
	var N int

	// Leitura do valor de x e N
	fmt.Println("Digite o valor de x (em radianos):")
	fmt.Scan(&x)
	fmt.Println("Digite a quantidade de termos N:")
	fmt.Scan(&N)

	// Inicialização das variáveis
	cosX := 1.0
	term := 1.0
	sign := -1.0

	// Cálculo da série
	for i := 1; i <= N; i++ {
		term *= (x * x) / float64((2 * i * (2*i - 1)))
		cosX += sign * term
		sign *= -1
	}

	// Impressão do resultado
	fmt.Printf("cos(x) = %.2f\n", cosX)
}
