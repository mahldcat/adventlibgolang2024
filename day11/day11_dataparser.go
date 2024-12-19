package day11

import (
	"strconv"
	"strings"
)

func day11DataParser(rawData string) []int {
	stones := make([]int, 0)

	for _, stoneRaw := range strings.Split(rawData, " ") {
		stone, _ := strconv.Atoi(strings.TrimSpace(stoneRaw))
		stones = append(stones, stone)
	}

	return stones
}
