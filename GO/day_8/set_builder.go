package day8

func (p *pairs) buildSets(setMap *[]map[int]struct{}, all int) int {
	isAllMerged := func() int {
		if len(*setMap) == 1 && len((*setMap)[0]) == all {
			return p.idxLX * p.tarIdxLX
		}

		return -1
	}

	idx1, idx2 := p.getSetIdx(setMap)

	switch {
	case idx1 == -1 && idx2 == -1:
		newSet := make(map[int]struct{})
		newSet[p.idx] = struct{}{}
		newSet[p.tarIdx] = struct{}{}
		*setMap = append(*setMap, newSet)
	case idx1 == -1:
		(*setMap)[idx2][p.idx] = struct{}{}
	case idx2 == -1:
		(*setMap)[idx1][p.tarIdx] = struct{}{}
	case idx1 != idx2:
		m := (*setMap)[idx2]

		for k := range m {
			(*setMap)[idx1][k] = struct{}{}
		}

		*setMap = append((*setMap)[:idx2], (*setMap)[idx2+1:]...)
	}

	return isAllMerged()
}
