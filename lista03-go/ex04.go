package main

import (
	"fmt"
)

func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil || n >= 5000 {
		return
	}

	values := make([]int, n)
	for i := 0; i < n; i++ {
		_, err := fmt.Scan(&values[i])
		if err != nil {
			return
		}
	}

	for i := 0; i < n; i++ {
		fmt.Print(values[i])
		if i != n-1 {
			fmt.Print(" ")
		}
	}
	fmt.Println()
}
