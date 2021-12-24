package main

import (
	_ "embed"
	"testing"

	"github.com/stretchr/testify/require"
)

//go:embed 16.txt
var puzzleInput string

func Test_parse(t *testing.T) {
	t.Run("literal", func(t *testing.T) {
		input := "D2FE28"
		bs := parse(input)
		want := bitstream([]byte{1, 1, 0, 1, 0, 0, 1, 0, 1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 1, 0, 1, 0, 0, 0})
		require.Equal(t, want, bs)
	})
}

func TestBitstream_next(t *testing.T) {
	t.Run("literal", func(t *testing.T) {
		input := "D2FE28"
		bs := parse(input)
		p, err := bs.next()

		require.NoError(t, err)
		require.Equal(t, 6, p.version)
		require.Equal(t, 4, p.id)
		require.Equal(t, 2021, p.literal)
	})

	t.Run("length type 0", func(t *testing.T) {
		input := "38006F45291200"
		bs := parse(input)
		p, err := bs.next()

		require.NoError(t, err)
		require.Equal(t, 1, p.version)
		require.Equal(t, 6, p.id)
		require.Len(t, p.subpackets, 2)
		require.Equal(t, p.subpackets[0].literal, 10)
		require.Equal(t, p.subpackets[1].literal, 20)
	})

	t.Run("length type 1", func(t *testing.T) {
		input := "EE00D40C823060"
		bs := parse(input)
		p, err := bs.next()

		require.NoError(t, err)
		require.Equal(t, 7, p.version)
		require.Equal(t, 3, p.id)
		require.Len(t, p.subpackets, 3)
		require.Equal(t, p.subpackets[0].literal, 1)
		require.Equal(t, p.subpackets[1].literal, 2)
		require.Equal(t, p.subpackets[2].literal, 3)
	})
}

func Test_versionSum(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want int
	}{
		{"nested", "8A004A801A8002F478", 16},
		{"tree 1", "620080001611562C8802118E34", 12},
		{"tree 2", "C0015000016115A2E0802F182340", 23},
		{"roots", "A0016C880162017C3686B18A3D4780", 31},
		{"puzzle A", puzzleInput, 934},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := parse(tt.hex)
			ps, err := bs.packets()
			require.NoError(t, err)
			vs := versionSum(ps)
			require.Equal(t, tt.want, vs)
		})
	}
}
