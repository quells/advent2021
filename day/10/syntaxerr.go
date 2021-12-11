package main

import (
	"sort"
	"strings"
)

type stack []rune

func (s stack) push(b rune) stack {
	return append(s, b)
}

func (s stack) peek() rune {
	return s[len(s)-1]
}

func (s stack) pop() stack {
	return s[:len(s)-1]
}

func checkCorruption(line string) (rune, stack) {
	var s stack
	for _, c := range line {
		switch c {
		case '<', '{', '[', '(':
			s = s.push(c)
		case '>':
			if s.peek() != '<' {
				return c, s
			}
			s = s.pop()
		case '}':
			if s.peek() != '{' {
				return c, s
			}
			s = s.pop()
		case ']':
			if s.peek() != '[' {
				return c, s
			}
			s = s.pop()
		case ')':
			if s.peek() != '(' {
				return c, s
			}
			s = s.pop()
		default:
			panic("invalid character")
		}
	}

	return 0, s
}

func scoreCorruption(lines []string) (score int) {
	scores := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}
	for _, line := range lines {
		c, _ := checkCorruption(line)
		score += scores[c]
	}
	return
}

func discardCorrupted(lines []string) (incomplete []stack) {
	for _, line := range lines {
		c, s := checkCorruption(line)
		if c == 0 {
			incomplete = append(incomplete, s)
		}
	}
	return
}

func scoreIncomplete(s stack) (score int) {
	for i := len(s) - 1; i >= 0; i-- {
		score *= 5
		switch s[i] {
		case '(':
			score += 1
		case '[':
			score += 2
		case '{':
			score += 3
		case '<':
			score += 4
		}
	}
	return
}

func scoreAutocompletes(s string) int {
	incomplete := discardCorrupted(strings.Split(s, "\n"))

	scores := make([]int, len(incomplete))
	for idx, i := range incomplete {
		scores[idx] = scoreIncomplete(i)
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}
