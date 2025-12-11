package day11

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadStringLines(path[f])
	adj := createMap(data)

	var dfs func(s string) int
	dfs = func(s string) int {
		if s == "out" {
			return 1
		}

		total := 0

		for k := range adj[s] {
			total += dfs(k)
		}

		return total
	}

	sum := dfs("you")

	fmt.Println(sum)
}
