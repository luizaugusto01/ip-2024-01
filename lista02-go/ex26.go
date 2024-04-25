package main

import (
	"fmt"
	"math"
)

func main() {
	var x float64
	var N int

	// Leitura do valor de x e N
	fmt.Println("Digite o valor de x em radianos:")
	fmt.Scan(&x)
	fmt.Println("Digite a quantidade de termos N:")
	fmt.Scan(&N)

	// Inicialização das variáveis
	sinX := 0.0

	// Cálculo da série
	sign := 1.0
	for i := 0; i <= N; i++ {
		term := factorial(2*i+1) * math.Pow(x, float64(2*i+1)) * sign
		sinX += term
		sign *= -1
	}

	// Impressão do resultado
	fmt.Printf("seno(x) = %.6f\n", sinX)
}

// Função para calcular o fatorial de um número inteiro
func factorial(n int) float64 {
	result := 1.0
	for i := 1; i <= n; i++ {
		result *= float64(i)
	}
	return result
}
