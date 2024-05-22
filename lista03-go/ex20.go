package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y, z float64
}

func main() {
	var N int
	fmt.Scan(&N)

	points := make([]Point, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&points[i].x, &points[i].y, &points[i].z)
	}

	maxCoord := 0.0
	for i := 0; i < N-1; i++ {
		vector := calculateVector(points[i], points[i+1])
		maxCoord = math.Max(maxCoord, math.Abs(vector.x))
		maxCoord = math.Max(maxCoord, math.Abs(vector.y))
		maxCoord = math.Max(maxCoord, math.Abs(vector.z))
	}

	fmt.Printf("%.2f\n", maxCoord)
}

func calculateVector(a, b Point) Point {
	return Point{b.x - a.x, b.y - a.y, b.z - a.z}
}
