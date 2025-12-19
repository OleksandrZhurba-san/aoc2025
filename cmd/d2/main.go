package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strconv"
	"strings"
)

func main() {

	path := input.GetArgs("data/day2.txt")
	data, err := input.Read(path)
	if err != nil {
		panic(err)
	}
	parsedStrs := strings.SplitSeq(strings.TrimSpace(string(data)), ",")
	sum := 0

	/* evenLen := func(s string) bool {
		return len(s)%2 == 0
	} */

	for v := range parsedStrs {
		if v == "" {
			continue
		}

		parts := strings.SplitN(v, "-", 2)
		lStr, rStr := parts[0], parts[1]

		/* if len(lStr) == len(rStr) && !evenLen(lStr) {
			continue
		} */

		start, err := strconv.Atoi(lStr)
		if err != nil {
			panic(err)
		}

		end, err := strconv.Atoi(rStr)
		if err != nil {
			panic(err)
		}

		for n := start; n <= end; n++ {
			/* s := strconv.Itoa(n)
			if len(s)%2 != 0 {
				continue
			}
			half := len(s) / 2
			if s[:half] == s[half:] {
				sum += n
			} */
			// 999
			s := strconv.Itoa(n)       // "1010"
			l := len(s)                // 4
			for i := 0; i < l/2; i++ { // l/2 = 2
				t := ""
				chunk := s[:i+1]                    //10
				for j := 0; j < l/len(chunk); j++ { //
					t += chunk // 1010
				}
				if t == s {
					sum += n
					break
				}
			}
		}
	}
	fmt.Printf("Sum of Invalid IDs: %d\n", sum)
}
