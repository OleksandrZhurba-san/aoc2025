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
	parsedStrs := strings.SplitSeq(strings.TrimSpace(string(data)), "\n")

	sum := 0

	for v := range parsedStrs {
		first := -1
		second := -1
		idx := -1

		for i := 0; i < len(v)-1; i++ {
			digit := int(v[i] - '0')
			if first < digit {
				first = digit
				idx = i
			}
		}
		for _, ch := range v[idx+1:] {
			digit := int(ch - '0')
			if second < digit {
				second = digit
			}
		}
		sum += (first * 10) + second
	}
	fmt.Printf("Sum: %d\n", sum)
}
