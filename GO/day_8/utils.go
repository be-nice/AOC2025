package day8

import (
	"cmp"
	"math"
	"slices"
	"strconv"
	"strings"
)

type pairs struct {
	idx      int
	tarIdx   int
	idxLX    int
	tarIdxLX int
	dist     float64
}

func (p *pairs) buildSets(setMap *[]map[int]struct{}, tar int) int {
	n1 := -1
	n2 := -1
	for i, m := range *setMap {
		if _, ok := m[p.idx]; ok && n1 == -1 {
			n1 = i
		}
		if _, ok := m[p.tarIdx]; ok && n2 == -1 {
			n2 = i
		}
	}

	if n1 == -1 && n2 == -1 {
		newSet := make(map[int]struct{})
		newSet[p.idx] = struct{}{}
		newSet[p.tarIdx] = struct{}{}
		*setMap = append(*setMap, newSet)
	}

	if n1 == -1 && n2 != -1 {
		(*setMap)[n2][p.idx] = struct{}{}
	}

	if n2 == -1 && n1 != -1 {
		(*setMap)[n1][p.tarIdx] = struct{}{}
	}

	if n1 != n2 && n1 != -1 && n2 != -1 {
		m := (*setMap)[n2]

		for k := range m {
			(*setMap)[n1][k] = struct{}{}
		}

		*setMap = append((*setMap)[:n2], (*setMap)[n2+1:]...)
	}

	if len(*setMap) == 1 && len((*setMap)[0]) == tar {
		return p.idxLX * p.tarIdxLX
	}

	return -1
}

func parseNums(s string) (int, int, int) {
	nums := strings.Split(s, ",")
	for k, el := range nums {
		nums[k] = strings.TrimSpace(el)
	}

	n1, _ := strconv.Atoi(nums[0])
	n2, _ := strconv.Atoi(nums[1])
	n3, _ := strconv.Atoi(nums[2])

	return n1, n2, n3
}

func calcDistances(data []string) []pairs {
	dists := []pairs{}
	for i, line := range data[:len(data)-1] {
		a1, b1, c1 := parseNums(line)
		for j, tarLine := range data[i+1:] {
			a2, b2, c2 := parseNums(tarLine)

			dists = append(dists, pairs{i, j + 1 + i, a1, a2, eucDist(a1, a2, b1, b2, c1, c2)})
		}
	}

	slices.SortFunc(dists, func(a, b pairs) int {
		return cmp.Compare(a.dist, b.dist)
	})

	return dists
}

func eucDist(a1, a2, b1, b2, c1, c2 int) float64 {
	dx := float64(a1 - a2)
	dy := float64(b1 - b2)
	dz := float64(c1 - c2)
	return math.Hypot(math.Hypot(dx, dy), dz)
}
