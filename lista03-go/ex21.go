package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Ler a quantidade de números a serem processados
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// Slices para armazenar os números pares e ímpares
	var pares []int
	var impares []int

	// Ler os números e separar em pares e ímpares
	for i := 0; i < N; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		if num%2 == 0 {
			pares = append(pares, num)
		} else {
			impares = append(impares, num)
		}
	}

	// Ordenar os pares em ordem crescente
	sort.Ints(pares)

	// Ordenar os ímpares em ordem decrescente
	sort.Sort(sort.Reverse(sort.IntSlice(impares)))

	// Imprimir os pares
	for _, num := range pares {
		fmt.Println(num)
	}

	// Imprimir os ímpares
	for _, num := range impares {
		fmt.Println(num)
	}
}
