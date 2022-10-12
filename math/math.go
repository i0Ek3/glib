package math

import "math"

type Int interface {
	~int | ~uint
}

func Tmax[T Int](vals []T) T {
	maxv := T(math.Inf(-1))
	for _, v := range vals {
		if v > maxv {
			maxv = v
		}
	}

	return maxv
}

func Tmin[T Int](vals []T) T {
	minv := Tmax(vals)
	for _, v := range vals {
		if v < minv {
			minv = v
		}
	}

	return minv
}

func Tabs[T Int](vals []T) []T {
	t := []T{}
	for _, v := range vals {
		if v < 0 {
			t = append(t, -v)
			continue
		}
		t = append(t, v)
	}

	return t
}

// Ttrop equals if cond ? expr1 : expr2
func Ttrop[T any](cond bool, expr1, expr2 T) T {
	if cond {
		return expr1
	}

	return expr2
}
