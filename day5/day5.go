package day5

import "fmt"

func SolveDay5Part1(rawData string) int {
	midPointSum := 0

	rules, updates, err := day5DataParser(rawData)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return -1
	}

	for _, updateEntry := range updates {
		if verifyList(rules, updateEntry) {
			midPointSum += getMidPointValue(updateEntry)
		}

	}
	return midPointSum
}

func SolveDay5Part2(rawData string) int {
	midPointSum := 0

	rules, updates, err := day5DataParser(rawData)

	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return -1
	}

	for _, updateEntry := range updates {
		if !verifyList(rules, updateEntry) {
			repairUpdateOrder(rules, updateEntry)
			midPointSum += getMidPointValue(updateEntry)
		}
	}

	return midPointSum
}
