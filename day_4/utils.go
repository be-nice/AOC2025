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
		if pos.y+d.y < 0 || pos.y+d.y > len(data)-1 {
			continue
		}

		if pos.x+d.x < 0 || pos.x+d.x > len(data[0])-1 {
			continue
		}

		if data[pos.y+d.y][pos.x+d.x] == '@' {
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
