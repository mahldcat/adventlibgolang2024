package day11

import (
	"fmt"
	"strconv"
)

/*
If the stone is engraved with the number 0, it is replaced by a stone engraved with the number 1.
If the stone is engraved with a number that has an even number of digits,
it is replaced by two stones.
The left half of the digits are engraved on the new left stone,
the right half of the digits are engraved on the new right stone.
(The new numbers don't keep extra leading zeroes: 1000 would become stones 10 and 0.)

If none of the other rules apply, the stone is replaced by a new stone;
the old stone's number multiplied by 2024 is engraved on the new stone.
*/
type StoneMap map[string]int

func rule1(stone int) ([]int, bool) {
	stoneSet := make([]int, 0)
	updated := false

	if stone == 0 {
		stoneSet = append(stoneSet, 1)
		updated = true
	}
	return stoneSet, updated
}

func rule2(stone int) ([]int, bool) {
	stoneSet := make([]int, 0)
	updated := false

	stoneStr := fmt.Sprintf("%d", stone)

	if len(stoneStr)%2 == 0 {

		left, _ := strconv.Atoi(stoneStr[:len(stoneStr)/2])
		right, _ := strconv.Atoi(stoneStr[len(stoneStr)/2:])
		//fmt.Printf("Splitting %d into [%d] %d\n", stone, left, right)
		stoneSet = append(stoneSet, left, right)

		updated = true
	}
	return stoneSet, updated
}

func rule3(stone int) ([]int, bool) {
	stoneSet := make([]int, 0)

	return append(stoneSet, stone*2024), true
}

func updateStone(stone int) []int {
	newStones, updated := rule1(stone)

	if !updated {
		newStones, updated = rule2(stone)

		if !updated {
			newStones, _ = rule3(stone)
		}
	}

	return newStones
}

func GetStoneCountDaySln(stone int, finalIteration int, prevCompute *StoneMap) int {
	return getStoneCount(stone, 1, finalIteration, stone, prevCompute)
}

// map[num][iteration] sum //this is prev
//
// iteration will start at 1
func getStoneCount(stone int, iteration int, finalIteration int, targetValue int, prevCompute *StoneMap) int {
	pc := *prevCompute

	memoKey := fmt.Sprintf("%d-%d", stone, iteration)
	value, ok := pc[memoKey]

	if ok {
		return value
	}

	update := updateStone(stone)

	if iteration == finalIteration { //this iteration is on the final iteration
		pc[memoKey] = len(update) //update the map so that we know that a stone on the final iteration should now have the given count
		return len(update)
	}

	sum := 0
	for _, nuStone := range update {
		sum += getStoneCount(nuStone, (iteration + 1), finalIteration, targetValue, prevCompute)
	}

	pc[memoKey] = sum

	return sum
}

//naive :D
func UpdateStoneSet(stoneSet []int, iterations int) []int {
	fmt.Printf("Enter UpdateStoneCt Iter[%d]:=SetCt(%d)\n", iterations, len(stoneSet))

	if iterations == 0 {
		return stoneSet
	}

	iterations--
	updatedSet := make([]int, 0)
	for _, stone := range stoneSet {
		updatedSet = append(updatedSet, updateStone(stone)...)
	}

	return UpdateStoneSet(updatedSet, iterations)
}
