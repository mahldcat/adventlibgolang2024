package day7

import (
	"testing"
)

var exampleRaw = `190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20`

func TestParser(t *testing.T) {

	result := day7DataParser(exampleRaw)

	expectedLength := 9
	if len(result) != expectedLength {
		t.Fatalf("unexpected length of equation list expected %d, actual %d", expectedLength, len(result))
	}

	expectedListLength := []int{3, 4, 3, 3, 5, 4, 4, 5, 5}
	for i, list := range result {
		if len(list) != expectedListLength[i] {
			t.Fatalf("unexpected list length entry[%d] expected %d, actual %d", i, expectedListLength[i], len(list))
		}
	}

}

func TestDay7Part1(t *testing.T) {
	expected := 3749
	sln := SolveDay7Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay7Part2(t *testing.T) {
	expected := 11387
	sln := SolveDay7Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
