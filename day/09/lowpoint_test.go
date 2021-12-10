package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = `2199943210
3987894921
9856789892
8767896789
9899965678`

//go:embed 09.txt
var puzzleInput string

func Test_parseInput(t *testing.T) {
	g := parseInput(exampleInput)
	require.Equal(t, 50, len(g.points))
	require.Equal(t, 10, g.w)
	require.Equal(t, 5, g.h)
	require.Equal(t, 2, g.points[0])
	require.Equal(t, 8, g.points[49])
}

func TestGrid_lowPoints(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		g := parseInput(exampleInput)
		lps := g.lowPoints()
		r := risk(lps)
		require.Equal(t, 15, r)
	})

	t.Run("puzzle A", func(t *testing.T) {
		g := parseInput(puzzleInput)
		lps := g.lowPoints()
		r := risk(lps)
		require.Equal(t, 580, r)
	})
}

func TestTopThree_insert(t *testing.T) {
	tt := topThree{}
	tt.insert(1)
	tt.insert(4)
	tt.insert(3)
	tt.insert(2)
	require.Equal(t, tt.a, 4)
	require.Equal(t, tt.b, 3)
	require.Equal(t, tt.c, 2)
}

func TestGrid_basinSize(t *testing.T) {
	g := parseInput(exampleInput)
	bs := g.basinSize(3, 2)
	require.Equal(t, 14, bs)
}

func TestGrid_largestBasins(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		g := parseInput(exampleInput)
		tt := g.largestBasins()
		require.Equal(t, 14, tt.a)
		require.Equal(t, 9, tt.b)
		require.Equal(t, 9, tt.b)
	})

	t.Run("puzzle B", func(t *testing.T) {
		g := parseInput(puzzleInput)
		tt := g.largestBasins()
		product := tt.a * tt.b * tt.c
		require.Equal(t, 856716, product)
	})
}
