package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o número de elementos na sequência: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("O número de elementos na sequência deve ser maior que zero.")
		return
	}

	// Lendo a sequência de números
	sequencia := make([]int, n)
	fmt.Println("Digite a sequência de números:")
	for i := 0; i < n; i++ {
		fmt.Scan(&sequencia[i])
	}

	// Encontrando a maior sub-sequência de números iguais consecutivos
	maiorSubsequencia := 1
	contagemAtual := 1
	for i := 1; i < n; i++ {
		if sequencia[i] == sequencia[i-1] {
			contagemAtual++
		} else {
			contagemAtual = 1
		}
		if contagemAtual > maiorSubsequencia {
			maiorSubsequencia = contagemAtual
		}
	}

	if maiorSubsequencia == 1 {
		fmt.Printf("Não existe uma subsequencia ")
		return

	}

	// Imprimindo o resultado
	fmt.Printf("A maior subsequência de elementos iguais possui %d elementos.\n", maiorSubsequencia)
}
