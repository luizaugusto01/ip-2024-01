package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	data := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&data[i])
	}

	sort.Ints(data)

	var median float64
	if N%2 == 1 {
		median = float64(data[N/2])
	} else {
		median = float64(data[N/2-1]+data[N/2]) / 2.0
	}

	fmt.Printf("%.2f\n", median)
}
