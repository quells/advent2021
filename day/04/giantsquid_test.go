package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

//go:embed 04.txt
var puzzleInput string

func Test_initBingoState(t *testing.T) {
	order, boards := initBingoState(exampleInput)
	wantOrder := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}
	require.Equal(t, wantOrder, order)
	require.Len(t, boards, 3)
	require.Len(t, boards[0].rows, 5)
	require.Len(t, boards[0].rows[0], 5)
}

func Test_playBingo(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		lastCalled, winningBoard := playBingo(initBingoState(exampleInput))
		require.Equal(t, 24, lastCalled)
		require.Equal(t, 188, winningBoard.sum())
		// 24 * 188 = 4512
	})

	t.Run("puzzle A", func(t *testing.T) {
		lastCalled, winningBoard := playBingo(initBingoState(puzzleInput))
		require.Equal(t, 57, lastCalled)
		require.Equal(t, 439, winningBoard.sum())
		// 57 * 439 = 25023
	})
}

func Test_loseBingo(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		lastCalled, losingBoard := playBingo(loseBingo(initBingoState(exampleInput)))
		require.Equal(t, 13, lastCalled)
		require.Equal(t, 148, losingBoard.sum())
		// 13 * 148 = 1924
	})

	t.Run("puzzle B", func(t *testing.T) {
		lastCalled, losingBoard := playBingo(loseBingo(initBingoState(puzzleInput)))
		require.Equal(t, 6, lastCalled)
		require.Equal(t, 439, losingBoard.sum())
		// 6 * 439 = 2634
	})
}
