package day7

func HasSolution(total int, subtotal int, remainder []int) bool {
	if len(remainder) == 0 {
		return subtotal == total
	}

	if HasSolution(total, subtotal+remainder[0], remainder[1:]) {
		return true
	}
	return HasSolution(total, subtotal*remainder[0], remainder[1:])
}
