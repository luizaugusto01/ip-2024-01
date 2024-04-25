package main

import "fmt"

func main() {
	var (
		numero, somaPar, somaImpar float64
		contagemPar, contagemImpar int
	)

	// Ler a sequência de números inteiros
	fmt.Println("Digite uma sequência de números inteiros (digite 0 para encerrar):")
	for {
		fmt.Scan(&numero)
		if numero == 0 {
			break // Encerra a entrada de dados
		}

		// Verificar se o número é par ou ímpar e somar à soma correspondente
		if numero%2 == 0 {
			somaPar += numero
			contagemPar++
		} else {
			somaImpar += numero
			contagemImpar++
		}
	}

	// Calcular as médias dos números pares e ímpares
	mediaPar := somaPar / float64(contagemPar)
	mediaImpar := somaImpar / float64(contagemImpar)

	// Imprimir as médias com até duas casas decimais
	fmt.Printf("MEDIA PAR = %.2f\n", mediaPar)
	fmt.Printf("MEDIA IMPAR = %.2f\n", mediaImpar)
}
