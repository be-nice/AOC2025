package day9

import (
	"strings"

	"aoc2025/utils"
)

type point struct {
	x int
	y int
}

type segment struct {
	start point
	end   point
}

func (s *segment) isIntersect(a, b point) bool {
	recMinX := min(a.x, b.x) + 1
	recMaxX := max(a.x, b.x) - 1
	recMinY := min(a.y, b.y) + 1
	recMaxY := max(a.y, b.y) - 1

	edgeMinX := min(s.start.x, s.end.x)
	edgeMaxX := max(s.start.x, s.end.x)
	edgeMinY := min(s.start.y, s.end.y)
	edgeMaxY := max(s.start.y, s.end.y)

	if edgeMaxX < recMinX || edgeMinX > recMaxX {
		return false
	}
	if edgeMaxY < recMinY || edgeMinY > recMaxY {
		return false
	}

	return true
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
