package main

import (
	"strconv"
	"strings"
)

func CalculateNumberOfValidPasswordsSledPolicy(passwordData []string) (count int) {
	result := 0

	for _, rulePasswordPair := range passwordData {
		rule := strings.Split(rulePasswordPair, ":")[0] // "1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"

		expectedCount := strings.Split(rule, " ")[0]
		least, _ := strconv.Atoi(strings.Split(expectedCount, "-")[0])
		most, _ := strconv.Atoi(strings.Split(expectedCount, "-")[1])

		char := strings.Split(rule, " ")[1]
		password := strings.Split(rulePasswordPair, ":")[1]
		count := strings.Count(password, char)

		if count >= least && count <= most {
			result++
		}
	}

	return result
}
