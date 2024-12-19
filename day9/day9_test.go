package day9

import (
	"testing"
)

var exampleRaw = "2333133121414131402"

func TestDay9Part1(t *testing.T) {
	expected := 1928
	sln := SolveDay9Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}
func TestDay9Part2(t *testing.T) {
	expected := 2858
	sln := SolveDay9Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}

func TestDay9Part2Checksum(t *testing.T) {
	sampleText := "12345"
	expected := 132

	entries, _ := day9DataParserPart2(sampleText)

	checkSum := entries.CalculateCheckSum()

	if checkSum != expected {
		t.Fatalf("mismatched checksum expected %d, actual %d", expected, checkSum)
	}
}
