package day10

type Direction int

const (
	N Direction = iota
	S
	E
	W
)

var (
	PointN = Point{X: 0, Y: -1}
	PointS = Point{X: 0, Y: 1}
	PointE = Point{X: 1, Y: 0}
	PointW = Point{X: -1, Y: 0}
)

func (d *Direction) ToIndexVector() Point {

	if *d == N {
		return PointN
	}
	if *d == S {
		return PointS
	}
	if *d == E {
		return PointE
	}

	return PointW
}

type Point struct {
	X int
	Y int
}

func (p *Point) Move(vector Point) Point {
	return Point{
		X: p.X + vector.X,
		Y: p.Y + vector.Y,
	}
}

type MapCoordinate struct {
	Point
	Ht int
}

func (mc *MapCoordinate) IsTrailHead() bool {
	return mc.Ht == 0
}

func (mc *MapCoordinate) IsPeak() bool {
	return mc.Ht == 9
}

func (mc *MapCoordinate) Traverse(topoMap TopoMap, direction Direction) (MapCoordinate, bool) {

	newPoint := mc.Move(direction.ToIndexVector())

	oob := topoMap.OutOfBounds(newPoint)

	if oob {
		return topoMap.GetEmptyCoordinate(), false
	}

	newCoord := topoMap.MapData[newPoint.Y][newPoint.X]
	if newCoord.Ht-mc.Ht != 1 {
		return topoMap.GetEmptyCoordinate(), false
	}

	return newCoord, true
}

type TopoMap struct {
	MapData   [][]MapCoordinate
	Dimension Point
}

func (tm *TopoMap) GetEmptyCoordinate() MapCoordinate {
	return MapCoordinate{}
}

func (tm *TopoMap) GetTrailEndpoints() ([]MapCoordinate, []MapCoordinate) {

	trailHeads := make([]MapCoordinate, 0)
	trailPeaks := make([]MapCoordinate, 0)

	for _, xMap := range tm.MapData {
		for _, mapPoint := range xMap {
			if mapPoint.IsTrailHead() {
				trailHeads = append(trailHeads, mapPoint)
			}
			if mapPoint.IsPeak() {
				trailPeaks = append(trailPeaks, mapPoint)
			}
		}
	}
	return trailHeads, trailPeaks
}

func (tm *TopoMap) OutOfBounds(p Point) bool {
	return p.X >= tm.Dimension.X || p.X < 0 || p.Y >= tm.Dimension.Y || p.Y < 0
}

func (topoMap *TopoMap) getTrailCount(coordinate MapCoordinate, count *int, totalPaths *int, peekReached *map[MapCoordinate]bool) {
	if coordinate.IsPeak() {
		(*totalPaths)++
		if !(*peekReached)[coordinate] {
			(*peekReached)[coordinate] = true
			(*count)++
		}
		return
	}

	mapCoord, valid := coordinate.Traverse(*topoMap, N)
	if valid {
		topoMap.getTrailCount(mapCoord, count, totalPaths, peekReached)
	}
	mapCoord, valid = coordinate.Traverse(*topoMap, S)
	if valid {
		topoMap.getTrailCount(mapCoord, count, totalPaths, peekReached)
	}
	mapCoord, valid = coordinate.Traverse(*topoMap, E)
	if valid {
		topoMap.getTrailCount(mapCoord, count, totalPaths, peekReached)
	}
	mapCoord, valid = coordinate.Traverse(*topoMap, W)
	if valid {
		topoMap.getTrailCount(mapCoord, count, totalPaths, peekReached)
	}
}

func (topoMap *TopoMap) GetPathCounts() (int, int) {

	pathCount := 0
	totalPathCount := 0

	trailHeads, _ := topoMap.GetTrailEndpoints()

	for _, trailHead := range trailHeads {

		peekReached := make(map[MapCoordinate]bool)

		topoMap.getTrailCount(trailHead, &pathCount, &totalPathCount, &peekReached)
	}

	return pathCount, totalPathCount
}

func (topoMap *TopoMap) ToString() string {
	return ""
}
