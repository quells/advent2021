package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var (
	//go:embed 19samepos.txt
	samePosInput string

	//go:embed 19example.txt
	exampleInput string
)

func TestParse(t *testing.T) {
	readings := Parse(exampleInput)

	require.Len(t, readings, 5)
	require.Equal(t, 25, readings[0].n)
	require.Equal(t, 404, readings[0].x[0])
	require.Equal(t, 686, readings[1].x[0])
}

func TestReading_Crosses(t *testing.T) {
	readings := Parse(samePosInput)

	zero := readings[0]
	want := CountAngleMatches(zero.angles)
	for i := 0; i < len(readings); i++ {
		ci := readings[i]
		for j := 0; j < len(readings); j++ {
			cj := readings[j]
			require.Equal(t, want, CountAngleMatches(ci.Intersect(cj)))
		}
	}

	for i := 1; i < len(readings); i++ {
		readings[i].alignTo(zero)
		require.Equal(t, zero.x, readings[i].x)
		require.Equal(t, zero.y, readings[i].y)
		require.Equal(t, zero.z, readings[i].z)
	}
}

func TestCrossAngleFingerprinting(t *testing.T) {
	readings := Parse(exampleInput)

	zero := readings[0]
	require.Equal(t, CountAngleMatches(zero.Intersect(readings[1])), 660)
	require.Equal(t, CountAngleMatches(zero.Intersect(readings[2])), 3)
	require.Equal(t, CountAngleMatches(zero.Intersect(readings[3])), 0)
	require.Equal(t, CountAngleMatches(zero.Intersect(readings[4])), 60)
}

func Test_rotationMatrix(t *testing.T) {
	zero := Cross{
		c1: -7, c2: 15, c3: 8,
	}
	other := Cross{
		c1: -15, c2: 8, c3: 7,
	}

	got := findRotationMatrix(other, zero)
	want := [9]int{
		0, 0, -1,
		-1, 0, 0,
		0, 1, 0,
	}
	require.Equal(t, want, got)

	other.applyRotationMatrix(got)
	require.Equal(t, zero, other)
}

func TestAlignAndCombine(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		readings := Parse(exampleInput)
		combined := AlignAndCombine(readings)
		require.Equal(t, 79, combined.n)
	})

	// t.Run("puzzle A", func(t *testing.T) {
	// 	readings := Parse(puzzleInput)
	// 	combined := AlignAndCombine(readings)
	// 	require.Equal(t, 79, combined.n)
	// })
}
