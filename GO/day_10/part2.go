package day10

import (
	"fmt"
	"math"
	"math/big"
	"slices"
	"strconv"
	"strings"
	"sync"

	"aoc2025/utils"
)

type State struct {
	i   int
	v   []int
	sum int
}

func Part2(f string) {
	data := utils.ReadStringLines(path[f])

	var wg sync.WaitGroup
	var cwg sync.WaitGroup
	resChan := make(chan int, 10)

	sum := 0

	cwg.Add(1)
	go func() {
		defer cwg.Done()

		for n := range resChan {
			sum += n
		}
	}()

	for _, line := range data {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			btns, target := parseLine2(l)
			minPressJoltsLinearSolve(target, btns, resChan)
		}(line)
	}

	wg.Wait()
	close(resChan)
	cwg.Wait()

	fmt.Println(sum)
}

func parseJolts(s string) []int {
	s = strings.Trim(strings.TrimSpace(s), "{}")
	parts := strings.Split(s, ",")
	res := make([]int, len(parts))

	for i, p := range parts {
		n, _ := strconv.Atoi(strings.TrimSpace(p))
		res[i] = n
	}

	return res
}

func parseBtn2(s string) [][]int {
	parts := strings.Fields(s)
	btns := make([][]int, len(parts))

	for i, part := range parts {
		part = strings.Trim(strings.TrimSpace(part), "()")

		indices := strings.Split(part, ",")

		for _, idx := range indices {
			n := utils.FastAtoi(idx)
			btns[i] = append(btns[i], n)
		}
	}

	return btns
}

func parseLine2(s string) ([][]int, []int) {
	idxStart := strings.Index(s, " ")
	idxEnd := strings.LastIndex(s, " ")
	btns := s[idxStart:idxEnd]
	jolts := s[idxEnd+1:]

	return parseBtn2(btns), parseJolts(jolts)
}

func rrefRat(aug [][]*big.Rat) ([][]*big.Rat, []int) {
	m := len(aug)
	if m == 0 {
		return aug, nil
	}

	n := len(aug[0]) - 1
	pivots := []int{}
	row := 0
	col := 0

	zero := big.NewRat(0, 1)

	for row < m && col < n {
		pivot := -1

		for i := row; i < m; i++ {
			if aug[i][col].Sign() != 0 {
				pivot = i
				break
			}
		}

		if pivot == -1 {
			col++
			continue
		}

		if pivot != row {
			aug[row], aug[pivot] = aug[pivot], aug[row]
		}

		pv := aug[row][col]

		for j := col; j <= n; j++ {
			aug[row][j] = new(big.Rat).Quo(aug[row][j], pv)
		}

		for i := range m {
			if i == row {
				continue
			}

			factor := aug[i][col]

			if factor.Sign() == 0 {
				continue
			}

			for j := col; j <= n; j++ {
				tmp := new(big.Rat).Mul(factor, aug[row][j])
				aug[i][j] = new(big.Rat).Sub(aug[i][j], tmp)
			}
		}

		pivots = append(pivots, col)
		row++
		col++
	}

	for i := range m {
		for j := 0; j <= n; j++ {
			if aug[i][j].Sign() == 0 {
				aug[i][j] = zero
			}
		}
	}

	return aug, pivots
}

func inconsistentRowRat(row []*big.Rat) bool {
	n := len(row) - 1

	for i := range n {
		if row[i].Sign() != 0 {
			return false
		}
	}

	return row[n].Sign() != 0
}

func convertRREFToFloat(rref [][]*big.Rat) [][]float64 {
	m := len(rref)
	if m == 0 {
		return nil
	}

	n := len(rref[0])
	out := make([][]float64, m)

	for i := range m {
		out[i] = make([]float64, n)

		for j := range n {
			if rref[i][j] == nil {
				out[i][j] = 0.0
				continue
			}

			f := new(big.Float).SetRat(rref[i][j])
			v, _ := f.Float64()
			out[i][j] = v
		}
	}

	return out
}

func solvePivotVarWithFree(rref [][]float64, numFree int, coeff func(int, int) float64,
	fvIndex map[int]int, pivotRow map[int]int, v []int, j int,
) int {
	if idx, ok := fvIndex[j]; ok {
		return v[idx]
	}

	pi := pivotRow[j]
	sum := rref[pi][len(rref[pi])-1]

	for k := range numFree {
		sum -= coeff(pi, k) * float64(v[k])
	}

	return int(math.Round(sum))
}

func partialPivotValue(rref [][]float64, coeff func(int, int) float64, v []int, idx int, pi int) float64 {
	sum := rref[pi][len(rref[pi])-1]

	for k := range idx {
		sum -= coeff(pi, k) * float64(v[k])
	}

	return sum
}

func canBecomeNonNeg(rref [][]float64, coeff func(int, int) float64,
	numFree int, v []int, idx int, pi int,
) bool {
	pv := partialPivotValue(rref, coeff, v, idx, pi)

	if pv >= -1e-12 {
		return true
	}

	for k := idx; k < numFree; k++ {
		if coeff(pi, k) < 0 {
			return true
		}
	}

	return false
}

func infeasible(rref [][]float64, numFree int,
	coeff func(int, int) float64, numPivots int, idx int, v []int,
) bool {
	for pi := range numPivots {
		if !canBecomeNonNeg(rref, coeff, numFree, v, idx, pi) {
			return true
		}
	}

	return false
}

func branchAndBoundFloat(rref [][]float64, pivots []int, n int, target []int, buttons [][]int) int {
	pivotSet := make(map[int]bool)

	for _, p := range pivots {
		pivotSet[p] = true
	}

	freeVars := []int{}

	for i := range n {
		if !pivotSet[i] {
			freeVars = append(freeVars, i)
		}
	}

	numFree := len(freeVars)
	fvIndex := map[int]int{}

	for i, fv := range freeVars {
		fvIndex[fv] = i
	}

	pivotRow := map[int]int{}

	for i, p := range pivots {
		pivotRow[p] = i
	}

	coeff := func(pi, k int) float64 {
		return rref[pi][freeVars[k]]
	}

	maxJ := 0
	sumJ := 0

	for _, v := range target {
		if v > maxJ {
			maxJ = v
		}

		sumJ += v
	}

	best := math.MaxInt
	stack := []State{{0, make([]int, numFree), 0}}

	for len(stack) > 0 {
		s := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if s.sum >= best || s.sum > sumJ {
			continue
		}

		if infeasible(rref, numFree, coeff, len(pivots), s.i, s.v) {
			continue
		}

		if s.i == numFree {
			sol := make([]int, n)
			valid := true
			for j := range n {
				sol[j] = solvePivotVarWithFree(rref, numFree, coeff, fvIndex, pivotRow, s.v, j)

				if sol[j] < 0 {
					valid = false
					break
				}
			}
			if valid && satisfiesTarget(sol, buttons, target) {
				total := 0

				for _, x := range sol {
					total += x
				}

				if total < best {
					best = total
				}
			}

			continue
		}

		maxVal := maxJ
		rem := sumJ - s.sum

		if rem < maxVal {
			maxVal = rem
		}

		for x := maxVal; x >= 0; x-- {
			nv := make([]int, len(s.v))
			copy(nv, s.v)
			nv[s.i] = x
			stack = append(stack, State{s.i + 1, nv, s.sum + x})
		}
	}

	if best == math.MaxInt32 {
		return -1
	}

	return best
}

func minPressJoltsLinearSolve(target []int, buttons [][]int, res chan int) {
	m := len(target)
	n := len(buttons)

	augRat := make([][]*big.Rat, m)

	for i := range m {
		row := make([]*big.Rat, n+1)

		for j := range n {
			row[j] = big.NewRat(0, 1)

			if slices.Contains(buttons[j], i) {
				row[j] = big.NewRat(1, 1)
			}
		}

		row[n] = big.NewRat(int64(target[i]), 1)
		augRat[i] = row
	}

	rrefRatMat, pivots := rrefRat(augRat)

	for i := len(pivots); i < len(rrefRatMat); i++ {
		if inconsistentRowRat(rrefRatMat[i]) {
			res <- -1
			return
		}
	}

	rrefFloat := convertRREFToFloat(rrefRatMat)
	ans := branchAndBoundFloat(rrefFloat, pivots, n, target, buttons)

	res <- ans
}

func satisfiesTarget(sol []int, buttons [][]int, target []int) bool {
	counts := make([]int, len(target))

	for btn, presses := range sol {
		if presses < 0 {
			return false
		}

		for _, c := range buttons[btn] {
			counts[c] += presses
		}
	}

	for i := range target {
		if counts[i] != target[i] {
			return false
		}
	}

	return true
}
