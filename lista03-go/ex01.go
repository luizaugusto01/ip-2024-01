package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var n int

	// Leitura do valor de N
	for {
		fmt.Scanln(&n)
		if n >= 1 && n <= 1000 {
			break
		}
	}

	// Leitura dos valores do vetor V
	var vInput string
	fmt.Scanln(&vInput)
	vStrings := strings.Split(vInput, " ")
	v := make([]int, n)
	for i, numStr := range vStrings {
		v[i], _ = strconv.Atoi(numStr)
	}

	// Leitura do valor de K
	var k int
	fmt.Scanln(&k)

	// Contabiliza quantos elementos são maiores ou iguais a K
	count := 0
	for _, num := range v {
		if num >= k {
			count++
		}
	}

	// Imprime o resultado
	fmt.Println(count)
}
