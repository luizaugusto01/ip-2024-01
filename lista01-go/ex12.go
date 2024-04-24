package main

import "fmt"

func main() {
	var horas int

	// Leitura do número de horas de uso da charrete
	fmt.Println("Digite a quantidade de horas de uso da charrete:")
	fmt.Scanln(&horas)

	resto := horas % 3
	// Cálculo do valor a pagar
	var valorAPagar float64
	if resto == 0 {
		valorAPagar = float64(horas / 3 * 10)
	} else {

		valorAPagar = float64((resto * 5) + int((horas/3)*10))
	}

	// Impressão do valor a pagar
	fmt.Printf("O VALOR A PAGAR E = %.2f\n", valorAPagar)
}
