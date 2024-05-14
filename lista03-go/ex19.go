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

	// Leitura do número de elementos
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	elements := make([]int, N)

	// Leitura dos N números
	for i := 0; i < N; i++ {
		scanner.Scan()
		num, _ := strconv.Atoi(scanner.Text())
		elements[i] = num
	}

	// Mapa para verificar elementos únicos
	uniqueElements := make(map[int]bool)
	for _, num := range elements {
		uniqueElements[num] = true
	}

	// Coletar os elementos únicos e colocá-los em um slice
	var result []int
	for num := range uniqueElements {
		result = append(result, num)
	}

	// Ordenar o slice result
	sort.Ints(result)

	// Imprimir os elementos únicos em ordem crescente
	for _, num := range result {
		fmt.Println(num)
	}
}
