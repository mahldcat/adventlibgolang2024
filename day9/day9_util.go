package day9

import (
	"fmt"
	"sort"
	"strings"
)

type FileType int

const (
	EmptySpace FileType = iota
	File
)

type DiskEntry struct {
	space           int      //number of units
	index           int      //file id
	locationPointer int      //starts at
	entryType       FileType //(0)space (1)file  [technically we don't need this as we are indexing space with neg indices]
}

func (file *DiskEntry) ToString() string {
	return fmt.Sprintf("%+v", file)
}

func (file *DiskEntry) FileCheckSum() int {
	checkSum := 0

	for i := 0; i < file.space; i++ {
		checkSum += file.index * (file.locationPointer + i)

	}
	return checkSum
}

type DiskDrive []DiskEntry

func (drive *DiskDrive) CalculateCheckSum() int {
	checkSum := 0
	/*
		12345
		0..111....22222
		        10  11 12 13 14
		0
	*/
	for _, d := range *drive {
		if d.entryType == File {
			checkSum += d.FileCheckSum()
		}
	}

	return checkSum
}

func (drive *DiskDrive) DiskEntriesToString() string {

	var builder strings.Builder

	for _, d := range *drive {
		for i := 0; i < d.space; i++ {
			if d.entryType == File {
				builder.WriteString(fmt.Sprintf("%d", d.index))
			} else {
				builder.WriteString(".")
			}
		}
	}
	return builder.String()
}

func MoveFileIndex(emptySpace DiskEntry, candidate DiskEntry) (DiskDrive, bool) {
	newFiles := make([]DiskEntry, 0)
	migrated := false

	if emptySpace.space >= candidate.space { //can be moved
		candidate.locationPointer = emptySpace.locationPointer
		newFiles = append(newFiles, candidate)

		newEmpty := DiskEntry{
			space:           emptySpace.space - candidate.space, //it's ok to have an empty space
			locationPointer: candidate.locationPointer + candidate.space,
			index:           emptySpace.index,
			entryType:       EmptySpace,
		}
		newFiles = append(newFiles, newEmpty)
		migrated = true
	}
	return newFiles, migrated
}

func DedupeAndSort(disk DiskDrive) DiskDrive {

	seenByIndex := make(map[int]bool)
	deduped := make(DiskDrive, 0, len(disk))

	for _, d := range disk {

		if !seenByIndex[d.index] {
			seenByIndex[d.index] = true
			if d.entryType == EmptySpace && d.space == 0 {
				continue
			}
			deduped = append(deduped, d)
		}
	}

	sort.Slice(deduped, func(i, j int) bool {
		return deduped[i].locationPointer < deduped[j].locationPointer
	})

	return deduped
}

func DefragIndexByFile(currentDisk DiskDrive, reverseFile Stack[DiskEntry]) int {
	fmt.Println("Starting")
	//fmt.Printf("%s\n", currentDisk.DiskEntriesToString())

	candidateDisk := make(DiskDrive, 0)
	candidateDiskChanged := false
	for {
		file, _ := reverseFile.Pop() // Pop the top of the stack
		//fmt.Printf("Working on file[%d] : %s\n", file.index, file.ToString())
		if file.index == 0 {
			break // Stop when index is 0
		}

		candidateDisk = candidateDisk[:0]
		candidateDiskChanged = false

		for _, possibleEmptySpace := range currentDisk {
			if possibleEmptySpace.entryType == File || //candidate is a file...skip it
				possibleEmptySpace.space < file.space || //candidate space not large enough to hold the file
				possibleEmptySpace.locationPointer > file.locationPointer { //we have gone past the file location on disk...no need to procceed
				continue
			}

			//fmt.Println("Entering the MoveFileToIndex logic")
			newFiles, moved := MoveFileIndex(possibleEmptySpace, file)

			//fmt.Println("split files")

			if moved {
				/*
					fmt.Printf("Moved File: %s", newFiles.DiskEntriesToString())
					for _, nf := range newFiles {
						fmt.Printf("splitFile: %s\n", nf.ToString())
					}
				*/
				candidateDisk = append(candidateDisk, newFiles...)
				candidateDisk = append(candidateDisk, currentDisk...)
				candidateDiskChanged = true
				break
			}
		}
		if candidateDiskChanged {
			//fmt.Println("Deduping/sorting disk")
			//fmt.Printf("   Base: %s\n", candidateDisk.DiskEntriesToString())
			currentDisk = DedupeAndSort(candidateDisk)
			//fmt.Printf("   Ddup: %s\n", currentDisk.DiskEntriesToString())
		}
	}

	//fmt.Printf("Final  Result: %s\n", currentDisk.DiskEntriesToString())
	return currentDisk.CalculateCheckSum()
}

func DefragIndex(rawDisk []int, reverseFiles Stack[int], sz int) int {
	checkSum := 0
	val := 0

	for i := 0; i < sz; i++ {
		if rawDisk[i] == -1 {
			val, _ = reverseFiles.Pop()
			checkSum += (val * i)
		} else {
			checkSum += (rawDisk[i] * i)
		}
	}

	return checkSum
}

/**************************************************************************/
/*                            STACK                                       */
/**************************************************************************/

// Stack represents a stack data structure
type Stack[T any] struct {
	data []T
}

// Push adds an element to the top of the stack
func (s *Stack[T]) Push(value T) {
	s.data = append(s.data, value)
}

// Pop removes and returns the top element of the stack
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var zeroVal T
		return zeroVal, false // Return zero value and false if stack is empty
	}
	val := s.data[len(s.data)-1]    // Get the top element
	s.data = s.data[:len(s.data)-1] // Remove the top element
	return val, true
}

// Peek returns the top element without removing it
func (s *Stack[T]) Peek() (T, bool) {
	if len(s.data) == 0 {
		var zeroVal T
		return zeroVal, false // Return zero value and false if stack is empty
	}
	return s.data[len(s.data)-1], true
}

// IsEmpty checks if the stack is empty
func (s *Stack[T]) IsEmpty() bool {
	return len(s.data) == 0
}

// Size returns the number of elements in the stack
func (s *Stack[T]) Size() int {
	return len(s.data)
}
