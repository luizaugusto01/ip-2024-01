package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Função que converte uma linha de entrada em um slice de inteiros
func readNumbers(line string) []int {
	parts := strings.Fields(line)
	numbers := make([]int, len(parts))
	for i, part := range parts {
		numbers[i], _ = strconv.Atoi(part)
	}
	return numbers
}

// Função que conta quantos números de uma aposta coincidem com os sorteados
func countMatches(guess, result []int) int {
	matches := 0
	for _, g := range guess {
		for _, r := range result {
			if g == r {
				matches++
				break
			}
		}
	}
	return matches
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Ler as dezenas sorteadas
	scanner.Scan()
	result := readNumbers(scanner.Text())

	// Ler a quantidade de apostas
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	// Contadores para sena, quina e quadra
	senaCount := 0
	quinaCount := 0
	quadraCount := 0

	// Ler e processar cada aposta
	for i := 0; i < n; i++ {
		scanner.Scan()
		guess := readNumbers(scanner.Text())
		matches := countMatches(guess, result)
		switch matches {
		case 6:
			senaCount++
		case 5:
			quinaCount++
		case 4:
			quadraCount++
		}
	}

	// Imprimir resultados
	if senaCount > 0 {
		fmt.Printf("Houve %d acertador(es) da sena\n", senaCount)
	} else {
		fmt.Println("Nao houve acertador para sena")
	}
	if quinaCount > 0 {
		fmt.Printf("Houve %d acertador(es) da quina\n", quinaCount)
	} else {
		fmt.Println("Nao houve acertador para quina")
	}
	if quadraCount > 0 {
		fmt.Printf("Houve %d acertador(es) da quadra\n", quadraCount)
	} else {
		fmt.Println("Nao houve acertador para quadra")
	}
}
