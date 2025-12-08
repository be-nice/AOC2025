package day8

import (
	"fmt"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadStringLines(path[f])
	res := -1

	dists := calcDistances(data)

	setMap := make([]map[int]struct{}, 0)

	for _, pair := range dists {
		res = pair.buildSets(&setMap, len(data))
		if res != -1 {
			break
		}
	}

	fmt.Println(res)
}
