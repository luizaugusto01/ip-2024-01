package main

import (
	"fmt"
)

func main() {
	for {
		var N int
		_, err := fmt.Scan(&N)
		if err != nil || N == 0 {
			break
		}

		if N <= 1 || N > 10000 {
			return
		}

		values := make([]int, N)
		for i := 0; i < N; i++ {
			_, err := fmt.Scan(&values[i])
			if err != nil {
				return
			}
		}

		maxValue := values[0]
		maxIndex := 0
		for i := 1; i < N; i++ {
			if values[i] > maxValue {
				maxValue = values[i]
				maxIndex = i
			}
		}

		fmt.Printf("%d %d\n", maxIndex, maxValue)
	}
}
