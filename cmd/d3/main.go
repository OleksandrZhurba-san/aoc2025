package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strings"
)

func main() {

	path := input.GetArgs("data/day3.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.Split(strings.TrimSpace(string(data)), "\n")
	sum := 0
	for _, v := range parsedStrs {
		num := 0
		res := solution(v, 12)

		for _, d := range res {
			num = num*10 + d
		}
		sum += num
	}

	fmt.Printf("Sum: %d\n", sum)
}

func solution(v string, k int) []int {
	result := make([]int, 0, k)
	start := 0

	for k > 0 {
		end := len(v) - k
		best := -1
		bestIdx := -1

		for i := start; i <= end; i++ {
			d := int(v[i] - '0')
			if d > best {
				best = d
				bestIdx = i
			}
		}
		result = append(result, best)
		start = bestIdx + 1
		k--
	}

	fmt.Println(result)
	return result
}
