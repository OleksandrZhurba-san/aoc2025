package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strings"
)

type Grid struct {
	width  int
	height int
	data   [][]byte
}

func NewGridFromLines(lines []string) Grid {
	maxW := 0
	for _, line := range lines {
		if len(line) > maxW {
			maxW = len(line)
		}
	}

	height := len(lines)
	data := make([][]byte, height)

	for y, line := range lines {
		row := make([]byte, maxW)
		copy(row, []byte(line))

		for x := len(line); x < maxW; x++ {
			row[x] = ' '
		}

		data[y] = row
	}

	return Grid{
		data:   data,
		width:  maxW,
		height: height,
	}
}

func (g Grid) At(r, c int) byte {
	return g.data[r][c]
}

func main() {
	path := input.GetArgs("data/day6.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(strings.TrimRight(string(data), "\n"), "\n")

	grid := NewGridFromLines(lines)

	for i, line := range lines {
		fmt.Printf("Line %d: %q\n", i, line)
	}
	bottom := grid.height - 1

	right := grid.width

	for c := right - 1; c >= 0; c-- {
		op := grid.data[bottom][c]
		if op == ' ' {
			continue
		}
		fmt.Printf("block: op=%c cols=[%d..%d)\n", op, c, right)

		right = c 
	}

	/* for i, c := range grid.data[bottom] {
		switch c {
		case '*':
			fmt.Printf("* at grid[%d][%d]\n", bottom, i )
		case '+':
			fmt.Printf("+ at grid[%d][%d]\n", bottom, i )
		}
	} */
 }
