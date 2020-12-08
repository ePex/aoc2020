package main

import (
	utils "aoc2020/util"
	"fmt"
)

func main() {
	values := utils.FileInt("day01/input.txt")

	fmt.Printf("Solution part one: %d\n", CalculateExpenseOfTwoEntries(values))
	fmt.Printf("Solution part two: %d\n", CalculateExpenseOfThreeEntries(values))
}
