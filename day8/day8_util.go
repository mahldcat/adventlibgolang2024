package day8

type Point struct {
	X int
	Y int
}

func (p Point) DistanceTo(p1 Point) Point {
	return Point{X: p1.X - p.X, Y: p1.Y - p.Y}
}

func (p Point) GetExtensionPoints(p1 Point) (Point, Point) {

	d := p.DistanceTo(p1)

	fromP := Point{X: p.X - d.X, Y: p.Y - d.Y}
	fromP1 := Point{X: p1.X + d.X, Y: p1.Y + d.Y}

	return fromP, fromP1
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
