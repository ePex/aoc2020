package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidateBatch(t *testing.T) {
	assert.Equal(t, 567, GetSeatId("BFFFBBFRRR"))
	assert.Equal(t, 119, GetSeatId("FFFBBBFRRR"))
	assert.Equal(t, 820, GetSeatId("BBFFBBFRLL"))
}
