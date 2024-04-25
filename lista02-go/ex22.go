package main

import (
	"fmt"
	"math/big"
)

func main() {
	var num float64
	fmt.Println("Digite um número decimal:")
	fmt.Scan(&num)

	// Converte o número decimal em uma fração
	fraction := big.NewRat(int64(num*10000), 10000)

	// Simplifica a fração
	fraction = fraction.Reduce()

	// Obtém o numerador e o denominador da fração
	numerator := fraction.Num()
	denominator := fraction.Denom()

	// Imprime a fração simplificada
	fmt.Printf("A fração simplificada é: %d/%d\n", numerator,​⬤