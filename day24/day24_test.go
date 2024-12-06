package day24

import (
	"testing"
)

var exampleRaw = ""

func TestParser(t *testing.T) {

}

func TestDay24Part1(t *testing.T) {
	expected := -1
	sln := SolveDay24Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay24Part2(t *testing.T) {
	expected := -1
	sln := SolveDay24Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
