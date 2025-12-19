package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	capacity := 100
	point := 50
	count := 0

	path := input.GetArgs("data/day1.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")

	for v := range parsedStrs {
		if v == "" {
			continue
		}

		current, err := strconv.Atoi(v[1:])
		if err != nil {
			panic(err)
		}

		switch v[0] {
		case 'R':
			/* if point = (point + current) % capacity; point == 0 {
				count++
			} */
			hits := (point + current) / capacity
			count += hits
			point = (point + current) % capacity
		case 'L':
			/* if point = (point - current + capacity) % capacity; point == 0 {
				count++
			} */
			hits := 0
			if point == 0 {
				hits = current / capacity
			} else {
				hits = (current + (capacity - point)) / capacity
			}

			count += hits
			point = (point - (current % capacity) + capacity) % capacity
		}
	}
	fmt.Println(count)
}
