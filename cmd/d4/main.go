package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strings"
)

type P struct {
	r, c int
}

var neighbors = []P{
	{-1, -1}, {-1, 0}, {-1, 1},
	{0, -1}, {0, 1},
	{1, -1}, {1, 0}, {1, 1},
}

func inBounds(grid []string, p P) bool {
	return p.r >= 0 && p.r < len(grid) && p.c >= 0 && p.c < len(grid[p.r])
}

func isRoll(grid []string, p P) bool {
	return grid[p.r][p.c] == '@'
}

func countAdjRolls(grid []string, at P) int {
	count := 0
	for _, d := range neighbors {
		np := P{at.r + d.r, at.c + d.c}
		if !inBounds(grid, np) {
			continue
		}
		if isRoll(grid, np) {
			count++
		}
	}
	return count
}

func main() {

	path := input.GetArgs("data/day4.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.Split(strings.TrimSpace(string(data)), "\n")

	xCount := 0

	for r := range parsedStrs {
		for c := range parsedStrs[r] {
			at := P{r, c}

			if !isRoll(parsedStrs, at) {
				continue
			}

			if countAdjRolls(parsedStrs, at) < 4 {
				xCount++
			}
		}
	}
	fmt.Printf("%d\n", xCount)
}
