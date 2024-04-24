package main

import "fmt"

func main() {
	var N int

	// Leitura do valor de N
	fmt.Println("Digite um valor inteiro N (5 < N < 2000):")
	fmt.Scanln(&N)

	// Verificação se N está no intervalo permitido
	if N <= 5 || N >= 2000 {
		fmt.Println("O valor de N está fora do intervalo permitido.")
		return
	}

	// Geração do quadrado de cada número par de 1 até N
	for i := 0; i <= N; i += 2 {
		quadrado := i * i
		fmt.Printf("%d^2 = %d\n", i, quadrado)
	}
}
