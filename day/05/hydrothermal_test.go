package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`

//go:embed 05.txt
var puzzleInput string

func Test_parseLine(t *testing.T) {
	tests := []struct {
		input string
		want  line
	}{
		{"0,9 -> 5,9", line{{0, 9}, {5, 9}}},
		{"10,29 -> 35,49", line{{10, 29}, {35, 49}}},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			got := parseLine(tt.input)
			require.Equal(t, tt.want, got)
		})
	}
}

func TestLine_isOrthogonal(t *testing.T) {
	tests := []struct {
		name string
		l    line
		want bool
	}{
		{"horizontal", line{{0, 9}, {5, 9}}, true},
		{"vertical", line{{7, 0}, {7, 4}}, true},
		{"diagonal", line{{8, 0}, {0, 8}}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.isOrthogonal()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestLine_points(t *testing.T) {
	tests := []struct {
		name string
		l    line
		want [][2]int
	}{
		{"horizontal", line{{0, 9}, {3, 9}}, [][2]int{{0, 9}, {1, 9}, {2, 9}, {3, 9}}},
		{"vertical", line{{7, 3}, {7, 0}}, [][2]int{{7, 0}, {7, 1}, {7, 2}, {7, 3}}},
		{"diagonal ne", line{{1, 2}, {3, 4}}, [][2]int{{1, 2}, {2, 3}, {3, 4}}},
		{"diagonal se", line{{1, 4}, {3, 2}}, [][2]int{{1, 4}, {2, 3}, {3, 2}}},
		{"diagonal sw", line{{3, 4}, {1, 2}}, [][2]int{{3, 4}, {2, 3}, {1, 2}}},
		{"diagonal nw", line{{3, 2}, {1, 4}}, [][2]int{{3, 2}, {2, 3}, {1, 4}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.l.points()
			require.Equal(t, tt.want, got)
		})
	}
}

func TestVentMap_countOverlaps(t *testing.T) {
	tests := []struct {
		name         string
		input        string
		wantOrtho    int
		wantOverlaps int
	}{
		{"example", exampleInput, 5, 12},
		{"puzzle", puzzleInput, 5698, 15463},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			allLines := parseInput(tt.input)

			t.Run("ortho", func(t *testing.T) {
				var lines []line
				for _, l := range allLines {
					if l.isOrthogonal() {
						lines = append(lines, l)
					}
				}

				vm := make(ventMap)
				for _, l := range lines {
					vm.draw(l)
				}

				overlaps := vm.countOverlaps()
				require.Equal(t, tt.wantOrtho, overlaps)
			})

			t.Run("all", func(t *testing.T) {
				vm := make(ventMap)
				for _, l := range allLines {
					vm.draw(l)
				}

				overlaps := vm.countOverlaps()
				require.Equal(t, tt.wantOverlaps, overlaps)
			})
		})
	}
}
