package main

import (
	"fmt"
)

func readSet(size int) []int {
	set := make([]int, size)
	elements := make(map[int]bool)

	for i := 0; i < size; {
		var element int
		fmt.Scan(&element)

		if !elements[element] {
			set[i] = element
			elements[element] = true
			i++
		} else {
			fmt.Println("Elemento repetido. Por favor, insira um elemento diferente.")
		}
	}

	return set
}

func union(setA, setB []int) []int {
	unionSet := make(map[int]bool)

	for _, element := range setA {
		unionSet[element] = true
	}

	for _, element := range setB {
		unionSet[element] = true
	}

	unionResult := make([]int, 0, len(unionSet))
	for element := range unionSet {
		unionResult = append(unionResult, element)
	}

	return unionResult
}

func intersection(setA, setB []int) []int {
	intersectSet := make(map[int]bool)
	for _, element := range setA {
		intersectSet[element] = true
	}

	intersectResult := make([]int, 0)
	for _, element := range setB {
		if intersectSet[element] {
			intersectResult = append(intersectResult, element)
		}
	}

	return intersectResult
}

func main() {
	var sizeA, sizeB int
	fmt.Println("Informe o tamanho do conjunto A:")
	fmt.Scan(&sizeA)

	fmt.Println("Informe o tamanho do conjunto B:")
	fmt.Scan(&sizeB)

	fmt.Println("Informe os elementos do conjunto A:")
	setA := readSet(sizeA)

	fmt.Println("Informe os elementos do conjunto B:")
	setB := readSet(sizeB)

	fmt.Printf("União: (%v)\n", union(setA, setB))
	fmt.Printf("Interseção: (%v)\n", intersection(setA, setB))
}
