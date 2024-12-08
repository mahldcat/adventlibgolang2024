package day8

func ContainsPoint(hayStack []Point, needle Point) bool {
	for _, p := range hayStack {
		if p.Equals(needle) {
			return true
		}
	}
	return false
}

func AppendBoundaryUnique(points []Point, p Point, gSz Point) []Point {
	if p.InBounds(gSz) && !ContainsPoint(points, p) {
		points = append(points, p)
	}

	return points
}

func SolveDay8Part1(rawData string) int {

	pointMap, gridSize := day8DataParser(rawData)

	collectedPoints := make([]Point, 0)

	for _, points := range pointMap {

		for i, p1 := range points {
			for _, p2 := range points[i+1:] {

				p3, p4 := p1.GetExtensionPoints(p2)

				collectedPoints = AppendBoundaryUnique(collectedPoints, p3, gridSize)
				collectedPoints = AppendBoundaryUnique(collectedPoints, p4, gridSize)
			}
		}

	}

	return len(collectedPoints)
}

func SolveDay8Part2(rawData string) int {
	solnResult := -1

	return solnResult
}
