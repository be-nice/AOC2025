package day10

import (
	"fmt"
	"strconv"
	"strings"
	"sync"

	"aoc2025/utils"
)

func Part1(f string) {
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
		go func() {
			defer wg.Done()
			state, btns := parseLine(line)
			minPresses(state, btns, resChan)
		}()
	}

	wg.Wait()
	close(resChan)
	cwg.Wait()

	fmt.Println(sum)
}

func parseLine(s string) (uint16, []uint16) {
	state, rest, _ := strings.Cut(s, " ")
	idx := strings.LastIndex(rest, " ")
	btns := rest[:idx]

	return parseTarget(state), parseBtn(btns)
}

func parseTarget(s string) uint16 {
	s = strings.Trim(strings.TrimSpace(s), "[]")
	tar := 0

	for i, el := range s {
		if el == '#' {
			tar |= 1 << i
		}
	}

	return uint16(tar)
}

func parseBtn(s string) []uint16 {
	parts := strings.Fields(s)

	btns := make([]uint16, len(parts))

	for i, part := range parts {
		part = strings.Trim(part, "()")
		var mask uint16 = 0
		for _, idx := range strings.Split(part, ",") {
			n, _ := strconv.Atoi(idx)
			mask |= 1 << n
		}
		btns[i] = mask
	}

	return btns
}

func minPresses(target uint16, buttons []uint16, res chan int) {
	n := len(buttons)
	best := -1

	for mask := range 1 << n {
		state := uint16(0)
		presses := 0

		for i := range n {
			if mask&(1<<i) != 0 {
				state ^= buttons[i]
				presses++
			}
		}

		if state == target {
			if best == -1 || presses < best {
				best = presses
			}
		}
	}

	res <- best
}
