package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Scanner para ler a entrada
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Ler o tamanho do vetor N
		scanner.Scan()
		nStr := scanner.Text()
		n, err := strconv.Atoi(nStr)
		if err != nil {
			fmt.Println("Erro ao converter N:", err)
			return
		}

		// Encerrar o loop se N for 0
		if n == 0 {
			break
		}

		// Ler os valores do vetor
		scanner.Scan()
		valoresStr := strings.Fields(scanner.Text())

		// Verificar se o número de valores corresponde a N
		if len(valoresStr) != n {
			fmt.Println("Número de elementos não corresponde a N")
			return
		}

		// Converter os valores para inteiros
		valores := make([]int, n)
		for i, valorStr := range valoresStr {
			valor, err := strconv.Atoi(valorStr)
			if err != nil {
				fmt.Println("Erro ao converter um valor:", err)
				return
			}
			valores[i] = valor
		}

		// Encontrar o maior valor no vetor
		maior := 0
		for _, valor := range valores {
			if valor > maior {
				maior = valor
			}
		}

		// Contar a frequência dos valores menores ou iguais a cada valor de i
		frequencias := make([]int, maior+1)
		for _, valor := range valores {
			for i := 0; i <= valor; i++ {
				frequencias[i]++
			}
		}

		// Imprimir a saída no formato desejado
		for i := 0; i <= maior; i++ {
			fmt.Printf("(%d) %d\n", i, frequencias[i])
		}
	}
}
