package day7

import "fmt"

func generalSolve(rawData string, includeConcat bool) int {
	combinedSolution := 0

	parsedList := day7DataParser(rawData)

	for _, entry := range parsedList {
		//this isn't needed for the AOC data, but nice to have
		//this is also something I'd be asking questionss on during an interview
		if len(entry) < 3 {
			fmt.Printf("Unexpected equation list size: %d", len(entry))
			continue
		}

		/*
		   original code was this if block that was....nasty?

		       if HasSolution(equationEntry[0], equationEntry[1]+equationEntry[2], equationEntry[3:], includeConcat) ||
		           HasSolution(equationEntry[0], equationEntry[1]*equationEntry[2], equationEntry[3:], includeConcat) ||
		           (includeConcat && HasSolution(equationEntry[0], concat(equationEntry[1], equationEntry[2]), equationEntry[3:], includeConcat)) {
		           solnResult += equationEntry[0]
		       }
		*/
		//the following is slightly more readable....
		solTtl := entry[0]
		e1 := entry[1]
		e2 := entry[2]
		remainder := entry[3:]

		if HasSolution(solTtl, e1+e2, remainder, includeConcat) {
			combinedSolution += solTtl
		} else if HasSolution(solTtl, e1*e2, remainder, includeConcat) {
			combinedSolution += solTtl
		} else if includeConcat && HasSolution(solTtl, concat(e1, e2), remainder, includeConcat) {
			combinedSolution += solTtl
		}

	}
	return combinedSolution
}

func SolveDay7Part1(rawData string) int {
	return generalSolve(rawData, false)
}

func SolveDay7Part2(rawData string) int {
	return generalSolve(rawData, true)
}
