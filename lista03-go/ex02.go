package main

import (
	"fmt"
)

func main() {
	var N int
	for {
		_, err := fmt.Scan(&N)
		if err == nil && N >= 1 && N <= 1000 {
			break
		}
	}

	V := make([]int, N)
	for i := 0; i < N; i++ {
		_, err := fmt.Scan(&V[i])
		if err != nil {
			return
		}
	}

	var K int
	_, err := fmt.Scan(&K)
	if err != nil {
		return
	}

	count := 0
	for _, value := range V {
		if value >= K {
			count++
		}
	}

	fmt.Println(count)
}
