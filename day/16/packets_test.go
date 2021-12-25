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

func TestPacket_value(t *testing.T) {
	tests := []struct {
		name string
		hex  string
		want int
	}{
		{"sum", "C200B40A82", 3},          // 1 + 2 = 3
		{"product", "04005AC33890", 54},   // 6 * 9 = 54
		{"minimum", "880086C3E88112", 7},  // min(7, 8, 9)
		{"maximum", "CE00C43D881120", 9},  // max(7, 8, 9)
		{"less than", "D8005AC2A8F0", 1},  // 5 < 15
		{"greater than", "F600BC2D8F", 0}, // !(5 > 15)
		{"equal to", "9C005AC2F8F0", 0},   // !(5 == 15)

		{"arithmetic comparison", "9C0141080250320F1802104A08", 1}, // 1 + 3 == 2 * 2

		{"puzzle B", puzzleInput, 912901337844},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bs := parse(tt.hex)
			ps, err := bs.packets()

			require.NoError(t, err)
			require.Len(t, ps, 1)

			got := ps[0].value()
			require.Equal(t, tt.want, got)
		})
	}
}
