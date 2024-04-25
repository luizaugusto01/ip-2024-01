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
		fmt.Printf("%d*%d*%d =", n, n, n)
		for i := 0; i < n; i++ {
			odd := 2*i + 1
			sum += odd
			if i > 0 {
				fmt.Printf(" +")
			}
			fmt.Printf(" %d", odd)
		}
		fmt.Printf(" = %d\n", sum)
	}
}
