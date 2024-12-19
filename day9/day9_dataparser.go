package day9

//import "fmt"

func day9DataParserPart2(rawData string) (DiskDrive, *Stack[DiskEntry]) {
	revDiskFiles := &Stack[DiskEntry]{}
	diskLayout := make(DiskDrive, 0)

	var converted int
	isFile := true
	fileIdx := 0
	spaceIdx := -1
	locationPointer := 0

	var diskEntry DiskEntry
	for _, cmpRune := range rawData {
		converted = int(cmpRune - '0')

		if isFile {
			diskEntry = DiskEntry{
				space:           converted,
				index:           fileIdx,
				locationPointer: locationPointer,
				entryType:       File,
			}
			revDiskFiles.Push(diskEntry)
			fileIdx++
		} else {
			diskEntry = DiskEntry{
				space:           converted,
				index:           spaceIdx,
				locationPointer: locationPointer,
				entryType:       EmptySpace,
			}
			spaceIdx--
		}

		//fmt.Printf("converted: %d, DiskEntry: %+v\n", converted, diskEntry)

		locationPointer += converted
		diskLayout = append(diskLayout, diskEntry)
		isFile = !isFile
	}

	return diskLayout, revDiskFiles
}

//                                   diskdescriptor--file index pushed into the stack
//                                          space on disk
//                                               expanded disk
func day9DataParser(rawData string) (*Stack[int], int, []int) {
	revDiskFiles := &Stack[int]{}
	spaceNeeded := 0
	expanded := make([]int, 0)

	var converted int
	isFile := true
	fileIdx := 0
	for _, cmpRune := range rawData {
		converted = int(cmpRune - '0')

		if isFile {
			spaceNeeded += converted
			for i := 0; i < converted; i++ {
				expanded = append(expanded, fileIdx)
				revDiskFiles.Push(fileIdx) //push the fileIdx onto the stack
			}
			fileIdx++
		} else {
			for i := 0; i < converted; i++ {
				expanded = append(expanded, -1) //use -1 to represent a slack value
			}
		}

		isFile = !isFile
	}

	return revDiskFiles, spaceNeeded, expanded
}
