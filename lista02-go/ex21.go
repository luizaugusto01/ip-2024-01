package main

import (
	"fmt"
)

func main() {
	var n int
	for {
		fmt.Scan(&n)
		if n == 0 {
			break
		}

		sequence := make([]float64, n)
		for i := 0; i < n; i++ {
			fmt.Scan(&sequence[i])
		}

		if isOrdered(sequence) {
			fmt.Println("ORDENADA")
		} else {
			fmt.Println("DESORDENADA")
		}
	}
}

func isOrdered(seq []float64) bool {
	for i := 1; i < len(seq); i++ {
		if seq[i-1] >= seq[i] {
			return false
		}
	}
	return true
}
