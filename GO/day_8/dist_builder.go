package day8

import (
	"runtime"
	"slices"
	"sync"
)

func mergeTwo(a, b []pairs) []pairs {
	if len(a) == 0 {
		return b
	}
	if len(b) == 0 {
		return a
	}

	out := make([]pairs, 0, len(a)+len(b))

	i, j := 0, 0
	for i < len(a) && j < len(b) {
		if a[i].dist <= b[j].dist {
			out = append(out, a[i])
			i++
		} else {
			out = append(out, b[j])
			j++
		}
	}

	if i < len(a) {
		out = append(out, a[i:]...)
	}

	if j < len(b) {
		out = append(out, b[j:]...)
	}

	return out
}

func mergeAll(allLists [][]pairs) []pairs {
	if len(allLists) == 0 {
		return nil
	}

	if len(allLists) == 1 {
		return allLists[0]
	}

	limit := runtime.GOMAXPROCS(0)
	sem := make(chan struct{}, limit)

	for len(allLists) > 1 {
		nextSize := (len(allLists) + 1) / 2
		next := make([][]pairs, nextSize)

		var wg sync.WaitGroup
		for i := 0; i < len(allLists); i += 2 {
			left := allLists[i]
			var right []pairs
			if i+1 < len(allLists) {
				right = allLists[i+1]
			} else {
				next[i/2] = left
				continue
			}

			wg.Add(1)
			sem <- struct{}{}
			go func(idx int, l, r []pairs) {
				defer wg.Done()
				merged := mergeTwo(l, r)
				next[idx] = merged
				<-sem
			}(i/2, left, right)
		}

		wg.Wait()
		allLists = next
	}

	return allLists[0]
}

func calcDistances(data []string) []pairs {
	var wg sync.WaitGroup
	var cwg sync.WaitGroup
	resultCh := make(chan []pairs, len(data)-1)
	allLists := make([][]pairs, 0, len(data)-1)

	cwg.Add(1)
	go func() {
		defer cwg.Done()
		for list := range resultCh {
			if len(list) > 0 {
				allLists = append(allLists, list)
			}
		}
	}()

	for i := range len(data) - 1 {
		wg.Add(1)
		go func(i int, line string) {
			defer wg.Done()
			a1, b1, c1 := parseNums(line)
			localDists := make([]pairs, 0, len(data)-i-1)

			for j := i + 1; j < len(data); j++ {
				a2, b2, c2 := parseNums(data[j])
				dist := eucDist(a1, a2, b1, b2, c1, c2)
				if dist > MAGICNUMDAY8 {
					continue
				}
				localDists = append(localDists, pairs{
					idx:      i,
					tarIdx:   j,
					idxLX:    a1,
					tarIdxLX: a2,
					dist:     dist,
				})
			}

			slices.SortFunc(localDists, func(x, y pairs) int {
				return x.dist - y.dist
			})

			resultCh <- localDists
		}(i, data[i])
	}

	go func() {
		wg.Wait()
		close(resultCh)
	}()

	cwg.Wait()

	if len(allLists) == 0 {
		return nil
	}

	dists := mergeAll(allLists)

	return dists
}
