package main

import (
	"fmt"
)

func mergeSortedArrays(arr1, arr2 []int) []int {
	merged := make([]int, len(arr1)+len(arr2))
	i, j, k := 0, 0, 0

	for i < len(arr1) && j < len(arr2) {
		if arr1[i] < arr2[j] {
			merged[k] = arr1[i]
			i++
		} else {
			merged[k] = arr2[j]
			j++
		}
		k++
	}

	for i < len(arr1) {
		merged[k] = arr1[i]
		i++
		k++
	}

	for j < len(arr2) {
		merged[k] = arr2[j]
		j++
		k++
	}

	return merged
}

func main() {
	var q1, q2 int
	fmt.Scan(&q1, &q2)

	v1 := make([]int, q1)
	v2 := make([]int, q2)

	for i := 0; i < q1; i++ {
		fmt.Scan(&v1[i])
	}

	for i := 0; i < q2; i++ {
		fmt.Scan(&v2[i])
	}

	result := mergeSortedArrays(v1, v2)

	for _, num := range result {
		fmt.Println(num)
	}
}
