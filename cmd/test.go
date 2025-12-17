package main

import (
	"aoc2025/internal/input"
	"fmt"
	"strings"
)

func main() {
	s, err := input.Read(input.GetArgs(""))
	if err != nil {
		panic(err)
	}
	s = strings.TrimSpace(s)
	for part := range strings.SplitSeq(s, "\n") {
		fmt.Printf("%s\n", part)
	}

	str := "!#!!asdfaf!$!!"
	fmt.Printf("%s", strings.Trim(str, "$!"))
}
