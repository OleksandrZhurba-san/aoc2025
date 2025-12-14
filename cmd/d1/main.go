package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strconv"
)

func main() {
	capacity := 100
	point := 50
	count := 0

	data, err := input.Read("data/day1.txt")
	if err != nil {
		panic(err)
	}

	for _, v := range data {
		if v == "" {
			continue
		}

		current, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		switch v[0] {
		case 'R':
			if point = (point + current) % capacity; point == 0 {
				count++
			}
		case 'L':
			if point = (point - current + capacity) % capacity; point == 0 {
				count++
			}
		}
	}
	fmt.Println(count)
}
