package day10

func SolveDay10Part1(rawData string) int {
	topoMap := day10DataParser(rawData)

	peak1, _ := topoMap.GetPathCounts()
	return peak1
}

func SolveDay10Part2(rawData string) int {
	topoMap := day10DataParser(rawData)

	_, peak2 := topoMap.GetPathCounts()
	return peak2
}
