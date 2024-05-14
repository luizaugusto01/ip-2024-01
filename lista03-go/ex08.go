package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// Criar um scanner para ler a entrada
	scanner := bufio.NewScanner(os.Stdin)

	// Ler até o final do arquivo
	for scanner.Scan() {
		// Ler a linha e converter para inteiro
		line := scanner.Text()
		number, err := strconv.Atoi(line)
		if err != nil {
			fmt.Println("Erro ao converter o número:", err)
			continue
		}

		// Converter o número inteiro para binário
		binary := strconv.FormatInt(int64(number), 2)

		// Imprimir o valor binário
		fmt.Println(binary)
	}

	// Verificar se ocorreu algum erro durante a leitura
	if err := scanner.Err(); err != nil {
		fmt.Println("Erro durante a leitura:", err)
	}
}
