package day17

import (
	"testing"
)

var exampleRaw = ""

func TestParser(t *testing.T) {

}

func TestDay17Part1(t *testing.T) {
	expected := -1
	sln := SolveDay17Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay17Part2(t *testing.T) {
	expected := -1
	sln := SolveDay17Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
