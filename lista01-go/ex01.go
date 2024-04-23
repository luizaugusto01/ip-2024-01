package main

import (
	"fmt"
)

var nota1, nota2, nota3 float64

func main() {
	fmt.Println("Digite a primeira nota: ")
	fmt.Scanln(&nota1)
	fmt.Println("Digite a segunda nota: ")
	fmt.Scanln(&nota2)
	fmt.Println("Digite a terceira nota: ")
	fmt.Scanln(&nota3)

	media := (nota1 + nota2 + nota3) / 3

	fmt.Println("media ", media)

	if media >= 6 {
		fmt.Println("aprovado")
	} else {
		fmt.Println("reprovado")
	}
}
