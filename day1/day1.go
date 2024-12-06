package day1

import (
	"errors"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func toFreq(slice []int) map[int]int {
	freqMap := make(map[int]int)

	// Calculate frequencies
	for _, num := range slice {
		freqMap[num]++
	}

	return freqMap
}

func SolveDay1Part1(rawData string) (int, error) {
	left, right := day1DataParser(rawData)

	if len(left) != len(right) {
		return -1, errors.New("left and right slices must be the same length")
	}
	sort.Ints(left)
	sort.Ints(right)

	fullDist := 0
	for i := 0; i < len(left); i++ {
		fullDist += abs(left[i] - right[i])
	}

	return fullDist, nil
}

func SolveDay1Part2(left []int, right []int) (int, error) {

	fMap := toFreq(right)

	distFreqCalc := 0

	for _, num := range left {

		distFreqCalc += (num * fMap[num])
	}

	return distFreqCalc, nil
}
