package day11

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadStringLines(path[f])
	adj := createAdjMap(data)

	memo := make(map[string]int)
	var dfs func(s string) int
	dfs = func(s string) int {
		if val, ok := memo[s]; ok {
			return val
		}
		if s == "out" {
			return 1
		}

		total := 0

		for _, v := range adj[s] {
			total += dfs(v)
		}

		memo[s] = total
		return total
	}

	sum := dfs("you")

	fmt.Printf("Day 11 | Part 1: %d\n", sum)
}
