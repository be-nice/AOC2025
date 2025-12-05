package day4

type pos struct {
	y int
	x int
}

var dirs = []pos{
	{-1, -1},
	{0, -1},
	{1, -1},
	{1, 0},
	{1, 1},
	{0, 1},
	{-1, 1},
	{-1, 0},
}

var (
	toDelete = make([]pos, 0, 3000)
	countMap = make(map[pos]int)
)

func countEdge(data [][]byte, pos pos, p1 bool) (moved bool) {
	count := 0

	for _, d := range dirs {
		dx, dy := pos.x+d.x, pos.y+d.y

		if dy < 0 || dy > len(data)-1 {
			continue
		}

		if dx < 0 || dx > len(data[0])-1 {
			continue
		}

		if data[dy][dx] == '@' {
			count++
		}

		// part 1 early return
		if p1 && count >= 4 {
			return
		}
	}

	// part 1 early return
	if p1 {
		return true
	}

	// part 2 extra logic     i
	if count < 4 {
		toDelete = append(toDelete, pos)
		return
	}

	countMap[pos] = count

	return
}
