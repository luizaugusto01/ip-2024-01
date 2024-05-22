package main

import (
	"fmt"
)

func main() {
	// Leitura do tamanho do vetor N e dos valores do vetor V
	var N int
	fmt.Scan(&N)
	vetorV := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&vetorV[i])
	}

	// Armazenando os valores de V em um mapa para busca eficiente
	mapV := make(map[int]struct{}, N)
	for _, num := range vetorV {
		mapV[num] = struct{}{}
	}

	// Leitura do número de buscas M
	var M int
	fmt.Scan(&M)

	// Leitura dos M números a serem buscados no vetor V e realização das buscas
	for i := 0; i < M; i++ {
		var busca int
		fmt.Scan(&busca)
		if _, found := mapV[busca]; found {
			fmt.Println("ACHEI")
		} else {
			fmt.Println("NAO ACHEI")
		}
	}
}
