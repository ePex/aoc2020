package main

import (
	utils "aoc2020/util"
	"fmt"
)

func main() {
	values := utils.FileString("day03/input.txt")

	fmt.Printf("Solution part one: %d\n", CalculateEncounteredTrees(values, Point{3, 1}))
	fmt.Printf("Solution part two: %d\n", CalculateEncounteredTreesMultipleSlopes(values, []Point{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	}))
}
