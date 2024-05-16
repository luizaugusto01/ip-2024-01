package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Println("digite o numero inteiro positivo:")
	fmt.Scan(&N)

	if N <= 1 {
		fmt.Println("Núemro invalido ")
		return
	}

	primo := true
	for i := 2; i*i <= N; i++ {
		if N%i == 0 {
			primo = false
			break
		}
	}

	if primo {
		fmt.Println("o número é primo")

	} else {
		fmt.Println("nao primo")
	}
}
