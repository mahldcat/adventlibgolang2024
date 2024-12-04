package day4

import (
	"bufio"
	"fmt"
	"strings"
)

func day4DataParser(rawData string) [][]rune {

	reader := strings.NewReader(rawData)
	scanner := bufio.NewScanner(reader)

	grid := make([][]rune, 0) // No need to preallocate unless you know the size
	for scanner.Scan() {
		line := scanner.Text()
		if line != "" { // Skip empty lines (optional, based on your needs)
			grid = append(grid, []rune(line))
		}
	}

	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

	return grid

}
