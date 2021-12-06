package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = "3,4,3,1,2"

//go:embed 06.txt
var puzzleInput string

func Test_newSchool(t *testing.T) {
	s := newSchool(exampleInput)
	want := uint64(2) // two 3's
	got := s[3]
	require.Equal(t, want, got)
}

func TestSchool_count(t *testing.T) {
	s := newSchool(exampleInput)
	want := uint64(5)
	got := s.count()
	require.Equal(t, want, got)
}

func TestSchool_tick(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		s := newSchool(exampleInput)

		for i := 1; i <= 18; i++ {
			s = s.tick()
		}
		require.Equal(t, uint64(26), s.count())

		for i := 19; i <= 80; i++ {
			s = s.tick()
		}
		require.Equal(t, uint64(5934), s.count())

		for i := 81; i <= 256; i++ {
			s = s.tick()
		}
		require.Equal(t, uint64(26984457539), s.count())
	})

	t.Run("puzzle", func(t *testing.T) {
		s := newSchool(puzzleInput)
		for i := 1; i <= 80; i++ {
			s = s.tick()
		}
		require.Equal(t, uint64(388419), s.count())

		for i := 81; i <= 256; i++ {
			s = s.tick()
		}
		require.Equal(t, uint64(1740449478328), s.count())
	})
}
