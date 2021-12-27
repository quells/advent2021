package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed 18.txt
var puzzleInput string

func TestParse(t *testing.T) {
	input := "[[[[1,3],[5,3]],[[1,3],[8,7]]],[[[4,9],[6,9]],[[8,2],[7,3]]]]"
	n := ParseString(input)
	require.Equal(t, input, n.String())
}

func TestReduce(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			"explode leftmost",
			"[[[[[9,8],1],2],3],4]",
			"[[[[0,9],2],3],4]",
		},
		{
			"explode rightmost",
			"[7,[6,[5,[4,[3,2]]]]]",
			"[7,[6,[5,[7,0]]]]",
		},
		{
			"explode middle",
			"[[6,[5,[4,[3,2]]]],1]",
			"[[6,[5,[7,0]]],3]",
		},
		{
			"nested explode",
			"[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			"[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			"split even",
			"[10,0]",
			"[[5,5],0]",
		},
		{
			"split odd",
			"[0,11]",
			"[0,[5,6]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			n := ParseString(tt.input)
			r := Reduce(n)
			require.Equal(t, tt.expect, r.String())
		})
	}
}

func TestAdd(t *testing.T) {
	tests := []struct {
		name   string
		a, b   string
		expect string
	}{
		{
			"simple",
			"[1,2]", "[[3,4],5]",
			"[[1,2],[[3,4],5]]",
		},
		{
			"example",
			"[[[[4,3],4],4],[7,[[8,4],9]]]", "[1,1]",
			"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := ParseString(tt.a)
			b := ParseString(tt.b)
			s := Add(a, b)
			require.Equal(t, tt.expect, s.String())
		})
	}
}

func TestSum(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect string
	}{
		{
			"four",
			`[1,1]
[2,2]
[3,3]
[4,4]`,
			"[[[[1,1],[2,2]],[3,3]],[4,4]]",
		},
		{
			"five",
			`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]`,
			"[[[[3,0],[5,3]],[4,4]],[5,5]]",
		},
		{
			"six",
			`[1,1]
[2,2]
[3,3]
[4,4]
[5,5]
[6,6]`,
			"[[[[5,0],[7,4]],[5,5]],[6,6]]",
		},
		{
			"slightly larger",
			`[[[0,[4,5]],[0,0]],[[[4,5],[2,6]],[9,5]]]
[7,[[[3,7],[4,3]],[[6,3],[8,8]]]]
[[2,[[0,8],[3,4]]],[[[6,7],1],[7,[1,6]]]]
[[[[2,4],7],[6,[0,5]]],[[[6,8],[2,8]],[[2,1],[4,5]]]]
[7,[5,[[3,8],[1,4]]]]
[[2,[2,2]],[8,[8,1]]]
[2,9]
[1,[[[9,3],9],[[9,0],[0,7]]]]
[[[5,[7,4]],7],1]
[[[[4,2],2],6],[8,7]]`,
			"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ns := ParseList(tt.input)
			s := Sum(ns)
			require.Equal(t, tt.expect, s.String())
		})
	}
}

const exampleInput = `[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]
[[[5,[2,8]],4],[5,[[9,9],0]]]
[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]
[[[6,[0,7]],[0,9]],[4,[9,[9,0]]]]
[[[7,[6,4]],[3,[1,3]]],[[[5,5],1],9]]
[[6,[[7,3],[3,2]]],[[[3,8],[5,7]],4]]
[[[[5,4],[7,7]],8],[[8,3],8]]
[[9,3],[[9,9],[6,[4,9]]]]
[[2,[[7,7],7]],[[5,8],[[9,3],[0,2]]]]
[[[[5,2],5],[8,[3,7]]],[[5,[7,5]],[4,4]]]`

func TestNumber_Magnitude(t *testing.T) {
	tests := []struct {
		input  string
		expect int
	}{
		{"[[1,2],[[3,4],5]]", 143},
		{"[[[[0,7],4],[[7,8],[6,0]]],[8,1]]", 1384},
		{"[[[[1,1],[2,2]],[3,3]],[4,4]]", 445},
		{"[[[[3,0],[5,3]],[4,4]],[5,5]]", 791},
		{"[[[[5,0],[7,4]],[5,5]],[6,6]]", 1137},
		{"[[[[8,7],[7,7]],[[8,6],[7,7]]],[[[0,7],[6,6]],[8,7]]]", 3488},
	}
	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			m := ParseString(tt.input).Magnitude()
			require.Equal(t, tt.expect, m)
		})
	}

	t.Run("example", func(t *testing.T) {
		m := Sum(ParseList(exampleInput)).Magnitude()
		require.Equal(t, 4140, m)
	})

	t.Run("puzzle A", func(t *testing.T) {
		m := Sum(ParseList(puzzleInput)).Magnitude()
		require.Equal(t, 4184, m)
	})
}

func TestLargestSum(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := LargestSum(ParseList(exampleInput))
		require.Equal(t, 3993, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		got := LargestSum(ParseList(puzzleInput))
		require.Equal(t, 4731, got)
	})
}
