package main

import "fmt"

func main() {
	var horas, minutos, segundos int

	// Entrada das horas, minutos e segundos
	fmt.Println("Digite o valor em horas:")
	fmt.Scan(&horas)

	fmt.Println("Digite o valor em minutos:")
	fmt.Scan(&minutos)

	fmt.Println("Digite o valor em segundos:")
	fmt.Scan(&segundos)

	// Conversão para segundos
	tempoEmSegundos := horas*3600 + minutos*60 + segundos

	// Saída do tempo em segundos
	fmt.Printf("O TEMPO EM SEGUNDOS E = %d\n", tempoEmSegundos)
}