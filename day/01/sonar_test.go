package main

import (
	_ "embed"
	"strconv"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed 01.txt
	puzzleText string

	exampleInput = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
)

func TestDeeperCount(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := deeperCount(exampleInput)
		want := 7
		require.Equal(t, want, got)
	})

	var puzzleInput []int
	for _, s := range strings.Split(puzzleText, "\n") {
		v, _ := strconv.Atoi(s)
		puzzleInput = append(puzzleInput, v)
	}

	t.Run("puzzle A", func(t *testing.T) {
		got := deeperCount(puzzleInput)
		want := 1791
		require.Equal(t, want, got)
	})

	t.Run("sliding example", func(t *testing.T) {
		got := deeperCount(slidingSums(exampleInput, 3))
		want := 5
		require.Equal(t, want, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		got := deeperCount(slidingSums(puzzleInput, 3))
		want := 1822
		require.Equal(t, want, got)
	})
}

func TestSlidingSums(t *testing.T) {
	got := slidingSums(exampleInput, 3)
	want := []int{607, 618, 618, 617, 647, 716, 769, 792}
	require.Equal(t, want, got)
}
