package day6

import (
	"testing"
)

var exampleRaw = `....#.....
.........#
..........
..#.......
.......#..
..........
.#..^.....
........#.
#.........
......#...`

func TestParser(t *testing.T) {
	expectedGuard := Vector{X: 4, Y: 5}
	expectedGridXSz := 10
	expectedGridYSz := 10

	grid, guard, err := day6DataParser(exampleRaw)
	if err != nil {
		t.Fatalf("Failure in parsing %v", err)
	}

	if guard.X != expectedGuard.X && guard.Y != expectedGuard.Y {
		t.Fatalf("unexpected guard location (%d,%d) expected: (%d,%d)", guard.X, guard.Y, expectedGuard.X, expectedGuard.Y)
	}
	if len(grid) != expectedGridYSz {
		t.Fatalf("unexpected row cound in grid %d, expected %d", len(grid), expectedGridYSz)
	}
	if len(grid[0]) != expectedGridXSz {
		t.Fatalf("unexpected row cound in grid %d, expected %d", len(grid[0]), expectedGridXSz)
	}
}

func TestDay6Part1(t *testing.T) {
	expected := 41
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
