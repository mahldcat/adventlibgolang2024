package day1

import (
	"testing"
)

func getExampleData(expected int) []struct {
	name      string
	left      []int
	right     []int
	expected  int
	expectErr bool
} {
	return []struct {
		name      string
		left      []int
		right     []int
		expected  int
		expectErr bool
	}{
		{
			name:      "Single test case",
			left:      []int{3, 4, 2, 1, 3, 3},
			right:     []int{4, 3, 5, 3, 9, 3},
			expected:  expected,
			expectErr: false,
		},
	}
}

func TestSolveDay1Part1ExampleData(t *testing.T) {

	expected := 11
	for _, tc := range getExampleData(expected) {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SolveDay1Part1(tc.left, tc.right)

			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}

}

func TestSolveDay1Part2ExampleData(t *testing.T) {

	expected := 31
	for _, tc := range getExampleData(expected) {
		t.Run(tc.name, func(t *testing.T) {
			result, err := SolveDay1Part2(tc.left, tc.right)

			if tc.expectErr && err == nil {
				t.Errorf("Expected an error but got nil")
			}
			if !tc.expectErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if result != tc.expected {
				t.Errorf("Expected %d, got %d", tc.expected, result)
			}
		})
	}

}
