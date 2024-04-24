package main

import "fmt"

var a, b, c float64

func main() {

	// Leitura dos coeficientes A, B e C
	fmt.Println("Digite o coeficiente A:")
	fmt.Scanln(&a)
	fmt.Println("Digite o coeficiente B:")
	fmt.Scanln(&b)
	fmt.Println("Digite o coeficiente C:")
	fmt.Scanln(&c)

	// Cálculo do discriminante
	delta := b*b - 4*a*c

	// Impressão do valor do discriminante
	fmt.Printf("O VALOR DE DELTA E = %.2f\n", delta)
}
