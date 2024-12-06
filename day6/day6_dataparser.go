package day6

import (
	"bufio"
	"strings"
)

func day6DataParser(rawData string) ([][]rune, Vector, error) {
	reader := strings.NewReader(rawData)
	scanner := bufio.NewScanner(reader)

	grid := make([][]rune, 0)
	guard := Vector{X: -1, Y: -1, Facing: 'N'}

	for y := 0; scanner.Scan(); y++ {
		line := scanner.Text()
		if line == "" { // Skip empty lines (optional, based on your needs)
			continue
		}

		guardIdx := strings.IndexRune(line, '^')
		if guardIdx != -1 {
			guard.X = guardIdx
			guard.Y = y
		}

		grid = append(grid, []rune(line))

	}

	return grid, guard, scanner.Err()
}
