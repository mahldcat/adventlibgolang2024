package day6

type Vector struct {
	X      int
	Y      int
	Facing rune
}

func (v *Vector) Reset(x int, y int) {
	v.X = x
	v.Y = y
	v.Facing = 'N'
}

func (v *Vector) UpdateFacing() {
	switch v.Facing {
	case 'N':
		v.Facing = 'E'
	case 'E':
		v.Facing = 'S'
	case 'S':
		v.Facing = 'W'
	case 'W':
		v.Facing = 'N'
	}
}

func (v *Vector) Move() {
	switch v.Facing {
	case 'N':
		v.Y--
	case 'E':
		v.X++
	case 'S':
		v.Y++
	case 'W':
		v.X--
	}
}

/**********************************************************/

type VectorPath struct {
	path []Vector
}

func New() *VectorPath {
	return &VectorPath{
		path: make([]Vector, 0), // Initializes an empty slice of pointers to Vector
	}
}

func (vp *VectorPath) Length() int {
	return len(vp.path)
}

func (vp *VectorPath) Append(v Vector) {
	vp.path = append(vp.path, v)
}

func (vp *VectorPath) GetConsumedCoordinates() []Vector {
	copyPaths := make([]Vector, len(vp.path))

	copy(copyPaths, vp.path)

	return copyPaths
}

func (vp *VectorPath) VectorCausesCycle(v Vector) bool {

	for _, pElement := range vp.path {

		if pElement.Facing == v.Facing && pElement.X == v.X && pElement.Y == v.Y {
			return true
		}

	}
	return false
}

func (vp *VectorPath) LocationTraversed(v Vector) bool {

	for _, pElement := range vp.path {

		if pElement.X == v.X && pElement.Y == v.Y {
			return true
		}

	}
	return false
}
