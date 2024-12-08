package day8

import (
	"bufio"
	"strings"
)

/*
............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............

ignore the .
map[rune][]Vector

*/

func day8DataParser(rawData string) (map[rune][]Point, Point) {
	reader := strings.NewReader(rawData)
	scanner := bufio.NewScanner(reader)

	pointMap := make(map[rune][]Point)

	var y int
	var line string
	for y = 0; scanner.Scan(); y++ {
		line = scanner.Text()
		if line == "" { // Skip empty lines
			continue
		}

		for x, val := range line {
			if val != '.' {
				p := Point{X: x, Y: y}
				pointMap[val] = append(pointMap[val], p)
			}
		}
	}
	return pointMap, Point{X: len(line), Y: y}
}
