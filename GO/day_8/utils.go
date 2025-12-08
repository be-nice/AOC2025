package day8

import (
	"strings"
)

const MAGICNUMDAY8 = 292_000_000

type pairs struct {
	idx      int
	tarIdx   int
	idxLX    int
	tarIdxLX int
	dist     int
}

func parseNums(s string) (int, int, int) {
	nums := strings.Split(s, ",")

	n1 := fastAtoi(strings.TrimSpace(nums[0]))
	n2 := fastAtoi(strings.TrimSpace(nums[1]))
	n3 := fastAtoi(strings.TrimSpace(nums[2]))

	return n1, n2, n3
}

func fastAtoi(s string) int {
	n := 0

	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}

	return n
}

func eucDist(a1, a2, b1, b2, c1, c2 int) int {
	dx := a1 - a2
	dy := b1 - b2
	dz := c1 - c2
	return dx*dx + dy*dy + dz*dz
}

func getTopN(n int, setMap []map[int]struct{}) []int {
	top := make([]int, n)

	for _, m := range setMap {
		switch {
		case len(m) > top[0]:
			top[2] = top[1]
			top[1] = top[0]
			top[0] = len(m)
		case len(m) > top[1]:
			top[2] = top[1]
			top[1] = len(m)
		case len(m) > top[2]:
			top[2] = len(m)
		}
	}

	return top
}

func (p *pairs) getSetIdx(setMap *[]map[int]struct{}) (int, int) {
	n1 := -1
	n2 := -1

	for i, m := range *setMap {
		if _, ok := m[p.idx]; ok {
			n1 = i
		}
		if _, ok := m[p.tarIdx]; ok {
			n2 = i
		}
	}

	return n1, n2
}
