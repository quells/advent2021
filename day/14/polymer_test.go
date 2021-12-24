package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

var exampleInput = `NNCB

CH -> B
HH -> N
CB -> H
NH -> C
HB -> C
HC -> B
HN -> C
NN -> C
BH -> H
NC -> B
NB -> B
BN -> B
BB -> N
BC -> B
CC -> N
CN -> C`

//go:embed 14.txt
var puzzleInput string

func Test_parse(t *testing.T) {
	state, rules := parse(exampleInput)
	require.Equal(t, "NNCB", string(state))
	require.Len(t, rules, 16)
	require.Equal(t, "CH", string(rules[0].pair))
	require.Equal(t, byte('B'), rules[0].insert)
}

func Test_apply(t *testing.T) {
	state, rules := parse(exampleInput)

	state = apply(state, rules)
	require.Equal(t, "NCNBCHB", string(state))

	state = apply(state, rules)
	require.Equal(t, "NBCCNBBBCBHCB", string(state))

	state = apply(state, rules)
	require.Equal(t, "NBBBCNCCNBBNBNBBCHBHHBCHB", string(state))

	state = apply(state, rules)
	require.Equal(t, "NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB", string(state))
}

func Test_countFreqs(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		state, rules := parse(exampleInput)
		for i := 0; i < 10; i++ {
			state = apply(state, rules)
		}

		most, least := countFreqs(state)
		require.Equal(t, 1749, most)
		require.Equal(t, 161, least)
	})

	t.Run("puzzle A", func(t *testing.T) {
		state, rules := parse(puzzleInput)
		for i := 0; i < 10; i++ {
			state = apply(state, rules)
		}

		most, least := countFreqs(state)
		diff := most - least
		require.Equal(t, 3048, diff)
	})
}
