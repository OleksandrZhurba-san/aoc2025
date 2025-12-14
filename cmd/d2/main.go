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

	data, err := input.Read("data/day2.txt")
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
			if b, a := point, (point+current)%capacity; b > a || point == 0 {
				count++
			}
			point = (point + current) % capacity
		case 'L':
			if b, a := point, (point-current+capacity)%capacity; b < a || point == 0 {
				count++
			}
			point = (point - current + capacity) % capacity
		}
	}
	fmt.Println(count)
}
