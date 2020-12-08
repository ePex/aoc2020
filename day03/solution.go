package main

type Point struct {
	x int
	y int
}

func CalculateEncounteredTreesMultipleSlopes(world []string, slopes []Point) (count int) {
	count = 1
	for _, slope := range slopes {
		count *= CalculateEncounteredTrees(world, slope)
	}

	return count
}

func CalculateEncounteredTrees(world []string, slope Point) (count int) {
	width, height := getWorldSize(world)

	currentPos := Point{0, 0}
	for currentPos.y < (height - slope.y) {
		moveX := calculateMove(currentPos.x, slope.x, width)
		moveY := currentPos.y + slope.y

		currentPos.x = moveX
		currentPos.y = moveY

		if treeAtPos(world, currentPos) {
			count++
		}
	}

	return count
}

func treeAtPos(world []string, pos Point) bool {
	return string(world[pos.y][pos.x]) == "#"
}

func getWorldSize(world []string) (width int, height int) {
	return len(world[0]), len(world)
}

func calculateMove(currentPos int, moveX int, maxPos int) int {
	return (currentPos + moveX) % maxPos
}
