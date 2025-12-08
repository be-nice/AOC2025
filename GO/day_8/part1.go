package day8

import (
	"fmt"

	"aoc2025/utils"
)

func Part1(f string) {
	data := utils.ReadStringLines(path[f])

	var includeAmount int
	if f == "at" || f == "bt" {
		includeAmount = 10
	} else {
		includeAmount = 1000
	}

	dists := calcDistances(data)

	setMap := make([]map[int]struct{}, 0)

	for _, val := range dists[:includeAmount] {
		val.buildSets(&setMap, -1)
	}

	top := []int{0, 0, 0}

	for _, m := range setMap {
		switch {
		case len(m) > top[0]:
			top[2] = top[1]
			top[1] = top[0]
			top[0] = len(m)
		case len(m) > top[1]:
			top[2] = top[1]
			top[1] = len(m)
		case len(m) > top[2]:
			top[2] = len(m)
		}
	}

	fmt.Println(utils.MultList(top))
}
