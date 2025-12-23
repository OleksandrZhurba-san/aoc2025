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

func inBounds(grid [][]byte, p P) bool {
	return p.r >= 0 && p.r < len(grid) && p.c >= 0 && p.c < len(grid[p.r])
}

func isRoll(grid [][]byte, p P) bool {
	return grid[p.r][p.c] == '@'
}

func countAdjRolls(grid [][]byte, at P) int {
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

func replaceRolls(grid [][]byte, ps []P) {
	for _, p := range ps {
		grid[p.r][p.c] = 'x'
	}
}

func main() {

	path := input.GetArgs("data/day4.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.Split(strings.TrimSpace(string(data)), "\n")

	grid := make([][]byte, len(parsedStrs))
	for i := range parsedStrs {
		grid[i] = []byte(parsedStrs[i])
	}
	xCount := 0

	for {
		coord := make([]P, 0)

		for r := range grid {
			for c := range grid[r] {
				at := P{r, c}

				if !isRoll(grid, at) {
					continue
				}

				if countAdjRolls(grid, at) < 4 {
					coord = append(coord, at)
				}
			}
		}
		if len(coord) == 0 {
			break
		}
		xCount += len(coord)
		replaceRolls(grid, coord)
	}

	fmt.Printf("%d\n", xCount)

}
