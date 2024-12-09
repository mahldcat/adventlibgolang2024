package day9

func SolveDay9Part1(rawData string) int {
	compressedFiles, szExpanded := day9DataParser(rawData)

	uncompacted := GetUncompacted(compressedFiles, szExpanded)

	Defragment(uncompacted)

	//fmt.Printf("Defrag: %v\n", uncompacted)

	return GetCheckSum(uncompacted)
}

func SolveDay9Part2(rawData string) int {
	solnResult := -1

	return solnResult
}
