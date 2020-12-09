package main

import (
	utils "aoc2020/util"
	"fmt"
)

func main() {
	values := utils.FileString("day04/input.txt")

	fmt.Printf("Solution part one: %d\n", ValidatePassportBatch(values))
	fmt.Printf("Solution part two: %d\n", ValidatePassportBatchWithValidationRules(values))
}
