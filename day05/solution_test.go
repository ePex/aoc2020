package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetSeatId(t *testing.T) {
	assert.Equal(t, 357, GetSeatId("FBFBBFFRLR"))
	assert.Equal(t, 567, GetSeatId("BFFFBBFRRR"))
	assert.Equal(t, 119, GetSeatId("FFFBBBFRRR"))
	assert.Equal(t, 820, GetSeatId("BBFFBBFRLL"))
}
