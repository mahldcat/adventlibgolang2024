package day5

import (
	"testing"
)

var exampleRaw = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47`

func TestParser(t *testing.T) {

	rules, updates, err := day5DataParser(exampleRaw)
	expectedUpdatesCount := 6
	expectedRulesCount := 6
	expectedRuleListCount := 4

	if err != nil {
		t.Fatalf("Parsing Error occurred: %v", err)
	}

	if len(rules) != expectedRulesCount {
		t.Errorf("unexpected rules count expected %d, actual %d", expectedRulesCount, len(rules))
	}

	rulePage := 47
	if len(rules[rulePage]) != expectedRuleListCount {
		t.Errorf("unexpected rule list count expected %d, actual %d", expectedRuleListCount, len(rules[rulePage]))
	}

	if len(updates) != expectedUpdatesCount {
		t.Errorf("unexpected rules count expected %d, actual %d", expectedUpdatesCount, len(updates))
	}

}

func TestDay5Part1(t *testing.T) {

	expected := 143
	sln := SolveDay5Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}

}

func TestDay5Part2(t *testing.T) {

	expected := 123
	sln := SolveDay5Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}
