package day11

import (
	"fmt"
	"sort"
)

func SolveDay11Part1(rawData string) int {

	iterations := 25
	computeMap := make(StoneMap)
	data := day11DataParser(rawData)

	fmt.Printf("Solving for stones: %v\n", data)
	solution := 0
	for _, d := range data {
		solution += GetStoneCountDaySln(d, iterations, &computeMap)
	}

	return solution
}

func SolveDay11Part2(rawData string) int {
	iterations := 75
	computeMap := make(StoneMap)
	data := day11DataParser(rawData)

	sort.Ints(data)
	//fmt.Printf("Solving for stones: %v\n", data)
	solution := 0
	for _, d := range data {
		//fmt.Printf("Solving Stone:%d\n", d)
		solution += GetStoneCountDaySln(d, iterations, &computeMap)
	}

	return solution
}
