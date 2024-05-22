package main

import "fmt"

func main() {
	for {
		var N int
		fmt.Scan(&N)
		if N == 0 {
			break
		}

		if N <= 1 || N > 10000 {
			return
		}

		values := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&values[i])
		}

		// Encontrar o maior elemento do vetor
		maxValue := values[0]
		for i := 1; i < N; i++ {
			if values[i] > maxValue {
				maxValue = values[i]
			}
		}

		// Inicializar um slice de frequência com tamanho igual a maxValue+1
		frequencies := make([]int, maxValue+1)

		// Calcular a frequência dos valores menores ou iguais a cada valor entre 0 e maxValue
		for i := 0; i < N; i++ {
			for j := 0; j <= values[i]; j++ {
				frequencies[j]++
			}
		}

		// Imprimir as frequências
		for i := 0; i <= maxValue; i++ {
			fmt.Printf("(%d) %d\n", i, frequencies[i])
		}
	}
}
