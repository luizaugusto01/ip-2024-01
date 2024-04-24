package main

import "fmt"

func main() {
	var a1, r, n int
	fmt.Println("Digite o valor inicial, a razão e o número de elementos da progressão aritmética:")
	fmt.Scanln(&a1, &r, &n)

	soma := 0
	for i := 0; i < n; i++ {
		termo := a1 + i*r // Calcula o i-ésimo termo da progressão
		soma += termo     // Soma o termo à variável soma
		fmt.Println(soma)
	}

	fmt.Printf("A soma dos %d primeiros elementos da progressão aritmética é: %d\n", n, soma)
}
