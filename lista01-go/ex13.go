package main

import "fmt"

func main() {
	var nota float64

	// Leitura da nota
	fmt.Println("Digite a nota (entre 0 e 10 com uma casa decimal):")
	fmt.Scanln(&nota)

	var conceito string

	// Conversão da nota para conceito
	switch {
	case nota >= 9.0 && nota <= 10:
		conceito = "A"
	case nota >= 7.5 && nota < 9.0:
		conceito = "B"
	case nota >= 6.0 && nota < 7.5:
		conceito = "C"
	case nota >= 0 && nota < 6:
		conceito = "D"
	default:
		fmt.Printf("nota inválida")
		return
	}

	// Impressão da nota e do conceito correspondente
	fmt.Printf("NOTA = %.1f\nCONCEITO = %s", nota, conceito)
}
