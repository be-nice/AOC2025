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
