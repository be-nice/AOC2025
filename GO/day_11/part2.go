package day11

import (
	"fmt"

	"aoc2025/utils"
)

type state struct {
	key string
	f   bool
	d   bool
}

func Part2(f string) {
	data := utils.ReadStringLines(path[f])
	adj := createAdjMap(data)
	memo := make(map[state]int)

	var dfs func(string, bool, bool) int
	dfs = func(s string, f, d bool) int {
		if val, ok := memo[state{s, f, d}]; ok {
			return val
		}

		if s == "out" {
			if f && d {
				return 1
			}
			return 0
		}

		total := 0

		for _, v := range adj[s] {
			nf := f || v == "fft"
			nd := d || v == "dac"

			total += dfs(v, nf, nd)
		}

		memo[state{s, f, d}] = total

		return total
	}

	sum := dfs("svr", false, false)

	fmt.Println(sum)
}
