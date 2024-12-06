package day6

import (
	"testing"
)

var exampleRaw = ""

func TestParser(t *testing.T) {

}

func TestDay6Part1(t *testing.T) {
	expected := -1
	sln := SolveDay6Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay6Part2(t *testing.T) {
	expected := -1
	sln := SolveDay6Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
