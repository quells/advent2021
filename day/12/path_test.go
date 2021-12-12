package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const exampleInput = `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`

const puzzleInput = `EG-bj
LN-end
bj-LN
yv-start
iw-ch
ch-LN
EG-bn
OF-iw
LN-yv
iw-TQ
iw-start
TQ-ch
EG-end
bj-OF
OF-end
TQ-start
TQ-bj
iw-LN
EG-ch
yv-iw
KW-bj
OF-ch
bj-ch
yv-TQ`

func Test_countPaths(t *testing.T) {
	t.Run("example", func(t *testing.T) {
		ps := parsePaths(exampleInput)
		got := countPaths(ps)
		want := 19
		require.Equal(t, want, got)
	})

	t.Run("puzzle A", func(t *testing.T) {
		ps := parsePaths(puzzleInput)
		got := countPaths(ps)
		want := 4659
		require.Equal(t, want, got)
	})
}

func Benchmark_countPaths(b *testing.B) {
	ps := parsePaths(puzzleInput)
	want := 4659

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		got := countPaths(ps)
		if got != want {
			b.Fail()
		}
	}
}
