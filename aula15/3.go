package main

import (
	"fmt"
)

func inverterArray(arr []int, tamanho int) {
	esquerda := 0
	direita := tamanho - 1

	for esquerda < direita {
		
		arr[esquerda], arr[direita] = arr[direita], arr[esquerda]
		
		esquerda++
		direita--
	}
}

func main3() {
	
	array := []int{1, 2, 3, 4, 5}
	tamanho := len(array)

	fmt.Println("Array antes da inversão:", array)

	
	inverterArray(array, tamanho)

	fmt.Println("Array depois da inversão:", array)
}
