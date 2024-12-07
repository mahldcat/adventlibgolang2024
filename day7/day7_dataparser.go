package day7

import (
	"bufio"
	"strconv"
	"strings"
)

/*
190: 10 19
3267: 81 40 27
83: 17 5
156: 15 6
7290: 6 8 6 15
161011: 16 10 13
192: 17 8 14
21037: 9 7 18 13
292: 11 6 16 20
*/
func day7DataParser(rawData string) [][]int {
	reader := strings.NewReader(rawData)
	scanner := bufio.NewScanner(reader)

	parsedList := make([][]int, 0)

	/*
		couple of things regarding the parsing...Advent of Code data vs data I'd see in real life?

		First:  I'm playing fast and loose with the data--normally I would also pass back an Error entity with this...
		which I'd be appending any errors from Atoi calls, as well as the bufio scanner....

		Second:  I was able to see the data before hand and know I am always guaranteed at least 3 numbers per line
		[total and at least two values to combine]

		If this were a system without that data, I'd also want to clarify how to deal with
		count of numbers is less than 3...put them on, reject the line?
	*/

	for scanner.Scan() {
		line := scanner.Text()

		firstSplit := strings.Split(line, ":")
		total, _ := strconv.Atoi(firstSplit[0])
		//fmt.Printf("total:=%d: ", total)

		entry := []int{total}

		for _, rawNum := range strings.Split(firstSplit[1], " ") {

			if rawNum == "" {
				continue
			}
			num, _ := strconv.Atoi(rawNum)
			//fmt.Printf(" %d", num)
			entry = append(entry, num)
		}
		//fmt.Println("")

		parsedList = append(parsedList, entry)
	}

	return parsedList
}
