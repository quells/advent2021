package main

import (
	"strconv"
	"strings"
)

type school map[int]uint64 // days til spawn : count

func newSchool(input string) (fish school) {
	fish = make(school)
	for _, s := range strings.Split(input, ",") {
		v, _ := strconv.Atoi(s)
		oldCount := fish[v]
		fish[v] = oldCount + 1
	}
	return
}

func (s school) tick() (next school) {
	next = make(school)
	for counter, count := range s {
		switch counter {
		case 0:
			// reset spawn countdown for existing fish
			{
				oldCount := next[6]
				newCount := oldCount + count
				next[6] = newCount
			}
			// add babies
			{
				oldCount := next[8]
				newCount := oldCount + count
				next[8] = newCount
			}

		default:
			newCounter := counter - 1
			oldCount := next[newCounter]
			newCount := oldCount + count
			next[newCounter] = newCount
		}
	}
	return
}

func (s school) count() (total uint64) {
	for _, count := range s {
		total += count
	}
	return
}
