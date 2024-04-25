package main

import (
	"fmt"
	"strconv"
)

func main() {
	var T int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		var A, B int
		fmt.Scan(&A, &B)

		// Invertendo os números
		AInvertido := inverterNumero(A)
		BInvertido := inverterNumero(B)

		// Verificando qual é o maior número invertido
		var maiorInvertido int
		if AInvertido > BInvertido {
			maiorInvertido = AInvertido
		} else {
			maiorInvertido = BInvertido
		}

		// Imprimindo o maior número invertido
		fmt.Println(maiorInvertido)
	}
}

// Função para inverter um número
func inverterNumero(num int) int {
	// Convertendo o número para string para facilitar a inversão
	numString := strconv.Itoa(num)

	// Invertendo a string
	invertido := ""
	for i := len(numString) - 1; i >= 0; i-- {
		invertido += string(numString[i])
	}

	// Convertendo a string invertida de volta para número
	numInvertido, _ := strconv.Atoi(invertido)

	return numInvertido
}
