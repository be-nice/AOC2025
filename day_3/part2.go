package day3

import (
	"fmt"
	"strconv"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadByteLines(path[f])

	sum := 0

	for _, line := range data {
		num, _ := strconv.Atoi(string(subSeq(line)))
		sum += num
	}

	fmt.Println(sum)
}

func subSeq(digits []byte) []byte {
	toRemove := len(digits) - 12
	stack := make([]byte, 0, 12)

	for _, c := range digits {

		for len(stack) > 0 && toRemove > 0 && stack[len(stack)-1] < c {
			stack = stack[:len(stack)-1]
			toRemove--
		}

		stack = append(stack, c)
	}

	return stack[:12]
}
