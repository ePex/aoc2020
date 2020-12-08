package main

import (
	"strconv"
	"strings"
)

func CalculateNumberOfValidPasswordsSledPolicy(passwordData []string) (count int) {
	for _, rulePasswordPair := range passwordData {
		min, max, char, password := extractPasswordAndRule(rulePasswordPair)

		charCount := strings.Count(password, char)
		if charCount >= min && charCount <= max {
			count++
		}
	}

	return count
}

func CalculateNumberOfValidPasswordsTobogganPolicy(passwordData []string) (count int) {
	for _, rulePasswordPair := range passwordData {
		min, max, char, password := extractPasswordAndRule(rulePasswordPair)

		firstChar := string(password[min-1])
		secondChar := string(password[max-1])

		if neededCharAppearsOnce(char, firstChar, secondChar) {
			count++
		}
	}

	return count
}

func neededCharAppearsOnce(char string, chars ...string) bool {
	return strings.Count(strings.Join(chars, ""), char) == 1
}

func extractPasswordAndRule(policyWithPassword string) (min int, max int, char string, password string) {
	rule := strings.Split(policyWithPassword, ":")[0]
	password = strings.TrimSpace(strings.Split(policyWithPassword, ":")[1])

	expectedCount := strings.Split(rule, " ")[0]
	min, _ = strconv.Atoi(strings.Split(expectedCount, "-")[0])
	max, _ = strconv.Atoi(strings.Split(expectedCount, "-")[1])

	char = strings.Split(rule, " ")[1]

	return min, max, char, password
}
