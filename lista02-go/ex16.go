package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo: ")
	fmt.Scan(&n)

	if n <= 0 {
		fmt.Println("O número deve ser maior que zero.")
		return
	}

	for h := 1; h <= n; h++ {
		for c1 := 1; c1 <= h; c1++ {
			c2 := math.Sqrt(float64(h*h - c1*c1))
			if c2 == math.Trunc(c2) && c2 <= float64(n) {
				fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, c1, int(c2))
			}
		}
	}
}
