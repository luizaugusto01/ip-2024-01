package main

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	// Armazenar as notas e a frequência de cada nota
	notas := make(map[int]int)
	var ultimaNota, maiorNota, indicePrimeiraOcorrencia int

	for i := 0; i < N; i++ {
		var nota int
		fmt.Scan(&nota)

		// Atualizar a frequência da última nota informada
		ultimaNota = nota
		notas[nota]++

		// Atualizar a maior nota e sua posição (primeira ocorrência)
		if nota > maiorNota {
			maiorNota = nota
			indicePrimeiraOcorrencia = i
		}
	}

	// Obter a frequência da última nota informada
	frequenciaUltimaNota := notas[ultimaNota]

	// Imprimir a frequência da última nota e a maior nota com sua posição
	fmt.Printf("Nota %d, %d vezes ", ultimaNota, frequenciaUltimaNota)
	fmt.Printf("Nota %d, indice %d\n", maiorNota, indicePrimeiraOcorrencia)
}
