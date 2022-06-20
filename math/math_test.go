package math

import (
	"reflect"
	"testing"
)

func TestTmax(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 5}, 5},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, 1},
		{[]int{-1, -2, -3}, -1},
	}

	for _, tt := range tests {
		if got := Tmax(tt.nums); got != tt.want {
			t.Errorf("Tmax() = %v, want %v", got, tt.want)
		}
	}
}

func FuzzTmax(f *testing.F) {
	testcases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 5}, 5},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, 1},
		{[]int{-1, -2, -3}, -1},
	}
	for _, tc := range testcases {
		f.Add(tc.want)
	}
	f.Fuzz(func(t *testing.T, _ int) {
		for _, tc := range testcases {
			if got := Tmax(tc.nums); got != tc.want {
				t.Errorf("Tmax() = %v, want %v", got, tc.want)
			}
		}
	})
}

func TestTmin(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 5}, 1},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, -1},
		{[]int{-1, -2, -3}, -3},
	}

	for _, tt := range tests {
		if got := Tmin(tt.nums); got != tt.want {
			t.Errorf("Tmin() = %v, want %v", got, tt.want)
		}
	}
}

func FuzzTmin(f *testing.F) {
	testcases := []struct {
		nums []int
		want int
	}{
		{[]int{1, 2, 3, 5}, 1},
		{[]int{0, 0, 0}, 0},
		{[]int{-1, 0, 1}, -1},
		{[]int{-1, -2, -3}, -3},
	}
	for _, tc := range testcases {
		f.Add(tc.want)
	}
	f.Fuzz(func(t *testing.T, _ int) {
		for _, tc := range testcases {
			if got := Tmin(tc.nums); got != tc.want {
				t.Errorf("Tmin() = %v, want %v", got, tc.want)
			}
		}
	})
}

func TestTabs(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 5}, []int{1, 2, 3, 5}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{-1, 0, 1}, []int{1, 0, 1}},
		{[]int{-1, -2, -3}, []int{1, 2, 3}},
	}

	sum := 0
	for _, tt := range tests {
		got := Tabs(tt.nums)
		want := tt.want

		for _, v1 := range got {
			sum += v1
		}
		for _, v2 := range want {
			sum -= v2
		}

		if sum != 0 {
			t.Errorf("Tabs() = %v, want %v", got, tt.want)
		}
	}
}

func FuzzTabs(f *testing.F) {
	testcases := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 5}, []int{1, 2, 3, 5}},
		{[]int{0, 0, 0}, []int{0, 0, 0}},
		{[]int{-1, 0, 1}, []int{1, 0, 1}},
		{[]int{-1, -2, -3}, []int{1, 2, 3}},
	}
	for _, tc := range testcases {
		sum := 0
		for _, v := range tc.nums {
			sum += v
		}
		for _, v := range tc.want {
			sum -= v
		}
		f.Add(sum)
	}
	f.Fuzz(func(t *testing.T, _ int) {
		for _, tc := range testcases {
			got := Tabs(tc.nums)
			want := Tabs(tc.want)
			if !reflect.DeepEqual(got, want) {
				t.Errorf("Tabs() = %v, want %v", got, want)
			}
		}
	})
}

var (
	tests = [][]int{[]int{1, 2, 3, 5}, []int{0, 0, 0}, []int{-1, 0, 1}, []int{-1, -2, -3}}
)

func BenchmarkTmax(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			_ = Tmax(tt)
		}
	}
}

func BenchmarkTmin(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			_ = Tmin(tt)
		}
	}
}

func BenchmarkTabs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tt := range tests {
			_ = Tabs(tt)
		}
	}
}
