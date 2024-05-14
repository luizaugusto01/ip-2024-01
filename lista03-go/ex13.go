package main

import (
	"fmt"
)

func main() {
	var N int
	fmt.Scan(&N)

	freq := make(map[int]int)

	var maxFreq, mostFreq int
	for i := 0; i < N; i++ {
		var num int
		fmt.Scan(&num)
		freq[num]++
		if freq[num] > maxFreq || (freq[num] == maxFreq && num < mostFreq) {
			mostFreq = num
			maxFreq = freq[num]
		}
	}

	fmt.Println(mostFreq)
	fmt.Println(maxFreq)

	fmt.Println()
}
