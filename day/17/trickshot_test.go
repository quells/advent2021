package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

// target area: x=124..174, y=-123..-86

/*
Puzzle A:
When the probe starts with a positive initial Y-velocity,
it will pass through Y=0 again on its way down.
The highest altitude is achieved when the probe moves from
Y=0 to the bottom edge of the target area in a single step,
since this corresponds to the highest possible speed.
This means that the initial velocity is one less than this
distance. The maximum altitude is the sum of positive integers,
n*(n+1)/2. 122*123/2 = 7503
*/

var exampleTarget = target{
	x: [2]int{20, 30},
	y: [2]int{-10, -5},
}

var puzzleTarget = target{
	x: [2]int{124, 174},
	y: [2]int{-123, -86},
}

func Test_fire(t *testing.T) {
	tests := []struct {
		vx, vy  int
		wantHit bool
	}{
		{7, 2, true},
		{6, 3, true},
		{9, 0, true},
		{17, -4, false},
		{6, 9, true},
	}
	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d,%d", tt.vx, tt.vy), func(t *testing.T) {
			gotHit := fire(tt.vx, tt.vy, exampleTarget)
			require.Equal(t, tt.wantHit, gotHit)
		})
	}
}

func Test_countTrajectories(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := countValidTrajectories(exampleTarget)
		require.Equal(t, 112, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		got := countValidTrajectories(puzzleTarget)
		require.Equal(t, 3229, got)
	})
}
