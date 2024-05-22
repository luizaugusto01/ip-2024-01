package main

import (
	"fmt"
	"math"
)

// Função para calcular o máximo divisor comum (MDC) usando o algoritmo de Euclides
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// Função para converter um número decimal em uma fração simplificada
func decimalToFraction(decimal float64) (int, int) {
	// Obter a parte inteira e a parte fracionária do número
	integralPart := math.Floor(decimal)
	fractionalPart := decimal - integralPart

	// Se não há parte fracionária, é um número inteiro
	if fractionalPart == 0 {
		return int(integralPart), 1
	}

	// Converter a parte fracionária para uma fração
	precision := 1e9 // Precisão alta para evitar problemas de arredondamento
	denominator := int(precision)
	numerator := int(fractionalPart * precision)

	// Simplificar a fração
	divisor := gcd(numerator, denominator)
	numerator /= divisor
	denominator /= divisor

	// Adicionar a parte inteira ao numerador
	numerator += int(integralPart) * denominator

	return numerator, denominator
}

func main() {
	var decimal float64
	fmt.Print("Digite um número decimal: ")
	fmt.Scanln(&decimal)

	// Converter para fração simplificada
	num, den := decimalToFraction(decimal)
	fmt.Printf("%d/%d\n", num, den)
}
