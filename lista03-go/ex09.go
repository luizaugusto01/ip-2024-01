package main

import (
	"fmt"
	"math"
)

type Point struct {
	X float64
	Y float64
	Z float64
}

func distance(p1, p2 Point) float64 {
	dx := p1.X - p2.X
	dy := p1.Y - p2.Y
	dz := p1.Z - p2.Z
	return math.Sqrt(dx*dx + dy*dy + dz*dz)
}

func main() {
	var N int
	fmt.Scanln(&N)

	points := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Scanln(&points[i].X, &points[i].Y, &points[i].Z)
	}

	for i := 1; i < N; i++ {
		d := distance(points[i-1], points[i])
		fmt.Printf("%.2f\n", d)
	}
}
