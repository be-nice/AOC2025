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

func countEdge(data [][]byte, pos pos, p2 bool) bool {
	count := 0
	for _, el := range dirs {
		if pos.y+el.y < 0 || pos.y+el.y > len(data)-1 {
			continue
		}
		if pos.x+el.x < 0 || pos.x+el.x > len(data[0])-1 {
			continue
		}

		if data[pos.y+el.y][pos.x+el.x] == '@' {
			count++
		}

		if count >= 4 {
			return false
		}
	}

	if p2 {
		data[pos.y][pos.x] = '.'
	}

	return true
}
