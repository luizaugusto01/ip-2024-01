package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	var input string
	fmt.Scanln(&input)

	// Verifica se o formato é válido
	parts := strings.Split(input, ";")
	if len(parts) != 2 {
		fmt.Println("FORMATO INVALIDO!")
		return
	}

	A := parts[0]
	B := parts[1]

	// Calcula os vetores de frequência de vogais para cada string
	freqA := vowelFrequency(A)
	freqB := vowelFrequency(B)

	// Imprime os vetores de frequência
	fmt.Printf("(%d,%d,%d,%d,%d)\n", freqA[0], freqA[1], freqA[2], freqA[3], freqA[4])
	fmt.Printf("(%d,%d,%d,%d,%d)\n", freqB[0], freqB[1], freqB[2], freqB[3], freqB[4])

	// Calcula a distância euclidiana entre os dois vetores de frequência
	distance := euclideanDistance(freqA, freqB)
	fmt.Printf("%.2f\n", distance)
}

// Função para calcular a frequência das vogais em uma string
func vowelFrequency(s string) [5]int {
	var freq [5]int
	s = strings.ToLower(s)
	for _, char := range s {
		switch char {
		case 'a':
			freq[0]++
		case 'e':
			freq[1]++
		case 'i':
			freq[2]++
		case 'o':
			freq[3]++
		case 'u':
			freq[4]++
		}
	}
	return freq
}

// Função para calcular a distância euclidiana entre dois vetores de frequência
func euclideanDistance(freqA, freqB [5]int) float64 {
	var sum float64
	for i := 0; i < 5; i++ {
		diff := float64(freqA[i] - freqB[i])
		sum += diff * diff
	}
	return math.Sqrt(sum)
}
