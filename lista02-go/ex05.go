package main

import (
	"fmt"
)

func main() {
	var n int

	fmt.Println("digite o valor de n")
	fmt.Scanln(&n)

	sequencia := make([]int, n)

	fmt.Println("digite o número de elementos na sequencia: ")
	for I := range sequencia {
		fmt.Println("digite o valor de ", I)
		fmt.Scanln(&sequencia[I])
	}

	comprimentomaximo := 0
	comprimentoatual := 1

	for i := 1; i < n; i++ {
		if sequencia[i] > sequencia[i-1] {
			comprimentoatual++
		} else {
			if comprimentoatual > comprimentomaximo {
				comprimentomaximo = comprimentoatual
			}
			comprimentoatual = 1
		}
	}
	if comprimentoatual > comprimentomaximo {
		comprimentomaximo = comprimentoatual
	}
	fmt.Printf("O comprimento do segmento crescente maximo é: %d\n", comprimentomaximo-1)
}
