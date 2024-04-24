package main

import "fmt"

func main() {
	var x, y int
	fmt.Print("Digite dois números inteiros separados por espaço: ")
	fmt.Scanln(&x, &y)

	if x%2 == 0 { // Verifica se x é par
		for i := 1; i <= y; i++ {
			fmt.Printf("%d ", x)
			x += 2 // Incrementa x para o próximo número par
		}
		fmt.Println() // Adiciona uma nova linha após a sequência
	} else {
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR") // Imprime mensagem se x não for par
	}
}
