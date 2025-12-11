package day11

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadStringLines(path[f])
	adj := createAdjMap(data)

	var dfs func(s string) int
	dfs = func(s string) int {
		if s == "out" {
			return 1
		}

		total := 0

		for _, v := range adj[s] {
			total += dfs(v)
		}

		return total
	}

	sum := dfs("you")

	fmt.Printf("Day 11 | Part 1: %d\n", sum)
}
