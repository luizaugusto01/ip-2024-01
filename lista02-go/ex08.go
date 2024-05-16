package main

import "fmt"

func main() {
	var N int
	fmt.Println("Digite a quantidade de times:")
	fmt.Scan(&N)

	if N < 2 {
		fmt.Println("Campeonato Inváido!")
		return
	}

	count := 1
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			TimeA := string('A' + i)
			TimeB := string('A' + j)
			fmt.Printf("Final %d: Time %s contra Time %s\n", count, TimeA, TimeB)
			count++
		}
	}
}
