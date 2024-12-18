package day8

import (
	"testing"
)

var exampleRaw = `............
........0...
.....0......
.......0....
....0.......
......A.....
............
............
........A...
.........A..
............
............`

func VerifyToken(pMap map[rune][]Point, token rune, expected int, t *testing.T) {
	if pMap[token] == nil {
		t.Fatalf("results for  token %q do not exist", token)
	}
	if len(pMap[token]) != expected {
		t.Fatalf("parsed results for %q mismatch expected %d, actual %d", token, expected, len(pMap['A']))
	}
}

func TestParser(t *testing.T) {
	pMap, gSz := day8DataParser(exampleRaw)

	expSz := Point{X: 12, Y: 12}

	if !gSz.Equals(expSz) {
		t.Fatalf("unexpected Grid Size, epected (%d,%d), actual (%d,%d)", gSz.X, gSz.Y, expSz.X, expSz.Y)
	}

	VerifyToken(pMap, 'A', 3, t)
	VerifyToken(pMap, '0', 4, t)
}

func TestDay8Part1(t *testing.T) {
	expected := 14
	sln := SolveDay8Part1(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}

func TestDay8Part2(t *testing.T) {
	expected := 34
	sln := SolveDay8Part2(exampleRaw)

	if sln != expected {
		t.Errorf("mismatched solution value expected %d, actual %d", expected, sln)
	}
}

func TestDist(t *testing.T) {
	p := Point{X: 2, Y: 2}
	p2 := Point{X: 6, Y: 1}

	d := p.DistanceTo(p2)

	expectedX := 4
	expectedY := -1

	if d.X != expectedX {
		t.Fatalf("Invalid Distance x, expected:%d, actual:%d", expectedX, d.X)
	}

	if d.Y != expectedY {
		t.Fatalf("Invalid Distance y, expected:%d, actual:%d", expectedY, d.Y)
	}
}

func TestFromPoint(t *testing.T) {
	p := Point{X: 2, Y: 2}
	p2 := Point{X: 6, Y: 1}

	gSz := Point{X: 9999, Y: 9999}

	/*
		this one will be calculated but thrown away because its OOB  X <0)
		expectedP3 := Point{X: -2, Y: 3}
	*/
	expectedP4 := Point{X: 10, Y: 0}

	pts := p.GetExtensionPoints(p2, gSz)

	if len(pts) != 1 {
		t.Fatalf("GetExtensionPoints slice count failed, expected 2, actual %d", len(pts))
	}

	if !pts[0].Equals(expectedP4) {
		t.Fatalf("invalid point, expected (%d,%d), actual (%d,%d)", pts[1].X, pts[1].Y, expectedP4.X, expectedP4.Y)
	}
}
