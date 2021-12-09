package main

import (
	"math/bits"
	"strings"
)

/*
_GFEDCBA : pack signals into byte

 AAA
B   C
 DDD
E   F
 GGG
*/

type wiring struct {
	signals [10]byte
	digits  [4]byte
}

func packSignal(s string) (p byte) {
	for _, c := range s {
		switch c {
		case 'a':
			p |= 1 << 0
		case 'b':
			p |= 1 << 1
		case 'c':
			p |= 1 << 2
		case 'd':
			p |= 1 << 3
		case 'e':
			p |= 1 << 4
		case 'f':
			p |= 1 << 5
		case 'g':
			p |= 1 << 6
		default:
			panic("invalid character")
		}
	}
	return
}

func parseLine(line string) (w wiring) {
	split := strings.Split(line, " | ")
	for i, si := range strings.Split(split[0], " ") {
		w.signals[i] = packSignal(si)
	}
	for i, di := range strings.Split(split[1], " ") {
		w.digits[i] = packSignal(di)
	}
	return
}

func parseLines(s string) (ws []wiring) {
	lines := strings.Split(s, "\n")
	ws = make([]wiring, len(lines))
	for i, li := range lines {
		ws[i] = parseLine(li)
	}
	return
}

func countEasyDigits(ws []wiring) (easyDigits int) {
	for _, w := range ws {
		for _, d := range w.digits {
			switch bits.OnesCount8(d) {
			case 2: // 1
				easyDigits++
			case 4: // 4
				easyDigits++
			case 3: // 7
				easyDigits++
			case 7: // 8
				easyDigits++
			}
		}
	}
	return
}

type wireMap [10]byte

func (wm wireMap) decode(b byte) int {
	for idx, bi := range wm {
		if bi == b {
			return idx
		}
	}
	panic("incomplete wire map")
}

func (wire wiring) deduce() (m wireMap) {
	var one, four, seven, eight byte
	for _, signal := range wire.signals {
		switch bits.OnesCount8(signal) {
		case 2:
			one = signal
		case 4:
			four = signal
		case 3:
			seven = signal
		case 7:
			eight = signal
		}
	}

	var e, nine byte
	{
		L := one ^ four ^ seven ^ eight // lower L
		var hi byte = 1 << 6
		for {
			if L&hi != 0 {
				break
			}
			hi >>= 1
		}
		lo := L ^ hi

		maybeNine := eight ^ hi
		for _, signal := range wire.signals {
			if signal == maybeNine {
				e = hi
				// g = lo
				nine = signal
				break
			}
		}
		if nine == 0 {
			e = lo
			// g = hi
			nine = eight ^ lo
		}
	}

	var c, three, five byte
	for _, signal := range wire.signals {
		segment := signal ^ nine
		if bits.OnesCount8(signal) == 5 && bits.OnesCount8(segment) == 1 {
			if segment&one == 0 {
				// b = segment
				three = signal
			} else {
				c = segment
				five = signal
			}
		}
	}
	d := (three & four) ^ one
	zero := eight ^ d
	six := five | e

	E := eight ^ one
	bars := E & three
	two := bars | c | e

	m[0] = zero
	m[1] = one
	m[2] = two
	m[3] = three
	m[4] = four
	m[5] = five
	m[6] = six
	m[7] = seven
	m[8] = eight
	m[9] = nine

	return
}

func (wire wiring) decode() (code int) {
	wm := wire.deduce()
	for _, digit := range wire.digits {
		code *= 10
		code += wm.decode(digit)
	}
	return
}

func decodeAndSum(input string) (sum int) {
	ws := parseLines(input)
	for _, w := range ws {
		sum += w.decode()
	}
	return
}
