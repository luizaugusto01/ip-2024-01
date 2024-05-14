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

	// Leitura da primeira linha que contém N e K
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Split(firstLine, " ")
	N, _ := strconv.Atoi(parts[0])
	K, _ := strconv.Atoi(parts[1])

	// Leitura da segunda linha que contém os tempos de chegada
	scanner.Scan()
	secondLine := scanner.Text()
	arrivalTimes := strings.Split(secondLine, " ")

	// Contagem dos alunos presentes antes do início da aula
	presentCount := 0
	earlyStudents := []int{}

	for i := 0; i < N; i++ {
		arrivalTime, _ := strconv.Atoi(arrivalTimes[i])
		if arrivalTime <= 0 {
			presentCount++
			earlyStudents = append(earlyStudents, i+1) // Guardar o índice do aluno (i+1)
		}
	}

	// Verificação se a aula será cancelada
	if presentCount < K {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		// Impressão dos alunos que chegaram antes do início da aula, em ordem reversa
		for i := len(earlyStudents) - 1; i >= 0; i-- {
			fmt.Print(earlyStudents[i])
			if i > 0 {
				fmt.Print(" ")
			}
		}
		fmt.Println()
	}
}
