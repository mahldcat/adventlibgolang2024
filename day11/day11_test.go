package day11

import (
	"testing"
)

var exampleRaw = "125 17"

func TestParser(t *testing.T) {

	stones := day11DataParser(exampleRaw)

	expectedStones := []int{125, 17}

	if len(stones) != len(expectedStones) {
		t.Errorf("mismatched stone set len expected %d actual %d", len(expectedStones), len(stones))

		for i, stone := range stones {
			if stone != expectedStones[i] {

				t.Errorf("mismatched stone[%d], expected %d actual %d", i, expectedStones[i], stone)
			}
		}
	}

}

func TestGetStoneCountSingleIter(t *testing.T) {

	stone := 0
	iterCt := 25
	maxIterCt := 25
	compMap := make(StoneMap)

	stoneCt := getStoneCount(stone, iterCt, maxIterCt, stone, &compMap)

	if stoneCt != 1 {
		t.Errorf("mismatched stoneCt expected %d actual %d", 1, stoneCt)

	}
}

func TestGetStoneCountZerowith25Iter(t *testing.T) {
	stone := 0
	maxIterCt := 25
	compMap := make(StoneMap)

	//this is just verifying we don't stack overflow!!!!
	_ = GetStoneCountDaySln(stone, maxIterCt, &compMap)

}

func TestSecondInputSingleIter(t *testing.T) {

	stoneRaw := "0 1 10 99 999"
	expectedCt := 7

	stones := day11DataParser(stoneRaw)
	computeMap := make(StoneMap)
	stoneCt := 0
	for _, stone := range stones {
		stoneCt += GetStoneCountDaySln(stone, 1, &computeMap)
	}

	if stoneCt != expectedCt {
		t.Fatalf("invalid stone count after one iteration, expected:%d actual:%d", expectedCt, stoneCt)
	}
}

func TestDay11Part1(t *testing.T) {
	expected := 55312
	sln := SolveDay11Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay11Part2(t *testing.T) {
	expected := -1
	sln := SolveDay11Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
