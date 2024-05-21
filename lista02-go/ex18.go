package main

import (
	"fmt"
)

func gcd(x, y int) int {
	for y != 0 {
		x, y = y, x%y
	}
	return x
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func calculateMMC(a, b, c int) int {
	lcmAB := lcm(a, b)
	lcmTotal := lcm(lcmAB, c)
	return lcmTotal
}

func main() {
	var a, b, c int
	fmt.Println("Digite três números inteiros diferentes de zero:")
	fmt.Scan(&a, &b, &c)

	mmc := calculateMMC(a, b, c)
	fmt.Printf("MMC: %d\n", mmc)
}
