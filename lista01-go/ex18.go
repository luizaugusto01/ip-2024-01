package main

import "fmt"

func main() {
	var x, y int

	// Entrada dos números inteiros x e y
	fmt.Print("Digite dois números inteiros separados por espaço: ")
	fmt.Scan(&x, &y)

	// Verificação se x é par
	if x%2 == 0 {
		// Imprimir sequência de números pares a partir de x
		for i := 0; i < y; i++ {
			fmt.Printf("%d ", x)
			x += 2
		}
		fmt.Println()
	} else {
		// Imprimir mensagem se x não for par
		fmt.Println("O PRIMEIRO NUMERO NAO E PAR")
	}
}