package day8

import (
	"testing"
)

var exampleRaw = ""

func TestParser(t *testing.T) {

}

func TestDay8Part1(t *testing.T) {
	expected := -1
	sln := SolveDay8Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay8Part2(t *testing.T) {
	expected := -1
	sln := SolveDay8Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
