package main

import (
	"regexp"
	"strconv"
)

var (
	cmdPattern = regexp.MustCompile(`^(\w)\w+ (\d+)$`)
)

func parseCmd(s string) (dx, dd int) {
	match := cmdPattern.FindAllStringSubmatch(s, 1)[0]
	dir := match[1]
	dist, _ := strconv.Atoi(match[2])
	switch dir {
	case "f":
		return dist, 0
	case "u":
		return 0, -dist
	case "d":
		return 0, dist
	}
	panic("invalid input")
}

type pos struct {
	x     int
	depth int
	aim   int
}

func (p *pos) followCmds(cmds []string) {
	for _, cmd := range cmds {
		dx, dd := parseCmd(cmd)
		p.x += dx
		p.depth += dd
	}
}

func (p *pos) followAimCmds(cmds []string) {
	for _, cmd := range cmds {
		dx, dd := parseCmd(cmd)
		p.aim += dd
		p.x += dx
		p.depth += p.aim * dx
	}
}
