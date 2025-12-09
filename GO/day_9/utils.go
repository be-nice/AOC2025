package day9

import (
	"strings"

	"aoc2025/utils"
)

type pairs struct {
	ax   int
	ay   int
	bx   int
	by   int
	area int
}

type point struct {
	x int
	y int
}

type segment struct {
	start        point
	end          point
	isHorizontal bool
}

func parseNum(s string) (int, int) {
	nums := strings.Split(s, ",")
	n1 := utils.FastAtoi(nums[0])
	n2 := utils.FastAtoi(nums[1])

	return n1, n2
}

func calcArea(ax, ay, bx, by int) int {
	return (utils.Abs(ax-bx) + 1) * (utils.Abs(ay-by) + 1)
}
