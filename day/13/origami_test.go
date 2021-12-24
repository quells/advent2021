package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleInput = `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`

//go:embed 13.txt
var puzzleInput string

func Test_parse(t *testing.T) {
	points, folds := parse(exampleInput)

	require.Len(t, points, 18)
	require.Equal(t, [2]int{6, 10}, points[0])
	require.Equal(t, [2]int{9, 0}, points[17])

	require.Len(t, folds, 2)
	require.Equal(t, fold{y: 7}, folds[0])
	require.Equal(t, fold{x: 5}, folds[1])
}

func Test_countVisible(t *testing.T) {
	tests := []struct {
		name   string
		points [][2]int
		want   int
	}{
		{"separate", [][2]int{{0, 0}, {1, 1}, {2, 2}}, 3},
		{"overlaps", [][2]int{{1, 1}, {1, 1}, {2, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := countVisible(tt.points)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_apply(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		points, folds := parse(exampleInput)
		require.Equal(t, 18, countVisible(points))

		points = apply(points, folds[0])
		require.Equal(t, 17, countVisible(points))

		points = apply(points, folds[1])
		require.Equal(t, 16, countVisible(points))
	})

	t.Run("puzzle A", func(t *testing.T) {
		points, folds := parse(puzzleInput)
		points = apply(points, folds[0])
		got := countVisible(points)
		require.Equal(t, 618, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		points, folds := parse(puzzleInput)
		for _, f := range folds {
			points = apply(points, f)
		}
		got := draw(points)
		want := `
.##..#....###..####.#..#.####.#..#.#..#
#..#.#....#..#.#....#.#..#....#.#..#..#
#..#.#....#..#.###..##...###..##...#..#
####.#....###..#....#.#..#....#.#..#..#
#..#.#....#.#..#....#.#..#....#.#..#..#
#..#.####.#..#.####.#..#.#....#..#..##.`
		require.Equal(t, want, got)
		// ALREKFKU
	})
}
