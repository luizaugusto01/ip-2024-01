package main

import (
	"fmt"
)

func main() {
	var x float64
	var N int

	// Leitura do valor de x e N
	fmt.Println("Digite o valor de x:")
	fmt.Scan(&x)
	fmt.Println("Digite a quantidade de termos N:")
	fmt.Scan(&N)

	// Inicialização das variáveis
	ex := 1.0
	fact := 1.0

	// Cálculo da série
	for i := 1; i <= N; i++ {
		fact *= float64(i)
		ex += power(x, i) / fact
	}

	// Impressão do resultado
	fmt.Printf("e^x = %.2f\n", ex)
}

// Função para calcular a potência de x elevado a n
func power(x float64, n int) float64 {
	result := 1.0
	for i := 0; i < n; i++ {
		result *= x
	}
	return result
}
