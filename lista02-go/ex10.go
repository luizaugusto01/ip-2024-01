package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		// Lendo os dados de entrada
		var matricula int
		var horasTrabalhadas, valorHora float64

		fmt.Print("Digite a matrícula, horas trabalhadas e valor por hora (separados por espaço): ")
		_, err := fmt.Scan(&matricula, &horasTrabalhadas, &valorHora)
		if err != nil {
			fmt.Println("Erro ao ler os dados:", err)
			return
		}

		// Verificando se a matrícula é zero (fim da entrada)
		if matricula == 0 {
			break
		}

		// Calculando o salário
		salario := horasTrabalhadas * valorHora

		// Imprimindo o resultado
		fmt.Printf("Matrícula: %d Salário: %.2f\n", matricula, salario)

		// Consumindo o caractere de quebra de linha da entrada
		if scanner.Scan() {
			scanner.Text()
		}
	}
}
