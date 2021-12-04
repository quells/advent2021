package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleInput = `00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010`

//go:embed 03.txt
var puzzleInput string

func Test_decodeRates(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		codes := strings.Split(exampleInput, "\n")
		gamma, epsilon := decodeRates(codes)
		require.Equal(t, 22, gamma)
		require.Equal(t, 9, epsilon)
	})

	t.Run("puzzle A", func(t *testing.T) {
		codes := strings.Split(puzzleInput, "\n")
		gamma, epsilon := decodeRates(codes)
		require.Equal(t, 1337, gamma)
		require.Equal(t, 2758, epsilon)
		// 1337 * 2758 = 3687446
	})
}

func Test_decodeGases(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		codes := strings.Split(exampleInput, "\n")
		ox, co2 := decodeGases(codes)
		require.Equal(t, 23, ox)
		require.Equal(t, 10, co2)
	})

	t.Run("puzzle B", func(t *testing.T) {
		codes := strings.Split(puzzleInput, "\n")
		ox, co2 := decodeGases(codes)
		require.Equal(t, 1599, ox)
		require.Equal(t, 2756, co2)
		// 1599 * 2756 = 4406844
	})
}
