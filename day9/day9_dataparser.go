package day9

func day9DataParser(rawData string) ([]int, int) {
	spaceNeeded := 0
	memVals := make([]int, 0)
	var converted int

	for _, cmpRune := range rawData {
		converted = int(cmpRune - '0')

		spaceNeeded += converted
		memVals = append(memVals, converted)
	}

	return memVals, spaceNeeded
}
