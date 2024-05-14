package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	for {
		// Leitura do valor de N
		var n int
		fmt.Scanln(&n)

		// Verifica se N é 0, indicando o fim da entrada
		if n == 0 {
			break
		}

		// Leitura dos valores do vetor
		var valuesStr string
		fmt.Scanln(&valuesStr)
		values := strings.Split(valuesStr, " ")

		maxIndex := 0
		maxValue := 0

		// Encontra o maior valor e seu índice
		for i := 0; i < n; i++ {
			val, _ := strconv.Atoi(values[i])
			if i == 0 || val > maxValue {
				maxValue = val
				maxIndex = i
			}
		}

		// Imprime o índice e o maior valor
		fmt.Printf("%d %d\n", maxIndex, maxValue)
	}
}
