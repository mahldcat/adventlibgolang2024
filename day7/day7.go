package day7

import "fmt"

func generalSolve(rawData string, includeConcat bool) int {
	solnResult := 0

	parsedList := day7DataParser(rawData)

	for _, equationEntry := range parsedList {
		//this isn't needed for the AOC data, but nice to have
		//this is also something I'd be asking questionss on during an interview
		if len(equationEntry) < 3 {
			fmt.Printf("Unexpected equation list size: %d", len(equationEntry))
			continue
		}

		if HasSolution(equationEntry[0], equationEntry[1]+equationEntry[2], equationEntry[3:], includeConcat) ||
			HasSolution(equationEntry[0], equationEntry[1]*equationEntry[2], equationEntry[3:], includeConcat) ||
			(includeConcat && HasSolution(equationEntry[0], concat(equationEntry[1], equationEntry[2]), equationEntry[3:], includeConcat)) {
			solnResult += equationEntry[0]
		}

	}
	return solnResult

}

func SolveDay7Part1(rawData string) int {
	return generalSolve(rawData, false)
}

func SolveDay7Part2(rawData string) int {
	return generalSolve(rawData, true)
}
