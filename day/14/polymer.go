package main

import (
	"encoding/binary"
	"strings"
)

type rule struct {
	pair   uint16
	insert byte
}

func parse(input string) (state []byte, rules []rule) {
	lines := strings.Split(input, "\n")

	state = []byte(lines[0])

	for _, line := range lines[2:] {
		parts := strings.Split(line, " -> ")
		pair := binary.BigEndian.Uint16([]byte(parts[0]))
		insert := parts[1][0]
		rules = append(rules, rule{pair, insert})
	}

	return
}

type polymer struct {
	pairCounts map[uint16]int
	last       byte
}

func initPolymer(state []byte) *polymer {
	pairCounts := make(map[uint16]int)
	for i := 0; i < len(state)-1; i++ {
		pair := binary.BigEndian.Uint16(state[i : i+2])
		pairCounts[pair]++
	}
	last := state[len(state)-1]
	return &polymer{pairCounts, last}
}

func (p *polymer) apply(rules []rule) {
	inserted := make(map[uint16]int)
	for _, r := range rules {
		count := p.pairCounts[r.pair]
		if count > 0 {
			lhs := (r.pair & 0xFF00) | uint16(r.insert)
			rhs := (r.pair & 0x00FF) | (uint16(r.insert) << 8)
			inserted[lhs] += count
			inserted[rhs] += count
			inserted[r.pair] -= count
		}
	}
	for pair, count := range inserted {
		p.pairCounts[pair] += count
		if p.pairCounts[pair] == 0 {
			delete(p.pairCounts, pair)
		}
	}
}

func (p *polymer) countFreqs() (most, least int) {
	freqs := make([]int, 128)
	for pair, count := range p.pairCounts {
		lhs := (pair & 0xFF00) >> 8
		freqs[int(lhs)] += count
		least += count
	}
	freqs[int(p.last)]++
	least++

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
