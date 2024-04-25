package main

import "fmt"

// Função para calcular a soma dos divisores de um número
func sumDivisors(n int) int {
	sum := 0
	for i := 1; i <= n/2; i++ {
		if n%i == 0 {
			sum += i
		}
	}
	return sum
}

// Função para encontrar os pares de números amigos
func findAmicablePairs(n int) [][]int {
	pairs := [][]int{}
	for i := 1; len(pairs) < n; i++ {
		sum1 := sumDivisors(i)
		sum2 := sumDivisors(sum1)
		if sum2 == i && sum1 != sum2 && sum1 < sum2 {
			pair := []int{i, sum1}
			pairs = append(pairs, pair)
		}
	}
	return pairs
}

func main() {
	var n int
	fmt.Println("Digite o número de pares de números amigos que deseja encontrar:")
	fmt.Scan(&n)

	// Encontra os pares de números amigos
	amicablePairs := findAmicablePairs(n)

	// Imprime os pares encontrados
	fmt.Println("Os pares de números amigos são:")
	for _, pair := range amicablePairs {
		fmt.Printf("(%d,%d)\n", pair[0], pair[1])
	}
}
