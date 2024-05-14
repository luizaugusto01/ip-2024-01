package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

// Function to generate combinations of n numbers out of m
func combinations(arr []int, n int) [][]int {
	var result [][]int
	var combine func([]int, int, []int)
	combine = func(temp []int, start int, current []int) {
		if len(current) == n {
			comb := make([]int, n)
			copy(comb, current)
			result = append(result, comb)
			return
		}
		for i := start; i < len(arr); i++ {
			current = append(current, arr[i])
			combine(temp, i+1, current)
			current = current[:len(current)-1]
		}
	}
	combine(arr, 0, []int{})
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	t, _ := strconv.Atoi(scanner.Text())

	for caseNum := 0; caseNum < t; caseNum++ {
		dwarfs := make([]int, 9)
		for i := 0; i < 9; i++ {
			scanner.Scan()
			num, _ := strconv.Atoi(scanner.Text())
			dwarfs[i] = num
		}

		combs := combinations(dwarfs, 7)
		for _, comb := range combs {
			sum := 0
			for _, num := range comb {
				sum += num
			}
			if sum == 100 {
				sort.Ints(comb)
				for _, num := range comb {
					fmt.Println(num)
				}
				break
			}
		}
	}
}
