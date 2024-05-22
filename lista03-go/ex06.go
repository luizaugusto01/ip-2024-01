package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	if n <= 0 || n >= 5000 {
		return
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&values[i])
	}

	sum := 0
	for _, value := range values {
		sum += value
	}

	fmt.Println(sum)
}
