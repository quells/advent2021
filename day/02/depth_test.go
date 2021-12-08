package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	exampleInput = strings.Split(`forward 5
down 5
forward 8
up 3
down 8
forward 2`, "\n")

	//go:embed 02.txt
	puzzleInput string
)

func TestParseCmd(t *testing.T) {
	tests := []struct {
		s      string
		dx, dd int
	}{
		{"forward 5", 5, 0},
		{"down 1", 0, 1},
		{"up 3", 0, -3},
	}
	for _, tt := range tests {
		t.Run(tt.s, func(t *testing.T) {
			dx, dd := parseCmd(tt.s)
			require.Equal(t, tt.dx, dx)
			require.Equal(t, tt.dd, dd)
		})
	}
}

func TestFollowCmds(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		var p pos
		p.followCmds(exampleInput)
		require.Equal(t, 15, p.x)
		require.Equal(t, 10, p.depth)
	})

	t.Run("puzzle A", func(t *testing.T) {
		var p pos
		p.followCmds(strings.Split(puzzleInput, "\n"))
		require.Equal(t, 1962, p.x)
		require.Equal(t, 987, p.depth)
		// 1962 * 987 = 1936494
	})
}

func TestFollowAimCmds(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		var p pos
		p.followAimCmds(exampleInput)
		require.Equal(t, 15, p.x)
		require.Equal(t, 60, p.depth)
	})

	t.Run("puzzle B", func(t *testing.T) {
		var p pos
		p.followAimCmds(strings.Split(puzzleInput, "\n"))
		require.Equal(t, 1962, p.x)
		require.Equal(t, 1017893, p.depth)
		// 1962 * 1017893 = 1997106066
	})
}
