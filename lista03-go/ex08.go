package main

import (
	"fmt"
)

func decimalToBinary(N int) string {
	if N == 0 {
		return "0"
	}

	binary := ""
	for N > 0 {
		bit := N % 2
		binary = fmt.Sprintf("%d%s", bit, binary)
		N /= 2
	}

	return binary
}

func main() {
	for {
		var N int
		_, err := fmt.Scan(&N)
		if err != nil {
			break
		}

		binary := decimalToBinary(N)
		fmt.Println(binary)
	}
}
