package inputfetcher

import (
	"fmt"
	"strconv"
	"strings"
)

func day1DataParser(rawData string) ([]int, []int, err) {
	var left []int
	var right []int

	// Split the blob into lines
	lines := strings.Split(strings.TrimSpace(rawData), "\n")
	for _, line := range lines {
		// Split each line into parts
		parts := strings.Fields(line)
		if len(parts) != 2 {
			fmt.Printf("Skipping malformed line: %s\n", line)
			continue
		}

		// Parse the numbers
		leftNum, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			continue
		}
		rightNum, err := strconv.Atoi(parts[1])
		if err != nil {
			fmt.Printf("Error parsing number: %v\n", err)
			continue
		}

		// Append to the slices
		left = append(left, leftNum)
		right = append(right, rightNum)
	}

	return left, right, nil
}
