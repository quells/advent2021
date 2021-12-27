package main

import (
	"fmt"
	"strings"
)

func newLiteral(x int, parent *Number) *Number {
	return &Number{Literal: x, parent: parent}
}

func newPair(l, r int, parent *Number) *Pair {
	left := newLiteral(l, parent)
	right := newLiteral(r, parent)
	return &Pair{
		Left:  left,
		Right: right,
	}
}

type Number struct {
	Literal int
	Pair    *Pair

	parent *Number
}

func (n Number) String() string {
	buf := new(strings.Builder)
	n.write(buf)
	return buf.String()
}

func (n Number) Copy() *Number {
	// This sucks but is very easy...
	return ParseString(n.String())
}

func (n Number) write(buf *strings.Builder) {
	if n.Pair == nil {
		fmt.Fprintf(buf, "%d", n.Literal)
	} else {
		p := n.Pair
		buf.WriteByte('[')
		p.Left.write(buf)
		buf.WriteByte(',')
		p.Right.write(buf)
		buf.WriteByte(']')
	}
}

type Pair struct {
	Left  *Number
	Right *Number
}

func ParseList(input string) (l []*Number) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		n := ParseString(line)
		l = append(l, n)
	}
	return
}

func ParseString(input string) *Number {
	return Parse([]byte(input))
}

func Parse(input []byte) *Number {
	num, rest := parseNext(input, Open, nil)
	if len(rest) != 0 {
		panic(fmt.Sprintf("incomplete parse of %q", string(input)))
	}

	return num
}

type State int

const (
	Open State = iota
	Left
	Comma
	Right
	Close
)

func parseNext(input []byte, state State, parent *Number) (num *Number, rest []byte) {
	idx := 0
	var left, right *Number
	for {
		c := input[idx]
		switch state {
		case Open:
			if c != '[' {
				return nil, input
			}
			idx++
			state = Left
			num = &Number{
				parent: parent,
			}

		case Left:
			if c == '[' {
				left, rest = parseNext(input[idx:], Open, num)
				input = rest
				idx = 0
			} else {
				left = newLiteral(int(c-'0'), num)
				idx++
			}
			state = Comma

		case Comma:
			// multi-digit numbers
			if '0' <= c && c <= '9' {
				left.Literal = left.Literal*10 + int(c-'0')
				idx++
				continue
			}

			if c != ',' {
				return nil, input
			}
			idx++
			state = Right

		case Right:
			if c == '[' {
				right, rest = parseNext(input[idx:], Open, num)
				input = rest
				idx = 0
			} else {
				right = newLiteral(int(c-'0'), num)
				idx++
			}
			state = Close

		case Close:
			// multi-digit numbers
			if '0' <= c && c <= '9' {
				right.Literal = right.Literal*10 + int(c-'0')
				idx++
				continue
			}

			if c != ']' {
				return nil, input
			}
			num.Pair = &Pair{
				Left:  left,
				Right: right,
			}
			rest = input[idx+1:]
			return num, rest
		}
	}
}

type serial struct {
	depth int
	num   *Number
}

func (n *Number) walk(depth int) (s []serial) {
	if n.Pair != nil {
		s = append(s, n.Pair.Left.walk(depth+1)...)
		s = append(s, n.Pair.Right.walk(depth+1)...)
	} else {
		s = append(s, serial{depth, n})
	}
	return
}

func Reduce(n *Number) *Number {
	w := n.walk(-1)
	for i, s := range w {
		if s.depth >= 4 {
			pair := s.num.parent
			if 0 < i {
				prev := w[i-1].num
				prev.Literal += pair.Pair.Left.Literal
			}
			if i < len(w)-2 {
				next := w[i+2].num
				next.Literal += pair.Pair.Right.Literal
			}
			pair.Pair = nil
			pair.Literal = 0

			return Reduce(n)
		}
	}

	for _, s := range w {
		if 10 <= s.num.Literal {
			l := s.num.Literal / 2
			r := (s.num.Literal + 1) / 2
			s.num.Literal = 0
			s.num.Pair = newPair(l, r, s.num)

			return Reduce(n)
		}
	}

	return n
}

func Add(a, b *Number) *Number {
	sum := &Number{}

	l := a.Copy()
	r := b.Copy()
	l.parent = sum
	r.parent = sum

	sum.Pair = &Pair{
		Left:  l,
		Right: r,
	}

	return Reduce(sum)
}

func Sum(ns []*Number) *Number {
	switch len(ns) {
	case 0:
		return nil
	case 1:
		return ns[0]
	}

	s := Add(ns[0], ns[1])
	for i := 2; i < len(ns); i++ {
		s = Add(s, ns[i])
	}
	return s
}

func (n *Number) Magnitude() int {
	if n.Pair == nil {
		return n.Literal
	}

	l := n.Pair.Left.Magnitude()
	r := n.Pair.Right.Magnitude()

	return 3*l + 2*r
}

func LargestSum(ns []*Number) (largest int) {
	for i := 0; i < len(ns); i++ {
		for j := 0; j < len(ns); j++ {
			if i == j {
				continue
			}
			{
				s := Add(ns[i], ns[j]).Magnitude()
				if s > largest {
					largest = s
				}
			}
			{
				s := Add(ns[j], ns[i]).Magnitude()
				if s > largest {
					largest = s
				}
			}
		}
	}
	return
}
