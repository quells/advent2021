package main

import (
	_ "embed"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_checkCorruption(t *testing.T) {
	tests := []struct {
		line string
		want rune
	}{
		{"{([(<{}[<>[]}>{[]{[(<()>", '}'},
		{"[[<[([]))<([[{}[[()]]]", ')'},
		{"[{[{({}]{}}([{[{{{}}([]", ']'},
		{"[<(<(<(<{}))><([]([]()", ')'},
		{"<{([([[(<>()){}]>(<<{{", '>'},

		{"[<>({}){}[([])<>]]", 0},     // valid line
		{"[(()[<>])]({[<{<<[]>>(", 0}, // incomplete but not corrupted
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			invalid, _ := checkCorruption(tt.line)
			require.Equal(t, tt.want, invalid)
		})
	}
}

const exampleInput = `[({(<(())[]>[[{[]{<()<>>
[(()[<>])]({[<{<<[]>>(
{([(<{}[<>[]}>{[]{[(<()>
(((({<>}<{<{<>}{[]{[]{}
[[<[([]))<([[{}[[()]]]
[{[{({}]{}}([{[{{{}}([]
{<[[]]>}<{[{[{[]{()[[[]
[<(<(<(<{}))><([]([]()
<{([([[(<>()){}]>(<<{{
<{([{{}}[<[[[<>{}]]]>[]]`

//go:embed 10.txt
var puzzleInput string

func Test_scoreCorruption(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := scoreCorruption(strings.Split(exampleInput, "\n"))
		want := 26397
		require.Equal(t, want, got)
	})

	t.Run("puzzle A", func(t *testing.T) {
		got := scoreCorruption(strings.Split(puzzleInput, "\n"))
		want := 399153
		require.Equal(t, want, got)
	})
}

func Benchmark_scoreCorruption(b *testing.B) {
	lines := strings.Split(puzzleInput, "\n")
	want := 399153

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := scoreCorruption(lines)
		if got != want {
			b.Fail()
		}
	}
}

func Test_discardCorrupted(t *testing.T) {
	incomplete := discardCorrupted(strings.Split(exampleInput, "\n"))
	require.Len(t, incomplete, 5)
}

func Test_scoreIncomplete(t *testing.T) {
	tests := []struct {
		line string
		want int
	}{
		{"[({(<(())[]>[[{[]{<()<>>", 288957},
		{"[(()[<>])]({[<{<<[]>>(", 5566},
		{"(((({<>}<{<{<>}{[]{[]{}", 1480781},
		{"{<[[]]>}<{[{[{[]{()[[[]", 995444},
		{"<{([{{}}[<[[[<>{}]]]>[]]", 294},
	}
	for _, tt := range tests {
		t.Run(tt.line, func(t *testing.T) {
			_, s := checkCorruption(tt.line)
			got := scoreIncomplete(s)
			require.Equal(t, tt.want, got)
		})
	}
}

func Test_scoreAutocompletes(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		got := scoreAutocompletes(exampleInput)
		want := 288957
		require.Equal(t, want, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		got := scoreAutocompletes(puzzleInput)
		want := 2995077699
		require.Equal(t, want, got)
	})
}
