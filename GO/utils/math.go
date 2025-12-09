package utils

type Number interface {
	int | float64
}

func SumList[T Number](n []T) T {
	var res T

	for _, el := range n {
		res += el
	}

	return res
}

func MultList[T Number](n []T) T {
	var res T = 1

	for _, el := range n {
		res *= el
	}

	return res
}

func FastAtoi(s string) int {
	n := 0

	for _, ch := range s {
		n = n*10 + int(ch-'0')
	}

	return n
}

func Abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
