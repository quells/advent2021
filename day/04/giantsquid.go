package main

import (
	"regexp"
	"strconv"
	"strings"
)

const marked = -1

type board struct {
	rows [][]int
}

func (b *board) mark(v int) {
	for j, row := range b.rows {
		for i, cell := range row {
			if cell == v {
				row[i] = marked
			}
		}
		b.rows[j] = row
	}
}

func (b board) isComplete() bool {
	colCounts := make([]int, 5)
	for _, row := range b.rows {
		rowCount := 0
		for i, cell := range row {
			if cell == marked {
				rowCount++
				colCounts[i]++
			}
		}
		if rowCount == 5 {
			return true
		}
	}
	for _, cc := range colCounts {
		if cc == 5 {
			return true
		}
	}

	return false
}

func (b board) sum() (s int) {
	for _, row := range b.rows {
		for _, cell := range row {
			if cell != marked {
				s += cell
			}
		}
	}
	return
}

var bingoRowPattern = regexp.MustCompile(`^\s*(\d+)\s+(\d+)\s+(\d+)\s+(\d+)\s+(\d+)$`)

func initBingoState(input string) (order []int, boards []board) {
	lines := strings.Split(input, "\n")

	orderStr := strings.Split(lines[0], ",")
	order = make([]int, len(orderStr))
	for i, o := range orderStr {
		order[i], _ = strconv.Atoi(o)
	}

	lines = lines[1:]
	for {
		var b board
		for i := 1; i <= 5; i++ {
			rowStr := bingoRowPattern.FindAllStringSubmatch(lines[i], 1)[0][1:]
			row := make([]int, 5)
			for j, rs := range rowStr {
				row[j], _ = strconv.Atoi(rs)
			}
			b.rows = append(b.rows, row)
		}

		boards = append(boards, b)

		lines = lines[6:]
		if len(lines) == 0 {
			break
		}
	}

	return
}

func playBingo(order []int, boards []board) (int, board) {
	for _, called := range order {
		for i := range boards {
			boards[i].mark(called)
			if boards[i].isComplete() {
				return called, boards[i]
			}
		}
	}
	return -1, board{}
}

func loseBingo(order []int, boards []board) ([]int, []board) {
	for o, called := range order {
		var toRemove []int

		for i := range boards {
			boards[i].mark(called)
			if boards[i].isComplete() {
				toRemove = append(toRemove, i)
			}
		}

		for i := len(toRemove) - 1; i >= 0; i-- {
			idx := toRemove[i]
			if idx+1 == len(boards) {
				boards = boards[:idx]
			} else {
				boards = append(boards[:idx], boards[idx+1:]...)
			}
		}

		if len(boards) == 1 {
			return order[o:], boards
		}
	}
	return nil, nil
}
