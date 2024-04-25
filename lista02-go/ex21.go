package main

import (
	"fmt"
)

func main() {
	for {
		var size int
		fmt.Scan(&size)

		// Verifica se é o fim da entrada
		if size == 0 {
			break
		}

		// Lê a sequência de números
		sequence := make([]float64, size)
		for i := 0; i < size; i++ {
			fmt.Scan(&sequence[i])
		}

		// Verifica se a sequência está em ordem crescente
		ordered := true
		for i := 1; i < size; i++ {
			if sequence[i-1] >= sequence[i] {
				ordered = false
				break
			}
		}

		// Imprime o resultado
		if ordered {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}
	}
}
