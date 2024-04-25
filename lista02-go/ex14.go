package main

import "fmt"

func main() {
	var m, n int
	fmt.Print("Digite o número de linhas da matriz: ")
	fmt.Scan(&m)
	fmt.Print("Digite o número de colunas da matriz: ")
	fmt.Scan(&n)

	// Verificando se a matriz é válida
	if m <= 0 || n <= 0 {
		fmt.Println("As dimensões da matriz devem ser positivas.")
		return
	}

	// Imprimindo os pares de índices inferiores à diagonal principal
	fmt.Println("Pares de índices inferiores à diagonal principal:")
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if j < i {
				fmt.Printf("(%d-%d) ", i, j)
			}
		}
		fmt.Println()
	}
}
