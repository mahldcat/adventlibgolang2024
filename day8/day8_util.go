package day8

type Point struct {
	X int
	Y int
}

func (p Point) DistanceTo(p1 Point) Point {
	return Point{X: p1.X - p.X, Y: p1.Y - p.Y}
}

func (p1 Point) GetExtensionPoints(p2 Point, gSz Point) []Point {
	d := p1.DistanceTo(p2)
	inLine := make([]Point, 0)

	p3 := p1.Subtract(d)
	if p3.InBounds(gSz) {
		inLine = append(inLine, p3)
	}

	p4 := p2.Add(d)
	if p4.InBounds(gSz) {
		inLine = append(inLine, p4)
	}

	return inLine
}

func (p Point) GetExtensionPointsInLine(p1 Point, gSz Point) []Point {

	d := p.DistanceTo(p1)

	inLine := make([]Point, 0)

	for fromP := p.Subtract(d); fromP.InBounds(gSz); fromP = fromP.Subtract(d) {
		inLine = append(inLine, fromP)
	}

	for fromP := p.Add(d); fromP.InBounds(gSz); fromP = fromP.Add(d) {
		inLine = append(inLine, fromP)
	}

	return inLine
}

func (p Point) Add(p1 Point) Point {
	return Point{X: p.X + p1.X, Y: p.Y + p1.Y}
}

func (p Point) Subtract(p1 Point) Point {
	return Point{X: p.X - p1.X, Y: p.Y - p1.Y}
}

func (p Point) Equals(p1 Point) bool {
	return p.X == p1.X && p.Y == p1.Y
}

func (p Point) InBounds(gSz Point) bool {
	return p.X >= 0 && p.X < gSz.X && p.Y >= 0 && p.Y < gSz.Y
}

/*
   P1
   P2
   P3
   P4


*/
