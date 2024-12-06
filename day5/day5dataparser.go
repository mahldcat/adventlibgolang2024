package day5

import (
	"bufio"
	"strconv"
	"strings"
)

func day5DataParser(rawData string) (map[int][]int, [][]int, error) {

	reader := strings.NewReader(rawData)
	scanner := bufio.NewScanner(reader)

	parseRules := true
	rulesMap := make(map[int][]int)

	updateList := make([][]int, 0)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" { //empty line means to switch to parsing updates
			parseRules = false
			continue
		}

		if parseRules {
			ruleRaw := strings.Split(line, "|")

			before, _ := strconv.Atoi(ruleRaw[0])
			after, _ := strconv.Atoi(ruleRaw[1])

			rulesMap[before] = append(rulesMap[before], after)

		} else {
			updateRaw := strings.Split(line, ",")
			update := make([]int, 0)
			for _, numRaw := range updateRaw {
				num, _ := strconv.Atoi(numRaw)
				update = append(update, num)
			}
			updateList = append(updateList, update)

		}
	}

	return rulesMap, updateList, scanner.Err()

}
