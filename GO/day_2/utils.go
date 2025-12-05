package day2

import (
	"log"
	"strconv"
	"strings"
)

type limits struct{ a, b int }

func parseInput(data string) ([]limits, int, int) {
	var ranges []limits
	maxEnd := 0
	maxDigits := 1

	for _, p := range strings.Split(data, ",") {
		p = strings.TrimSpace(p)
		pe := strings.Split(p, "-")
		if len(pe) != 2 {
			continue
		}

		start, err1 := strconv.Atoi(strings.TrimSpace(pe[0]))
		end, err2 := strconv.Atoi(strings.TrimSpace(pe[1]))
		if err1 != nil || err2 != nil || start > end {
			log.Fatal("Conversion to int failed")
		}

		ranges = append(ranges, limits{start, end})
		if end > maxEnd {
			maxEnd = end
		}

		d := countDigits(end)
		if d > maxDigits {
			maxDigits = d
		}
	}

	return ranges, maxEnd, maxDigits
}

func countDigits(n int) int {
	if n == 0 {
		return 1
	}

	d := 0

	for n > 0 {
		d++
		n /= 10
	}

	return d
}

func buildPow10(max int) []int {
	pow10 := make([]int, max+1)
	pow10[0] = 1

	for i := 1; i <= max; i++ {
		pow10[i] = pow10[i-1] * 10
	}

	return pow10
}

func generateRepeatedIDs(ranges []limits, maxEnd int, maxDigits int, minR int, maxR int) int {
	pow10 := buildPow10(maxDigits)
	seen := map[int]struct{}{}
	sum := 0

	for L := 1; L <= maxDigits/2; L++ {
		base := pow10[L]
		pStart := pow10[L-1]
		pEnd := pow10[L] - 1

		for p := pStart; p <= pEnd; p++ {
			maxRep := maxDigits / L
			if maxRep < minR {
				continue
			}
			if maxRep > maxR {
				maxRep = maxR
			}

			for r := minR; r <= maxRep; r++ {
				N := 0
				overflow := false

				for rep := 0; rep < r; rep++ {
					if N > maxEnd/base && rep < r {
						overflow = true
						break
					}
					N = N*base + p
					if N > maxEnd {
						overflow = true
						break
					}
				}

				if overflow {
					continue
				}

				for _, R := range ranges {
					if N >= R.a && N <= R.b {
						if _, ok := seen[N]; !ok {
							seen[N] = struct{}{}
							sum += N
						}
						break
					}
				}
			}
		}
	}

	return sum
}
