package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Criar um scanner para ler a entrada
	scanner := bufio.NewScanner(os.Stdin)

	// Ler o primeiro valor que é o número de elementos
	scanner.Scan()
	n, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Erro ao converter o número de elementos:", err)
		return
	}

	// Ler a segunda linha que contém os valores
	scanner.Scan()
	valoresStr := strings.Fields(scanner.Text())

	// Verificar se o número de valores corresponde a n
	if len(valoresStr) != n {
		fmt.Println("Número de elementos não corresponde a n")
		return
	}

	// Converter os valores para inteiros e somá-los
	soma := 0
	for _, valorStr := range valoresStr {
		valor, err := strconv.Atoi(valorStr)
		if err != nil {
			fmt.Println("Erro ao converter um valor:", err)
			return
		}
		soma += valor
	}

	// Imprimir a soma dos valores
	fmt.Println(soma)
}
