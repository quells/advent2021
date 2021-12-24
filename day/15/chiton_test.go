package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleInput = `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`

//go:embed 15.txt
var puzzleInput string

func Test_parse(t *testing.T) {
	c := parse(exampleInput)
	require.Equal(t, 10, c.w)
	require.Equal(t, 10, c.h)
	require.Equal(t, 6, c.risks[2])
	require.Equal(t, 8, c.risks[98])
}

func Test_safestRoute(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		c := parse(exampleInput)
		risk := safestRoute(c)
		require.Equal(t, 40, risk)
	})

	t.Run("puzzle A", func(t *testing.T) {
		c := parse(puzzleInput)
		risk := safestRoute(c)
		require.Equal(t, 609, risk)
	})

	t.Run("example scaled", func(t *testing.T) {
		c := parse(exampleInput).scale(5)
		risk := safestRoute(c)
		require.Equal(t, 315, risk)
	})

	t.Run("puzzle B", func(t *testing.T) {
		c := parse(puzzleInput).scale(5)
		risk := safestRoute(c)
		require.Equal(t, 2925, risk)
	})
}

func TestCave_scale(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		c := parse("8")
		larger := c.scale(5)
		want := []int{
			8, 9, 1, 2, 3,
			9, 1, 2, 3, 4,
			1, 2, 3, 4, 5,
			2, 3, 4, 5, 6,
			3, 4, 5, 6, 7,
		}
		require.Equal(t, want, larger.risks)
	})
}
