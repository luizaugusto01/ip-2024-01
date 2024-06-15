package main

import (
	"fmt"
)

func decimalParaBinario(decimal int) int {
	if decimal == 0 {
		return 0
	}

	binario := 0
	potencia := 1

	for decimal > 0 {
		resto := decimal % 2
		binario += resto * potencia
		potencia *= 10
		decimal /= 2
	}

	return binario
}

func main4() {
	numeroDecimal := 23
	binario := decimalParaBinario(numeroDecimal)

	fmt.Printf("O número decimal %d em binário é %d\n", numeroDecimal, binario)
}
