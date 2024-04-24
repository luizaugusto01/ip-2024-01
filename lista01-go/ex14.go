package main

import (
	"fmt"
	"math"
)

func main() {
	var altura, aresta float64

	// Leitura da altura e da aresta do hexágono
	fmt.Println("Digite a altura da pirâmide e o comprimento da aresta do hexágono (em metros):")
	fmt.Scanln(&altura, &aresta)

	// Cálculo da área da base do hexágono
	areaBase := 3 * math.Pow(aresta, 2) * math.Sqrt(3)

	// Cálculo do volume da pirâmide
	volume := (1.0 / 3.0) * areaBase * altura

	// Impressão do volume da pirâmide
	fmt.Printf("O VOLUME DA PIRAMIDE E = %.2f METROS CUBICOS", volume)
}
