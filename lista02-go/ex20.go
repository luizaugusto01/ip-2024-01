package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Println("Digite um número inteiro maior que 1:")
	fmt.Scan(&n)

	sum := 0
	divisors := make([]int, 0)

	for i := 1; i < n; i++ {
		if n%i == 0 {
			sum += i
			divisors = append(divisors, i)
		}
	}

	fmt.Printf("%d =", n)
	for i, divisor := range divisors {
		fmt.Printf(" %d", divisor)
		if i < len(divisors)-1 {
			fmt.Printf(" +")
		}
	}
	fmt.Printf(" = %d ", sum)

	if sum == n {
		fmt.Println("(NUMERO PERFEITO)")
	} else {
		fmt.Println("(NUMERO NAO E PERFEITO)")
	}
}
