package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	// Leitura do número de pontos
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	// Leitura dos pontos
	points := make([][3]float64, N)
	for i := 0; i < N; i++ {
		scanner.Scan()
		coords := strings.Split(scanner.Text(), " ")
		x, _ := strconv.ParseFloat(coords[0], 64)
		y, _ := strconv.ParseFloat(coords[1], 64)
		z, _ := strconv.ParseFloat(coords[2], 64)
		points[i] = [3]float64{x, y, z}
	}

	// Processamento dos vetores e cálculo dos módulos das coordenadas
	for i := 0; i < N-1; i++ {
		vx := points[i+1][0] - points[i][0]
		vy := points[i+1][1] - points[i][1]
		vz := points[i+1][2] - points[i][2]

		absVx := math.Abs(vx)
		absVy := math.Abs(vy)
		absVz := math.Abs(vz)

		maxAbs := math.Max(absVx, math.Max(absVy, absVz))

		fmt.Printf("%.2f\n", maxAbs)
	}
}
