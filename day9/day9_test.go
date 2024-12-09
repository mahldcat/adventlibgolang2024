package day9

import (
	"testing"
)

var exampleRaw = "2333133121414131402"

func TestParsing(t *testing.T) {
	vals, spaceNeeded := day9DataParser(exampleRaw)

	expectedValCt := 19
	expected := 42

	if spaceNeeded != expected {
		t.Fatalf("mismatched space length expected %d, actual:%d", expected, spaceNeeded)
	}
	if len(vals) != expectedValCt {
		t.Fatalf("mismatched val count, expected %d, actual:%d", expectedValCt, len(vals))
	}
}

func TestUnCompacted(t *testing.T) {
	vals, uncompressedSz := day9DataParser(exampleRaw)

	uncompacted := GetUncompacted(vals, uncompressedSz)
	expected := 42

	if len(uncompacted) != expected {
		t.Logf("uncompacted: %v", uncompacted)
		t.Fatalf("mismatched space length expected %d, actual:%d", expected, len(uncompacted))
	}
}

func TestDay9Part1(t *testing.T) {
	expected := 1928
	sln := SolveDay9Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay9Part2(t *testing.T) {
	expected := -1
	sln := SolveDay9Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
