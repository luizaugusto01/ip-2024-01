package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int

	// Leitura do valor de n
	fmt.Scanln(&n)

	// Leitura dos n valores inteiros
	var valuesStr string
	fmt.Scanln(&valuesStr)
	valuesStrings := strings.Split(valuesStr, " ")
	values := make([]int, n)
	for i, numStr := range valuesStrings {
		values[i], _ = strconv.Atoi(numStr)
	}

	// Imprime os valores na ordem inversa
	for i := n - 1; i >= 0; i-- {
		if i != n-1 {
			fmt.Print(" ")
		}
		fmt.Print(values[i])
	}
	fmt.Println()
}