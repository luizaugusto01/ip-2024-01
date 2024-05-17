package main

import (
	"fmt"
	"math"
)

func main() {
	var n, e float64

	fmt.Scan(&e)
	fmt.Scan(&n)

	r := 1.0

	for {
		r_next := (r + n/r) / 2

		erro := math.Abs(n - r_next*r_next)

		fmt.Printf("r: %.9f, erro: %.9f\n", r_next, erro)

		if erro < e {
			break
		}
		r = r_next
	}

}
