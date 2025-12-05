package day5

import (
	"fmt"
	"strings"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadString(path[f])
	count := 0

	parts := strings.Split(data, "\n\n")

	ranges := buildRanges(parts[0])

	for _, inter := range ranges {
		count += inter.end - inter.start + 1
	}

	fmt.Println(count)
}
