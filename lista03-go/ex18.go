package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Leitura do número de casos de teste
	scanner.Scan()
	T, _ := strconv.Atoi(scanner.Text())

	results := make([]string, T)

	for t := 0; t < T; t++ {
		// Leitura de cada CPF
		scanner.Scan()
		cpfLine := scanner.Text()
		cpfParts := strings.Split(cpfLine, " ")

		// Convertendo a string para um slice de inteiros
		cpf := make([]int, 11)
		for i := 0; i < 11; i++ {
			cpf[i], _ = strconv.Atoi(cpfParts[i])
		}

		// Calculando o primeiro dígito verificador (b1)
		sum1 := 0
		for i := 0; i < 9; i++ {
			sum1 += cpf[i] * (i + 1)
		}
		b1 := sum1 % 11
		if b1 == 10 {
			b1 = 0
		}

		// Calculando o segundo dígito verificador (b2)
		sum2 := 0
		for i := 0; i < 9; i++ {
			sum2 += cpf[i] * (9 - i)
		}
		b2 := sum2 % 11
		if b2 == 10 {
			b2 = 0
		}

		// Verificando se o CPF é válido
		if b1 == cpf[9] && b2 == cpf[10] {
			results[t] = "CPF valido"
		} else {
			results[t] = "CPF invalido"
		}
	}

	// Imprimindo os resultados
	for _, result := range results {
		fmt.Println(result)
	}
}
