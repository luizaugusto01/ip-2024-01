package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Println("Digite três números inteiros diferentes de zero:")
	fmt.Scan(&a, &b, &c)

	mmcAB := mmc(a, b)
	mmcABC := mmc(mmcAB, c)

	fmt.Printf("%5d %5d %5d :%d\n", a, b, c, mmcABC)
}

// Função para calcular o Mínimo Múltiplo Comum (MMC) de dois números inteiros
func mmc(x, y int) int {
	return x * y / mdc(x, y)
}

// Função para calcular o Máximo Divisor Comum (MDC) de dois números inteiros (Algoritmo de Euclides)
func mdc(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}
