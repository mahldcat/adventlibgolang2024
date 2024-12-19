package day10

import (
	"testing"
)

var exampleRaw = `89010123
78121874
87430965
96549874
45678903
32019012
01329801
10456732`

func TestTrailHeadCount(t *testing.T) {
	topoMap := day10DataParser(exampleRaw)

	expected := 9

	th, _ := topoMap.GetTrailEndpoints()

	sln := len(th)

	if sln != expected {
		t.Errorf("mismatched trail head count expected %d, actual %d", expected, sln)
	}

}

func TestDay10Part1(t *testing.T) {
	expected := 36
	sln := SolveDay10Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay10Part2(t *testing.T) {
	expected := 81
	sln := SolveDay10Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
