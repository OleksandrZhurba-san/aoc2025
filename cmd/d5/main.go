package main

import (
	"aoc2025/internal/input"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func mergeRanges(sortedRanges []Range) []Range {
	mergedRanges := make([]Range, 0, len(sortedRanges))
	cur := sortedRanges[0]
	for _, r := range sortedRanges[1:] {
		if r.start <= cur.end {
			if r.end > cur.end {
				cur.end = r.end
			}
		} else {
			mergedRanges = append(mergedRanges, cur)
			cur = r
		}
	}
	mergedRanges = append(mergedRanges, cur)
	return mergedRanges
}

func findFreshIngredientId(mergedRanges []Range, ids []int) []int {
	result := make([]int, 0)
	for _, r := range mergedRanges {
		for _, id := range ids {
			if id >= r.start && id <= r.end {
				result = append(result, id)
			}
		}
	}
	return result
}

func main() {

	path := input.GetArgs("data/day5.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.Split(strings.TrimSpace(string(data)), "\n\n")

	ranges := strings.Split(parsedStrs[0], "\n")
	ids := strings.Split(parsedStrs[1], "\n")
	intIds := make([]int, 0, len(ids))
	for _, id := range ids {
		i, err := strconv.Atoi(id)
		if err != nil{
			continue
		}
		intIds = append(intIds, i)
	}
	powerRanges := make([]Range, 0, len(ranges))
	for _, v := range ranges {
		parts := strings.Split(v, "-")
		s, _ := strconv.Atoi(parts[0])
		e, _ := strconv.Atoi(parts[1])
		powerRanges = append(powerRanges, Range{start: s, end: e})
	}
	sort.Slice(powerRanges, func(i, j int) bool {
		if powerRanges[i].start == powerRanges[j].start {
			return powerRanges[i].end < powerRanges[j].end
		}
		return powerRanges[i].start < powerRanges[j].start
	})

	mergedRanges := mergeRanges(powerRanges)
	freshIds := findFreshIngredientId(mergedRanges, intIds)

	sum := 0
	for _, r := range mergedRanges {
		sum += r.end - r.start + 1
	}


	fmt.Println(len(freshIds))
	fmt.Println(sum)

}
