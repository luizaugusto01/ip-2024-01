package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// Ler a sequência de inteiros
	nums := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nums[i])
	}

	// Criar um mapa para contar a frequência de cada elemento
	frequency := make(map[int]int)
	for _, num := range nums {
		frequency[num]++
	}

	// Contar o número de elementos únicos que aparecem apenas uma vez
	count := 0
	for _, freq := range frequency {
		if freq == 1 {
			count++
		}
	}

	// Imprimir o número de elementos únicos que aparecem apenas uma vez
	fmt.Println(count)
}
