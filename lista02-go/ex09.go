package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	_, err := fmt.Scan(&n)

	if err != nil || n <= 1 {
		fmt.Println("Numero invalido!")
		return
	}

	primo := true
	limite := int(math.Sqrt(float64(n))) + 1
	for i := 2; i < limite; i++ {
		if n%i == 0 {
			primo = false
			break
		}
	}

	if primo {
		fmt.Println("PRIMO")
	} else {
		fmt.Println("NAO PRIMO")
	}
}
