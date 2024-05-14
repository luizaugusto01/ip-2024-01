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

	// Leitura da primeira linha que contém o valor de n
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Leitura da segunda linha que contém os valores inteiros
	scanner.Scan()
	secondLine := scanner.Text()
	elements := strings.Split(secondLine, " ")

	// Mapa para contar as ocorrências de cada número
	occurrences := make(map[int]int)

	// Preenchendo o mapa com as contagens
	for i := 0; i < n; i++ {
		num, _ := strconv.Atoi(elements[i])
		occurrences[num]++
	}

	// Contando os números que aparecem exatamente uma vez
	uniqueCount := 0
	for _, count := range occurrences {
		if count == 1 {
			uniqueCount++
		}
	}

	// Imprimindo o resultado
	fmt.Println(uniqueCount)
}
