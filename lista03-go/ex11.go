package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	V := make([]int, N)
	W := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&V[i])
		W[N-1-i] = V[i]
	}

	fmt.Println(V)
	fmt.Println(W)

	maxV := V[0]
	for _, num := range V {
		if num > maxV {
			maxV = num
		}
	}
	fmt.Println(maxV)

	minW := W[0]
	for _, num := range W {
		if num < minW {
			minW = num
		}
	}
	fmt.Println(minW)
}
