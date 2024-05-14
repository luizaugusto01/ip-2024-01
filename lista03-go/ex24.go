package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// Lê o tamanho do vetor
		n, _ := strconv.Atoi(scanner.Text())
		if n == 0 {
			break
		}

		// Lê os elementos do vetor
		scanner.Scan()
		line := scanner.Text()
		elements := strings.Split(line, " ")
		arr := make([]int, n)
		for i := range arr {
			arr[i], _ = strconv.Atoi(elements[i])
		}

		// Executa o Counting Sort
		sortedArr := countingSort(arr, 10000)

		// Imprime o vetor ordenado
		for i, val := range sortedArr {
			if i > 0 {
				fmt.Print(" ")
			}
			fmt.Print(val)
		}
		fmt.Println()
	}
}

// Função que implementa o Counting Sort
func countingSort(arr []int, maxVal int) []int {
	// Cria os vetores vCount e vOrd
	vCount := make([]int, maxVal+1)
	vOrd := make([]int, len(arr))

	// Passo 2: Inicializa vCount com 0
	for i := range vCount {
		vCount[i] = 0
	}

	// Passo 3: Conta a ocorrência de cada valor em arr
	for _, val := range arr {
		vCount[val]++
	}

	// Passo 4: Acumula as contagens em vCount
	for i := 1; i <= maxVal; i++ {
		vCount[i] += vCount[i-1]
	}

	// Passo 5: Ordena os valores em vOrd
	for i := len(arr) - 1; i >= 0; i-- {
		val := arr[i]
		vCount[val]--
		vOrd[vCount[val]] = val
	}

	// Passo 6: Copia vOrd de volta para arr
	copy(arr, vOrd)

	return arr
}
