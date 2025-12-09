package day9

import (
	"fmt"
	"slices"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadStringLines(path[f])
	points := make([]point, len(data))
	areas := make([]pairs, 0, len(data)*(len(data)-1)/2)
	maxRect := -1

	for i, s := range data {
		x, y := parseNum(s)
		points[i] = point{x, y}
	}

	segments := buildsegments(points)

	for i := range data {
		for j := i + 1; j < len(data); j++ {
			ax, ay := parseNum(data[i])
			bx, by := parseNum(data[j])
			area := calcArea(ax, ay, bx, by)
			areas = append(areas, pairs{ax, ay, bx, by, area})
		}
	}

	slices.SortFunc(areas, func(i, j pairs) int {
		return j.area - i.area
	})

	for _, p := range areas {
		rectMin := point{min(p.ax, p.bx), min(p.ay, p.by)}
		rectMax := point{max(p.ax, p.bx), max(p.ay, p.by)}

		if isValid(rectMin, rectMax, segments) {
			maxRect = p.area
			break
		}
	}

	fmt.Println(maxRect)
}
