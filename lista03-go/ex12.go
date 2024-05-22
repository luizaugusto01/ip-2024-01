package main

import (
	"fmt"
	"sort"
)

func main() {
	var N int
	fmt.Scan(&N)

	nums := make([]int, N)

	for i := 0; i < N; i++ {
		fmt.Scan(&nums[i])
	}

	sort.Ints(nums)

	for _, num := range nums {
		fmt.Println("-")
		fmt.Println(num)
	}

	fmt.Println()
}
