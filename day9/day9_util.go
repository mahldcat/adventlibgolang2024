package day9

func GetUncompacted(compactedFiles []int, decompressedSize int) []int {
	uncompacted := make([]int, decompressedSize)

	isFile := true //slackSpace is false
	fileIdx := 0
	uIdxOffSet := 0
	for _, cf := range compactedFiles {
		for i := 0; i < cf; i++ {
			if isFile {
				uncompacted[uIdxOffSet+i] = fileIdx
			} else {
				uncompacted[uIdxOffSet+i] = -1
			}
		}

		if isFile {
			fileIdx++
		}
		isFile = !isFile
		uIdxOffSet += cf
	}

	return uncompacted
}

func GetCheckSum(defragged []int) int {

	checkSum := 0
	for i, idx := range defragged {
		//fmt.Printf("checkSum on file[%d]:=%d\n", i, idx)
		if idx == -1 {
			return checkSum
		}
		checkSum += (i * idx)
	}

	return checkSum
}

func Swap(intSlice []int, idx1 int, idx2 int) {
	intSlice[idx1], intSlice[idx2] = intSlice[idx2], intSlice[idx1]
}

//gets the index and value of the last file entry in uncompacted
// starting from startFrom
func GetLastFileIndex(uncompacted []int, startFrom int) int {

	for i := startFrom; i >= 0; i-- {
		if uncompacted[i] > 0 {
			return i
		}
	}

	//should never get here?
	return -1
}

func GetNextSlackIndex(uncompacted []int, startFrom int) int {
	for i := startFrom; i < len(uncompacted); i++ {
		if uncompacted[i] == -1 {
			return i
		}
	}
	return -1
}

func Defragment(uncompacted []int) {

	forwardSlackIdx := GetNextSlackIndex(uncompacted, 0)
	reverseFileIdx := GetLastFileIndex(uncompacted, len(uncompacted)-1)

	//fmt.Printf("Forward:%d Reverse:%d\n", forwardSlackIdx, reverseFileIdx)

	for forwardSlackIdx < reverseFileIdx && forwardSlackIdx != -1 && reverseFileIdx != -1 {
		Swap(uncompacted, forwardSlackIdx, reverseFileIdx)
		forwardSlackIdx = GetNextSlackIndex(uncompacted, forwardSlackIdx+1)
		reverseFileIdx = GetLastFileIndex(uncompacted, reverseFileIdx-1)
		//fmt.Printf("Forward:%d Reverse:%d\n", forwardSlackIdx, reverseFileIdx)
	}
}

/*
"23 33 13 31 21 41 41 31 40 2"

00...111...2...333.44.5555.6666.777.888899
00...111...2...333.44.5555.6666.777.888899
f                                        b

compresses to:
0099811188827773336446555566..............
slack Indexes
{2, 3, 4, 8, 9, 10, 12, 13, 14, 18, 21, 26, 31, 35}

*/
