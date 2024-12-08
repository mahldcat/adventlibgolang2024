package day8

func ContainsPoint(hayStack []Point, needle Point) bool {
	for _, p := range hayStack {
		if p.Equals(needle) {
			return true
		}
	}
	return false
}

func AppendUnique(points []Point, pts []Point) []Point {

	for _, p := range pts {
		if !ContainsPoint(points, p) {
			points = append(points, p)
		}
	}
	return points
}

func SolveDay8Part1(rawData string) int {

	pointMap, gridSize := day8DataParser(rawData)
	collectedPoints := make([]Point, 0)

	for _, points := range pointMap {
		if len(points) == 1 {
			continue
		}
		for i, p1 := range points {
			for _, p2 := range points[i+1:] {

				pts := p1.GetExtensionPoints(p2, gridSize)
				collectedPoints = AppendUnique(collectedPoints, pts)
			}
		}

	}

	return len(collectedPoints)
}

func SolveDay8Part2(rawData string) int {
	pointMap, gridSize := day8DataParser(rawData)
	collectedPoints := make([]Point, 0)

	for _, points := range pointMap {
		if len(points) == 1 {
			continue
		}
		collectedPoints = AppendUnique(collectedPoints, points) //each point should be counted now

		for i, p1 := range points {
			for _, p2 := range points[i+1:] {

				pts := p1.GetExtensionPointsInLine(p2, gridSize)
				collectedPoints = AppendUnique(collectedPoints, pts)
			}
		}
	}

	return len(collectedPoints)
}
