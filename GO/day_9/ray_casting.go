package day9

func buildsegments(points []point) []segment {
	n := len(points)
	segments := make([]segment, n)

	for i := range n {
		p1 := points[i]
		p2 := points[(i+1)%n]
		start := point{min(p1.x, p2.x), min(p1.y, p2.y)}
		end := point{max(p1.x, p2.x), max(p1.y, p2.y)}
		segments[i] = segment{
			start:        start,
			end:          end,
			isHorizontal: p1.y == p2.y,
		}
	}

	return segments
}

func isValid(rectMin, rectMax point, segments []segment) bool {
	corners := []point{
		rectMin,
		{rectMax.x, rectMin.y},
		rectMax,
		{rectMin.x, rectMax.y},
	}

	for _, c := range corners {
		if !isPointInPoly(c, segments) {
			return false
		}
	}

	edges := []segment{
		{corners[0], corners[1], true},
		{corners[1], corners[2], false},
		{corners[3], corners[2], true},
		{corners[0], corners[3], false},
	}

	for _, e := range edges {
		for _, s := range segments {
			if isInternalIntersect(e, s) {
				return false
			}
		}
	}

	return true
}

func isPointInPoly(p point, segments []segment) bool {
	crossings := 0

	for _, seg := range segments {
		if seg.isHorizontal && p.y == seg.start.y && p.x >= seg.start.x && p.x <= seg.end.x {
			return true
		}
		if !seg.isHorizontal && p.x == seg.start.x && p.y >= seg.start.y && p.y <= seg.end.y {
			return true
		}

		if !seg.isHorizontal {
			x := seg.start.x
			y1, y2 := seg.start.y, seg.end.y
			if x > p.x && p.y >= y1 && p.y < y2 {
				crossings++
			}
		}
	}

	return crossings%2 == 1
}

func isInternalIntersect(rectEdge, polyEdge segment) bool {
	if rectEdge.isHorizontal && !polyEdge.isHorizontal {
		rx1, rx2 := rectEdge.start.x, rectEdge.end.x
		ry := rectEdge.start.y
		px := polyEdge.start.x
		py1, py2 := polyEdge.start.y, polyEdge.end.y

		if px > rx1 && px < rx2 && ry > py1 && ry < py2 {
			return true
		}
	} else if !rectEdge.isHorizontal && polyEdge.isHorizontal {
		rx := rectEdge.start.x
		ry1, ry2 := rectEdge.start.y, rectEdge.end.y
		py := polyEdge.start.y
		px1, px2 := polyEdge.start.x, polyEdge.end.x

		if rx > px1 && rx < px2 && py > ry1 && py < ry2 {
			return true
		}
	}

	return false
}
