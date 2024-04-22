package main

import (
	"fmt"
)

func main() {
	var populacaoA, populacaoB int
	fmt.Scan(&populacaoA, &populacaoB)

	anos := 0
	for populacaoA <= populacaoB {
		populacaoA = int(float64(populacaoA) * 1.03)
		populacaoB = int(float64(populacaoB) * 1.015)
		anos++
	}

	fmt.Printf("ANOS = %d\n", anos)
}