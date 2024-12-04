package day4

import (
	"testing"
)

func getExampleData() []struct {
	name             string
	rawData          string
	expectedRuneGrid [][]rune
} {
	return []struct {
		name             string
		rawData          string
		expectedRuneGrid [][]rune
	}{
		{
			name: "ExampleTestCase",
			rawData: `
MMMSXXMASM
MSAMXMSMSA
AMXSXMAAMM
MSAMASMSMX
XMASAMXAMM
XXAMMXXAMA
SMSMSASXSS
SAXAMASAAA
MAMMMXMMMM
MXMXAXMASX`,
			expectedRuneGrid: [][]rune{
				{'M', 'M', 'M', 'S', 'X', 'X', 'M', 'A', 'S', 'M'},
				{'M', 'S', 'A', 'M', 'X', 'M', 'S', 'M', 'S', 'A'},
				{'A', 'M', 'X', 'S', 'X', 'M', 'A', 'A', 'M', 'M'},
				{'M', 'S', 'A', 'M', 'A', 'S', 'M', 'S', 'M', 'X'},
				{'X', 'M', 'A', 'S', 'A', 'M', 'X', 'A', 'M', 'M'},
				{'X', 'X', 'A', 'M', 'M', 'X', 'X', 'A', 'M', 'A'},
				{'S', 'M', 'S', 'M', 'S', 'A', 'S', 'X', 'S', 'S'},
				{'S', 'A', 'X', 'A', 'M', 'A', 'S', 'A', 'A', 'A'},
				{'M', 'A', 'M', 'M', 'M', 'X', 'M', 'M', 'M', 'M'},
				{'M', 'X', 'M', 'X', 'A', 'X', 'M', 'A', 'S', 'X'},
			},
		},
	}
}

func TestStringToRuneGrid(t *testing.T) {

	for _, tc := range getExampleData() {
		t.Run(tc.name, func(t *testing.T) {

			result := day4DataParser(tc.rawData)

			for y := 0; y < len(result); y++ {
				row := result[y]
				if len(row) != len(tc.expectedRuneGrid[y]) {
					t.Errorf("Error(y:=%d) on row length", y)
				}
				for x := 0; x < len(row); x++ {
					exp := tc.expectedRuneGrid[y][x]
					res := result[y][x]
					if res != exp {
						t.Errorf("Error(x:=%d,y:=%d)  Expected %c got %c", x, y, exp, res)
					}

				}
			}
		})
	}

}
