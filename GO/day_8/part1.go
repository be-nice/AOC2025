package day8

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadStringLines(path[f])

	var firstN int
	if f == "at" || f == "bt" {
		firstN = 10
	} else {
		firstN = 1000
	}

	dists := calcDistances(data)

	setMap := make([]map[int]struct{}, 0)

	for _, pair := range dists[:firstN] {
		pair.buildSets(&setMap, -1)
	}

	fmt.Println(utils.MultList(getTopN(3, setMap)))
}
