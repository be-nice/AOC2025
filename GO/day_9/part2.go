package day9

import (
	"fmt"
	"math/rand/v2"

	"aoc2025/utils"
)

func Part2(f string) {
	data := utils.ReadStringLines(path[f])
	maxRect := -1

	points := createCorners(data)
	edges := createEdges(points)

	rand.Shuffle(len(points), func(i, j int) {
		points[i], points[j] = points[j], points[i]
	})

	for i := range len(points) - 1 {
	hot_path:
		for j := i + 1; j < len(points); j++ {
			a := points[i]
			b := points[j]
			area := calcArea(a.x, a.y, b.x, b.y)

			if area < maxRect {
				continue
			}

			recMinX := min(a.x, b.x) + 1
			recMaxX := max(a.x, b.x) - 1
			recMinY := min(a.y, b.y) + 1
			recMaxY := max(a.y, b.y) - 1

			for _, edge := range edges {
				if edge.isIntersect(recMinX, recMaxX, recMinY, recMaxY) {
					continue hot_path
				}
			}

			maxRect = area
		}
	}

	fmt.Println(maxRect)
}

func createCorners(data []string) []point {
	points := make([]point, len(data))

	for i, s := range data {
		x, y := parseNum(s)
		points[i] = point{x, y}
	}

	return points
}

func createEdges(points []point) []segment {
	segments := make([]segment, len(points)+1)

	for i := range len(points) - 1 {
		segments[i] = segment{points[i], points[i+1]}
	}
	segments[len(segments)-1] = segment{points[len(points)-1], points[0]}

	return segments
}
