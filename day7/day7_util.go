package day7

import (
	"fmt"
	"strconv"
)

func concat(a int, b int) int {

	//I could do this as this one liner....but readability suffers
	//ab,_ := strconv.Atoi( fmt.Sprintf("%s%s",  strconv.Itoa(a),strconv.Itoa(b)))
	//and then you realize....I don't need to convert int back to string (idjut)

	concatRaw := fmt.Sprintf("%d%d", a, b)
	parsed, _ := strconv.Atoi(concatRaw)

	return parsed
}

func HasSolution(total int, subtotal int, remainder []int, includeConCat bool) bool {
	if len(remainder) == 0 {
		return subtotal == total
	}

	if HasSolution(total, subtotal+remainder[0], remainder[1:], includeConCat) {
		return true
	}
	if HasSolution(total, subtotal*remainder[0], remainder[1:], includeConCat) {
		return true
	}

	if includeConCat {
		return HasSolution(total, concat(subtotal, remainder[0]), remainder[1:], includeConCat)
	}

	return false
}
