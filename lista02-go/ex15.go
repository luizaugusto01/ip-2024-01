package main

import "fmt"

func main() {
	var T, A, B int
	fmt.Scan(&T)

	for i := 0; i < T; i++ {
		fmt.Scan(&A, &B)

		// Inverte os dígitos de A e B
		A = reverseNumber(A)
		B = reverseNumber(B)

		// Determina qual número é maior
		if A > B {
			fmt.Println(reverseNumber(A))
		} else {
			fmt.Println(reverseNumber(B))
		}
	}
	// Determina qual número é maior
	if A > B {
		fmt.Println(reverseNumber(A))
	} else {
		fmt.Println(reverseNumber(B))
	}
}

// Função para inverter os dígitos de um número
func reverseNumber(n int) int {
	var reversed int
	for n > 0 {
		digit := n % 10
		reversed = reversed*10 + digit
		n /= 10
	}
	return reversed
}
