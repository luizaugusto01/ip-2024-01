package main

import (
	"fmt"
)

func main() {
	var ValorIngresso, ValorInicial, ValorFinal float64

	fmt.Scan(&ValorIngresso)
	fmt.Scan(&ValorInicial)
	fmt.Scan(&ValorFinal)

	if ValorInicial >= ValorFinal {
		fmt.Println("INTERVALO INVALIDO.")
		return
	}

	type Resultado struct {
		valorIngresso float64
		quantidade    int
		lucro         float64
	}

	var melhorResultado Resultado

	for v := ValorInicial; v <= ValorFinal; v++ {
		quantidadeVendida := 120 + int((ValorIngresso-v)*50)
		receita := v * float64(quantidadeVendida)
		custo := 200.00 + 0.05*float64(quantidadeVendida)
		lucro := receita - custo

		fmt.Printf("V: %.2f, N: %d, L: %.2f\n", v, quantidadeVendida, lucro)

		if lucro > melhorResultado.lucro {
			melhorResultado = Resultado{
				valorIngresso: v,
				quantidade:    quantidadeVendida,
				lucro:         lucro,
			}
		}
	}

	if melhorResultado.lucro > 0 {
		fmt.Printf("Melhor valor final: %.2f\n", melhorResultado.valorIngresso)
		fmt.Printf("Lucro: %.2f\n", melhorResultado.lucro)
		fmt.Printf("Numero de ingressos: %d\n", melhorResultado.quantidade)
	} else {
		fmt.Printf("Melhor valor final: 0.00\n")
		fmt.Printf("Lucro: 0.00\n")
		fmt.Printf("Numero de ingressos: 0\n")
	}
}
