package main

import (
	"strconv"
	"strings"
)

func CalculateNumberOfValidPasswordsSledPolicy(passwordData []string) (count int) {
	result := 0

	for _, rulePasswordPair := range passwordData {
		min, max, char, password := extractPasswordAndRule(rulePasswordPair)

		count := strings.Count(password, char)
		if count >= min && count <= max {
			result++
		}
	}

	return result
}

func extractPasswordAndRule(policyWithPassword string) (min int, max int, char string, password string) {
	rule := strings.Split(policyWithPassword, ":")[0] // "1-3 a: abcde", "1-3 b: cdefg", "2-9 c: ccccccccc"

	expectedCount := strings.Split(rule, " ")[0]
	min, _ = strconv.Atoi(strings.Split(expectedCount, "-")[0])
	max, _ = strconv.Atoi(strings.Split(expectedCount, "-")[1])

	char = strings.Split(rule, " ")[1]
	password = strings.Split(policyWithPassword, ":")[1]

	return min, max, char, password
}
