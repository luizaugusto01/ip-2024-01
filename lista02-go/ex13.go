package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scanln(&n)

	grains := calcularGraos(n)
	fmt.Println(grains)
}

func calcularGraos(n int) int {
	var totalGraos int

	// Percorrer o tabuleiro de 64 quadros
	for i := 0; i < 64; i++ {
		// Calcular o número de grãos para o quadro atual
		if i%2 == 0 {
			totalGraos += n // Quadro branco, mesma quantidade de grãos
		} else {
			totalGraos += n * (1 << (i / 2)) // Quadro escuro, dobrar progressivamente
		}
	}

	return totalGraos
}
