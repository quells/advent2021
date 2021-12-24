package main

import (
	"bytes"
	"strings"
)

type rule struct {
	pair   []byte
	insert byte
}

func parse(input string) (state []byte, rules []rule) {
	lines := strings.Split(input, "\n")

	state = []byte(lines[0])

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		pair := []byte{parts[0][0], parts[0][1]}
		insert := parts[1][0]
		rules = append(rules, rule{pair, insert})
	}

	return
}

func apply(state []byte, rules []rule) []byte {
	applied := []byte{state[0]}
	for i := 0; i < len(state)-1; i++ {
		for _, r := range rules {
			if bytes.Equal(state[i:i+2], r.pair) {
				applied = append(applied, r.insert)
				break
			}
		}
		applied = append(applied, state[i+1])
	}
	return applied
}

func countFreqs(state []byte) (most, least int) {
	freqs := make([]int, 128)
	for _, s := range state {
		freqs[int(s)]++
	}

	least = len(state)
	for _, f := range freqs {
		if f > most {
			most = f
		}
		if f > 0 && f < least {
			least = f
		}
	}
	return
}
