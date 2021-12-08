package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = "16,1,2,0,4,2,7,1,2,14"

//go:embed 07.txt
var puzzleInput string

func Test_median(t *testing.T) {
	t.Run("even", func(t *testing.T) {
		got := median(parseInput("1,2,3,4,5,6"))
		want := 3
		require.Equal(t, want, got)
	})

	t.Run("odd", func(t *testing.T) {
		got := median(parseInput("1,2,3,4,5,6,7"))
		want := 4
		require.Equal(t, want, got)
	})

	t.Run("example", func(t *testing.T) {
		got := median(parseInput(exampleInput))
		want := 2
		require.Equal(t, want, got)
	})
}

func Test_linearFuel(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		positions := parseInput(exampleInput)
		target := median(positions)
		got := linearFuel(positions, target)
		want := 37
		require.Equal(t, want, got)
	})

	t.Run("puzzle A", func(t *testing.T) {
		positions := parseInput(puzzleInput)
		target := median(positions)
		got := linearFuel(positions, target)
		want := 331067
		require.Equal(t, want, got)

		// median is the minimum
		left := linearFuel(positions, target-1)
		right := linearFuel(positions, target+1)
		require.Less(t, got, left)
		require.Less(t, got, right)
	})
}

func Test_mean(t *testing.T) {
	t.Run("die", func(t *testing.T) {
		got := mean(parseInput("1,2,3,4,5,6"))
		want := 4 // 3.5
		require.Equal(t, want, got)
	})

	t.Run("example", func(t *testing.T) {
		got := mean(parseInput(exampleInput))
		want := 5
		require.Equal(t, want, got)
	})
}

func Test_sumFuel(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		positions := parseInput(exampleInput)
		target := mean(positions)
		got := sumFuel(positions, target)
		want := 168
		require.Equal(t, want, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		// mean is very close to the minimum

		positions := parseInput(puzzleInput)
		target := mean(positions)
		got := sumFuel(positions, target)
		for {
			if target > 0 {
				nextTarget := target - 1
				next := sumFuel(positions, nextTarget)
				if next < got {
					target = nextTarget
					got = next
				} else {
					break
				}
			}
		}

		want := 92881128
		require.Equal(t, want, got)
	})
}
