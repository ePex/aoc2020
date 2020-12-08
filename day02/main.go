package main

import (
	utils "aoc2020/util"
	"fmt"
)

func main() {
	values := utils.FileString("day02/input.txt", "\r\n")

	fmt.Printf("Solution part one: %d\n", CalculateNumberOfValidPasswordsSledPolicy(values))
	fmt.Printf("Solution part two: %d\n", CalculateNumberOfValidPasswordsTobogganPolicy(values))
}
