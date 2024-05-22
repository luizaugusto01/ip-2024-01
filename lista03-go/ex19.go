package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	// Ler os números inteiros
	nums := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	// Identificar e imprimir os elementos únicos
	uniqueElements := make(map[int]bool)
	for _, num := range nums {
		if !uniqueElements[num] {
			uniqueElements[num] = true
			fmt.Println(num)
		}
	}
}
