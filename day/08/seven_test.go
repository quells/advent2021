package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

const singleInput = "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"

const exampleInput = `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe
edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc
fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg
fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb
aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea
fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb
dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe
bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef
egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb
gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce`

//go:embed 08.txt
var puzzleInput string

func Test_countEasyDigits(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ws := parseLines(exampleInput)
		got := countEasyDigits(ws)
		want := 26
		require.Equal(t, want, got)
	})

	t.Run("puzzle A", func(t *testing.T) {
		ws := parseLines(puzzleInput)
		got := countEasyDigits(ws)
		want := 381
		require.Equal(t, want, got)
	})
}

func TestWiring_deduce(t *testing.T) {
	t.Run("in order", func(t *testing.T) {
		w := parseLine("abcefg cf acdeg acdfg bcdf abdfg abdefg acf abcdefg abcdfg | cf acdeg acdfg bcdf")
		got := w.deduce()
		want := wireMap{119, 36, 93, 109, 46, 107, 123, 37, 127, 111}
		require.Equal(t, want, got)
	})

	t.Run("single", func(t *testing.T) {
		w := parseLine(singleInput)
		got := w.deduce()
		want := wireMap{95, 3, 109, 47, 51, 62, 126, 11, 127, 63}
		require.Equal(t, want, got)
	})
}

func TestWiring_decode(t *testing.T) {
	t.Run("single", func(t *testing.T) {
		w := parseLine(singleInput)
		got := w.decode()
		want := 5353
		require.Equal(t, want, got)
	})

	t.Run("example", func(t *testing.T) {
		got := decodeAndSum(exampleInput)
		want := 61229
		require.Equal(t, want, got)
	})

	t.Run("puzzle B", func(t *testing.T) {
		got := decodeAndSum(puzzleInput)
		want := 1023686
		require.Equal(t, want, got)
	})
}

func BenchmarkDecode(b *testing.B) {
	w := parseLine(singleInput)
	want := 5353

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := w.decode()
		if got != want {
			b.Fail()
		}
	}
}
