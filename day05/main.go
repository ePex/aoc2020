package main

import (
	utils "aoc2020/util"
	"fmt"
)

func main() {
	values := utils.FileString("day05/input.txt")

	fmt.Printf("Solution part one: %d\n", GetHighestSeatId(values))
	fmt.Printf("Solution part two: %d\n", -1)
}
