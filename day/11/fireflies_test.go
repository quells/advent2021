package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`

const puzzleInput = `5665114554
4882665427
6185582113
7762852744
7255621841
8842753123
8225372176
7212865827
7758751157
1828544563`

func TestGrid_step(t *testing.T) {
	t.Run("overload", func(t *testing.T) {
		const step0 = `11111
19991
19191
19991
11111`
		const step1 = `34543
40004
50005
40004
34543`

		g := parseInput(step0)
		gotBlinked := g.step()
		gotGrid := g.points

		wantBlinked := 9
		wantGrid := parseInput(step1).points

		require.Equal(t, wantBlinked, gotBlinked)
		require.Equal(t, wantGrid, gotGrid)
	})

	t.Run("example", func(t *testing.T) {
		g := parseInput(exampleInput)

		var gotBlinked int
		for i := 1; i <= 10; i++ {
			gotBlinked += g.step()
		}

		require.Equal(t, 204, gotBlinked)

		for i := 11; i <= 100; i++ {
			gotBlinked += g.step()
		}

		require.Equal(t, 1656, gotBlinked)
	})

	t.Run("puzzle A", func(t *testing.T) {
		g := parseInput(puzzleInput)

		var gotBlinked int
		for i := 1; i <= 100; i++ {
			gotBlinked += g.step()
		}

		require.Equal(t, 1617, gotBlinked)
	})
}

func BenchmarkGrid_step(b *testing.B) {
	g := parseInput(exampleInput)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		g.step()
	}
}

func TestGrid_isSynched(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		g := parseInput(exampleInput)
		var gotT int
		for {
			gotT++
			g.step()
			if g.isSynched() {
				break
			}
		}

		require.Equal(t, 195, gotT)
	})

	t.Run("puzzle B", func(t *testing.T) {
		g := parseInput(puzzleInput)
		var gotT int
		for {
			gotT++
			g.step()
			if g.isSynched() {
				break
			}
		}

		require.Equal(t, 258, gotT)
	})
}
