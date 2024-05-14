package main

import "fmt"

func main() {
	var N int
	fmt.Scanln(&N)

	notas := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&notas[i])
	}

	// Última nota informada
	ultimaNota := notas[N-1]

	// Frequência da última nota
	frequencia := 0
	for _, nota := range notas {
		if nota == ultimaNota {
			frequencia++
		}
	}

	// Maior nota e sua primeira ocorrência
	maiorNota := notas[0]
	primeiraOcorrencia := 0
	for i, nota := range notas {
		if nota > maiorNota {
			maiorNota = nota
		}
		if nota == maiorNota && i < primeiraOcorrencia {
			primeiraOcorrencia = i
		}
	}

	// Exibição dos resultados
	fmt.Printf("%d\n", frequencia)
	fmt.Printf("%d %d\n", maiorNota, primeiraOcorrencia)
}
