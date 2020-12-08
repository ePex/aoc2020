package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var world = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestCalculateEncounteredTrees(t *testing.T) {
	actual := CalculateEncounteredTrees(world, Point{3, 1})

	assert.Equal(t, 7, actual)
}

func TestCalculateEncounteredTreesMultipleSlopes(t *testing.T) {
	actual := CalculateEncounteredTreesMultipleSlopes(world, []Point{
		{1, 1},
		{3, 1},
		{5, 1},
		{7, 1},
		{1, 2},
	})

	assert.Equal(t, 336, actual)
}
