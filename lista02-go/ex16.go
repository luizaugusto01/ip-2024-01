package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)

	for h := 1; h <= n; h++ {
		for c1 := 1; c1 < h; c1++ {
			c2 := math.Sqrt(float64(h*h - c1*c1))
			if c2 == math.Floor(c2) {
				fmt.Printf("hipotenusa = %d, catetos %d e %d\n", h, c1, int(c2))
			}
		}
	}
}
