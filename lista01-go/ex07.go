package main

import "fmt"

func main() {
	var temperaturaFahrenheit, chuvaPolegadas float64

	// Leitura da temperatura em Fahrenheit
	fmt.Println("Digite a temperatura em Fahrenheit:")
	fmt.Scanln(&temperaturaFahrenheit)

	// Leitura da quantidade de chuva em polegadas
	fmt.Println("Digite a quantidade de chuva em polegadas:")
	fmt.Scanln(&chuvaPolegadas)

	// Conversão de Fahrenheit para Celsius
	temperaturaCelsius := (5 * (temperaturaFahrenheit - 32)) / 9

	// Conversão de polegadas para milímetros
	quantidadeChuvaMM := chuvaPolegadas * 25.4

	// Impressão dos resultados
	fmt.Printf("O VALOR EM CELSIUS = %.2f\n", temperaturaCelsius)
	fmt.Printf("A QUANTIDADE DE CHUVA E = %.2f\n", quantidadeChuvaMM)
}
