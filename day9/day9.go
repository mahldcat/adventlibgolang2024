package day9

func SolveDay9Part1(rawData string) int {
	reverseFile, sizeCompacted, expandedDisk := day9DataParser(rawData)

	return DefragIndex(expandedDisk, *reverseFile, sizeCompacted)
}

func SolveDay9Part2(rawData string) int {
	diskEntries, reverserawData := day9DataParserPart2(rawData)

	return DefragIndexByFile(diskEntries, *reverserawData)
}
