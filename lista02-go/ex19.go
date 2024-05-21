package main

import (
	"fmt"
)

func main() {
	var m int
	fmt.Println("Digite um número inteiro maior que zero:")
	fmt.Scan(&m)

	for n := 1; n <= m; n++ {
		sum := 0
		fmt.Printf("%d*%d*%d = ", n, n, n)
		for i := 0; i < n; i++ {
			odd := (n-1)*n + 1 + 2*i
			sum += odd
			if i < n-1 {
				fmt.Printf("%d+", odd)
			} else {
				fmt.Printf("%d", odd)
			}
		}
		fmt.Println()
	}
}
