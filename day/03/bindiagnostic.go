package main

import "strconv"

func binaries(codes []string) (bs []int, width int) {
	width = len(codes[0])
	bs = make([]int, len(codes))
	for i, code := range codes {
		bi, _ := strconv.ParseInt(code, 2, 32)
		bs[i] = int(bi)
	}
	return
}

func countInPosition(codes []int, width, pos int) (count int) {
	mask := 1 << (width - pos - 1)
	for _, code := range codes {
		if code&mask != 0 {
			count++
		}
	}
	return
}

func countInPositions(codes []int, width int) []int {
	counts := make([]int, width)
	for pos := 0; pos < width; pos++ {
		counts[pos] = countInPosition(codes, width, pos)
	}
	return counts
}

func decodeRates(codes []string) (gamma, epsilon int) {
	counts := countInPositions(binaries(codes))
	h := len(codes) / 2

	epsilon = 1
	for i := range counts {
		gamma <<= 1
		epsilon <<= 1

		count := counts[i]
		if count > h {
			gamma |= 1
		}
	}
	epsilon -= 1 + gamma
	return
}

func filter(codes []int, width, pos int, keepCommon bool) (filtered []int) {
	count := countInPosition(codes, width, pos)

	var keepOnes bool
	if keepCommon {
		keepOnes = 2*count >= len(codes)
	} else {
		keepOnes = 2*count < len(codes)
	}

	mask := 1 << (width - pos - 1)
	for _, code := range codes {
		isOne := (code & mask) == mask
		if isOne == keepOnes {
			filtered = append(filtered, code)
		}
	}
	return
}

func repeatedFilter(codes []int, width int, keepCommon bool) int {
	for pos := 0; pos < width; pos++ {
		codes = filter(codes, width, pos, keepCommon)
		if len(codes) == 1 {
			break
		}
	}
	return codes[0]
}

func decodeGases(codes []string) (ox, co2 int) {
	bs, width := binaries(codes)
	ox = repeatedFilter(bs, width, true)
	co2 = repeatedFilter(bs, width, false)
	return
}
