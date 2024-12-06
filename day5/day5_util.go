package day5

func getMidPointValue(collectionList []int) int {
	midpoint := len(collectionList) / 2
	return collectionList[midpoint]
}

func contains(haystack []int, needle int) bool {
	for _, h := range haystack {
		if h == needle {
			return true
		}
	}
	return false
}

func verifyList(rules map[int][]int, orderList []int) bool {

	for i, pageToCheck := range orderList {
		if rules[pageToCheck] == nil { //freebie
			continue
		}

		pageRules := rules[pageToCheck]
		//forward scan on the rules
		for j := i + 1; j < len(orderList); j++ {
			againstPage := orderList[j]

			if !contains(pageRules, againstPage) {
				return false
			}
		}
		//now scan back
		// catches values that show up with no rules
		for j := i - 1; j >= 0; j-- {
			againstPage := orderList[j]

			if contains(pageRules, againstPage) {
				return false
			}
		}

	}

	return true
}

func repairUpdateOrder(rules map[int][]int, orderList []int) {

	//replicate the order list
	replica := make([]int, len(orderList))
	copy(replica, orderList)

	for idx := 0; idx < len(replica); idx++ {
		toAdjust := replica[idx]
		slotCt := 0

		for _, entry := range replica {
			if entry == toAdjust {
				continue //this is the value
			}

			//TIL that the maps also pass back a "key exists" param
			//if rules[entry] == nil {
			if _, exists := rules[entry]; !exists {
				continue //nothing on the entry we are checking
			}

			if contains(rules[entry], toAdjust) {
				slotCt++
			}

		}

		orderList[slotCt] = toAdjust
	}
}
