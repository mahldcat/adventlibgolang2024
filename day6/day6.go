package day6

func IsFacingObstruction(grid [][]rune, guard *Vector) bool {
	switch guard.Facing {
	case 'N':
		if guard.Y == 0 {
			return false //facing north at the top of the grid
		}
		return grid[guard.Y-1][guard.X] == '#'

	case 'E':
		if guard.X == (len(grid[0]) - 1) {
			return false //facing east on the right side of the grid
		}
		return grid[guard.Y][guard.X+1] == '#'

	case 'S':
		if guard.Y == (len(grid) - 1) {
			return false //facing south bottom of grid
		}
		return grid[guard.Y+1][guard.X] == '#'

	case 'W':
		if guard.X == 0 {
			return false //facing west on the left side of the grid
		}
		return grid[guard.Y][guard.X-1] == '#'
	}
	//should never reach here!
	return false
}

//return is escape,facingChange
func UpdateGuardStep(grid [][]rune, guard *Vector) (bool, bool) {

	facingChange := IsFacingObstruction(grid, guard)
	//fmt.Printf("Guard UPDATE Facing:(%c,%d,%d) changeType:%t \n", guard.Facing, guard.X, guard.Y, facingChange)

	if facingChange {
		guard.UpdateFacing()
	} else {
		guard.Move()
	}

	return guard.X < 0 || guard.Y < 0 || guard.X == len(grid[0]) || guard.Y == len(grid), facingChange
}

func SolveDay6Part1(rawData string) int {
	solnResult := 0

	grid, guard, _ := day6DataParser(rawData)

	vPath := New()

	for escape, facingChange := false, false; !escape; escape, facingChange = UpdateGuardStep(grid, &guard) {
		if !facingChange {

			if !vPath.LocationTraversed(guard) {
				solnResult++
				vPath.Append(guard)

			}
		}
	}
	return solnResult
}

func GetInitialPath(grid [][]rune, guard Vector) *VectorPath {
	vPath := New()

	//we know from soln one that vPath will be vectors that won't cycle
	//so we need to collect each step of the guard
	for escape, facingChange := false, false; !escape; escape, facingChange = UpdateGuardStep(grid, &guard) {
		if !facingChange {

			vPath.Append(guard)
		}
	}

	return vPath
}

func SolveDay6Part2(rawData string) int {
	cycleCoords := New()

	grid, guard, _ := day6DataParser(rawData)

	startX := guard.X
	startY := guard.Y

	guardSteps := GetInitialPath(grid, guard).GetConsumedCoordinates()

	for _, step := range guardSteps {
		grid[step.Y][step.X] = '#' //place a block
		guard.Reset(startX, startY)

		cycleDetected := false
		var testPath = New()

		for escape := false; !escape; escape, _ = UpdateGuardStep(grid, &guard) {

			if testPath.VectorCausesCycle(guard) {
				cycleDetected = true
				break
			}
			testPath.Append(guard)
		}

		if cycleDetected && !cycleCoords.LocationTraversed(step) {
			cycleCoords.Append(step)
		}

		grid[step.Y][step.X] = '.' //reset back to blankW
	}

	return cycleCoords.Length()
}
