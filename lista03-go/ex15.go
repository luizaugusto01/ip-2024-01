package main

import (
	"fmt"
	"math"
)

func main() {
	var T int
	fmt.Scan(&T)

	for t := 0; t < T; t++ {
		var N int
		fmt.Scan(&N)

		nums := make([]int, N)
		for i := 0; i < N; i++ {
			fmt.Scan(&nums[i])
		}

		minDistance, comparisons := minDistance(nums)
		fmt.Printf("%d %d\n", minDistance, comparisons)
	}
}

func minDistance(nums []int) (int, int) {
	minDist := math.MaxInt64
	comparisons := 0

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			dist := int(math.Abs(float64(nums[i] - nums[j])))
			comparisons++
			if dist < minDist {
				minDist = dist
			}
		}
	}

	return minDist, comparisons
}
