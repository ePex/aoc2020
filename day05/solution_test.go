package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSeatId(t *testing.T) {
	assert.Equal(t, 357, GetHighestSeatId([]string{"FBFBBFFRLR"}))
	assert.Equal(t, 567, GetHighestSeatId([]string{"BFFFBBFRRR"}))
	assert.Equal(t, 119, GetHighestSeatId([]string{"FFFBBBFRRR"}))
	assert.Equal(t, 820, GetHighestSeatId([]string{"BBFFBBFRLL"}))
}

func TestFindYourSeatId(t *testing.T) {
	assert.Equal(t, 6, FindYourSeatId([]string{
		"FFFFFFFRLL", // 4
		"FFFFFFFRLR", // 5
		"FFFFFFFRRR", // 7
	}))
}
