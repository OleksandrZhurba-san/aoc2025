package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strings"
)

func solution() {

}

func main() {

	path := input.GetArgs("data/day4.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.Split(strings.TrimSpace(string(data)), "\n")

	solution()

	xCount := 0

	rows := len(parsedStrs)
	for r := range parsedStrs {
		cols := len(parsedStrs[r])
		for c := range cols {
			if parsedStrs[r][c] == '@' {
				count := 0
				for dr := -1; dr <= 1; dr++ {
					for dc := -1; dc <= 1; dc++ {
						if dr == 0 && dc == 0 {
							continue
						}
						nr := r + dr
						nc := c + dc

						if nr < 0 || nr >= rows || nc < 0 || nc >= cols {
							continue
						}
						if parsedStrs[nr][nc] == '@' {
							count++
						}
					}
				}
				if count < 4 {
					xCount++
				}
			}
		}
	}
	fmt.Printf("%d\n", xCount)
}
