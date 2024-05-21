package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	// Inicializa um array para representar o tabuleiro
	var tabuleiro [64]int

	// Preenche o tabuleiro com a quantidade de grãos em cada quadrado
	for i := range tabuleiro {
		if i%2 == 0 {
			tabuleiro[i] = n
		} else {
			tabuleiro[i] = 2 * n
		}
		if i >= 8 && i%8 == 0 {
			n *= 2
		}
	}

	// Calcula a quantidade total de grãos no tabuleiro
	total := 0
	for _, valor := range tabuleiro {
		total += valor
	}

	fmt.Println(total)
}
