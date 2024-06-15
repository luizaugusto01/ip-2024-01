package main

import (
	"fmt"
)


func somaArray(arr []float64, n int) float64 {
	
	if n == 0 {
		return 0
	}
	
	return arr[n-1] + somaArray(arr, n-1)
}

func main() {
	
	array := []float64{1.5, 2.7, 3.2, 4.8, 5.1}

	
	total := somaArray(array, len(array))

	
	fmt.Printf("A soma de todos os elementos do array é: %.2f\n", total)
}
