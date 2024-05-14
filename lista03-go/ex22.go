package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	for scanner.Scan() {
		line := scanner.Text()
		if line == "0 0" {
			break
		}

		var n, d int
		fmt.Sscanf(line, "%d %d", &n, &d)
		scanner.Scan()
		num := scanner.Text()

		stack := make([]byte, 0, d)
		toRemove := n - d

		for i := 0; i < n; i++ {
			for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < num[i] {
				stack = stack[:len(stack)-1]
				toRemove--
			}
			stack = append(stack, num[i])
		}

		result := stack[:d]
		writer.WriteString(string(result) + "\n")
	}
}
