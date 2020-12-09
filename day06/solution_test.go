package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWithOneGroup(t *testing.T) {
	assert.Equal(t, 6, CalculateYesCount([]string{
		"abcx",
		"abcy",
		"abcz",
	}))
}

func TestWithMultipleGroups(t *testing.T) {
	assert.Equal(t, 11, CalculateYesCount([]string{
		"abc",
		"",
		"a",
		"b",
		"c",
		"",
		"ab",
		"ac",
		"",
		"a",
		"a",
		"a",
		"a",
		"",
		"b",
	}))
}

func TestWithMultipleGroups2(t *testing.T) {
	assert.Equal(t, 23, CalculateYesCount([]string{
		"c",
		"omp",
		"yk",
		"w",
		"",
		"u",
		"fq",
		"",
		"qkvm",
		"qvlmk",
		"vmq",
		"bvmrq",
		"ifmqxagvu",
	}))
}
