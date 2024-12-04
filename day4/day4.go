package day4

func findWord(wl [][]rune, x int, y int, xInc int, yInc int, word []rune, wordIdx int) bool {

	if wl[y][x] != word[wordIdx] {
		return false //mismatched rune
	}
	wordIdx++

	if wordIdx == len(word) {
		return true //no more letters to find, things match WOOHOO!!
	}

	y += yInc
	if y < 0 || y == len(wl) {
		return false //OOB on the Y axis (rows)
	}

	x += xInc
	if x < 0 || x == len(wl[y]) {
		return false //OOB on the X axis (cols)
	}

	//verify the next char in the sequence
	return findWord(wl, x, y, xInc, yInc, word, wordIdx)
}

func SolveDay4Part1(rawData string) int {
	wordCt := 0

	wordRune := []rune("XMAS")
	wl := day4DataParser(rawData)

	grid := []int{-1, 0, 1}

	//holy nesting batman!
	for y, row := range wl {
		for x, glyph := range row {
			if glyph == wordRune[0] {
				//check each direction in the 8 poin
				for _, incY := range grid {
					for _, incX := range grid {

						if incX == 0 && incY == 0 { //this is the center of the 3x3
							continue
						}

						if findWord(wl, x, y, incX, incY, wordRune, 0) {
							wordCt++
						}
					}
				}
			}
		}
	}

	return wordCt
}

func SolveDay4Part2(rawData string) int {
	wordCt := 0

	return wordCt
}
