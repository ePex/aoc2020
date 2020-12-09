package main

import (
	utils "aoc2020/util"
	"fmt"
)

func main() {
	values := utils.FileString("day06/input.txt")

	fmt.Printf("Solution part one: %d\n", CalculateYesCount(values, false))
	fmt.Printf("Solution part two: %d\n", CalculateYesCount(values, true))
}
