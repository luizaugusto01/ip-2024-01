package main

import (
	"fmt"
)

func main() {
	// Leitura do número de grãos para o primeiro quadrado
	var n int
	fmt.Print("Digite o número de grãos para o primeiro quadrado: ")
	fmt.Scan(&n)

	// Variável para armazenar o total de grãos
	total := 0

	// Loop para percorrer o tabuleiro de xadrez
	for i := 1; i <= 64; i++ {
		// Se o quadrado for escuro, dobrar a quantidade de grãos
		if i%2 == 0 {
			total += n
		} else {
			total += n / 2 // Se o quadrado for branco, manter a mesma quantidade de grãos
		}
		n *= 2 // Dobrar a quantidade de grãos para o próximo quadrado
	}

	// Imprimir o total de grãos no tabuleiro
	fmt.Println("Total de grãos no tabuleiro de xadrez:", total)
}
