package main

import "fmt"

func vetorcrescente(vetor []int, N int) []int {
	for i := 0; i < N; i++ {
		for j := N - 1; j > i; j-- {
			if vetor[j] < vetor[i] {
				vetor[i], vetor[j] = vetor[j], vetor[i]

			}

		}
	}
	return vetor
}

func main() {
	nummero := []int{9, 8, 7, 6}
	fmt.Printf("%d", vetorcrescente(nummero, 4))
}
