package day10

import (
	"bufio"
	"strings"
)

func LineToMapCoordinates(line string, y int) []MapCoordinate {

	mapCoord := make([]MapCoordinate, 0, len(line))

	for x, ht := range line {
		mapCoord = append(mapCoord, MapCoordinate{
			Point: Point{
				X: x,
				Y: y,
			},
			Ht: int(ht - '0'),
		})
	}

	return mapCoord
}

func day10DataParser(rawData string) TopoMap {
	reader := strings.NewReader(rawData)
	scanner := bufio.NewScanner(reader)

	topoMap := TopoMap{
		MapData: make([][]MapCoordinate, 0),
	}

	var y int
	var line string
	for y = 0; scanner.Scan(); y++ {
		line = scanner.Text()
		if line == "" {
			continue
		}

		topoMap.MapData = append(topoMap.MapData, LineToMapCoordinates(line, y))
	}

	topoMap.Dimension = Point{
		X: len(line),
		Y: y,
	}

	return topoMap
}
