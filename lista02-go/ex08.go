package main

import "fmt"

func main() {
	var n int
	fmt.Print("Digite o número de times no campeonato: ")
	fmt.Scan(&n)

	if n < 2 {
		fmt.Println("Campeonato inválido!")
		return
	}

	fmt.Println("Finais possíveis:")
	k := 1
	for i := 1; i <= n; i++ {
		for j := i + 1; j <= n; j++ {
			fmt.Printf("Final %d: Time%d X Time%d\n", k, i, j)
			k++
		}
	}
}
