package main

import (
	_ "embed"
	"encoding/binary"
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
	require.Equal(t, binary.BigEndian.Uint16([]byte("CH")), rules[0].pair)
	require.Equal(t, byte('B'), rules[0].insert)
}

func Test_initPolymer(t *testing.T) {
	state, _ := parse(exampleInput)
	p := initPolymer(state)

	require.Len(t, p.pairCounts, 3)
	require.Equal(t, byte('B'), p.last)
}

func TestPolymer_apply(t *testing.T) {
	state, rules := parse(exampleInput)
	p := initPolymer(state)

	p.apply(rules)
	require.Equal(t, initPolymer([]byte("NCNBCHB")), p)

	p.apply(rules)
	require.Equal(t, initPolymer([]byte("NBCCNBBBCBHCB")), p)

	p.apply(rules)
	require.Equal(t, initPolymer([]byte("NBBBCNCCNBBNBNBBCHBHHBCHB")), p)

	p.apply(rules)
	require.Equal(t, initPolymer([]byte("NBBNBNBBCCNBCNCCNBBNBBNBBBNBBNBBCBHCBHHNHCBBCBHCB")), p)
}

func TestPolymer_countFreqs(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		state, rules := parse(exampleInput)
		p := initPolymer(state)
		for i := 0; i < 10; i++ {
			p.apply(rules)
		}

		most, least := p.countFreqs()
		require.Equal(t, 1749, most)
		require.Equal(t, 161, least)

		for i := 10; i < 40; i++ {
			p.apply(rules)
		}
		most, least = p.countFreqs()
		require.Equal(t, 2192039569602, most)
		require.Equal(t, 3849876073, least)
	})

	t.Run("puzzle", func(t *testing.T) {
		state, rules := parse(puzzleInput)
		p := initPolymer(state)
		for i := 0; i < 10; i++ {
			p.apply(rules)
		}

		most, least := p.countFreqs()
		diff := most - least
		require.Equal(t, 3048, diff)

		for i := 10; i < 40; i++ {
			p.apply(rules)
		}
		most, least = p.countFreqs()
		diff = most - least
		require.Equal(t, 3288891573057, diff)
	})
}
