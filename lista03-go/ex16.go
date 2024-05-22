package main

import (
	"fmt"
	"sort"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	arrivalTimes := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&arrivalTimes[i])
	}

	sort.Ints(arrivalTimes)

	countOnTime := 0
	for _, time := range arrivalTimes {
		if time <= 0 {
			countOnTime++
		}
	}

	if countOnTime < K {
		fmt.Println("SIM")
	} else {
		fmt.Println("NAO")
		for i := countOnTime - 1; i >= 0; i-- {
			if arrivalTimes[i] <= 0 {
				fmt.Printf("%d ", i+1)
			}
		}
		fmt.Println()
	}
}
