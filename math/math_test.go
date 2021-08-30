// Code generated by go2go; DO NOT EDIT.


//line math_test.go2:1
package math

//line math_test.go2:1
import (
//line math_test.go2:1
 "math"
//line math_test.go2:1
 "testing"
//line math_test.go2:1
)

//line math_test.go2:7
func TestMaxFromNumbers(t *testing.T) {
	got := instantiate୦୦MaxFromNumbers୦int(1, 2, 3, 5, 3)
	want := 5

	if got != want {
		t.Errorf("got %q not equal want %q", got, want)
	}
}

func TestMax(t *testing.T) {
	got := instantiate୦୦Max୦int(1, 2)
	want := 2

	if got != want {
		t.Errorf("got %q not equal want %q", got, want)
	}
}

func TestAbs(t *testing.T) {
	got := instantiate୦୦Abs୦int(-4)
	want := 4

	if got != want {
		t.Errorf("got %q not equal want %q", got, want)
	}
}
//line math.go2:21
func instantiate୦୦MaxFromNumbers୦int(vals ...int,) int {
	return instantiate୦୦max୦int(vals...)
}

func instantiate୦୦Max୦int(a, b int,) int {
	if a > b {
		return a
	}
	return b
}

func instantiate୦୦Abs୦int(a int,) int {
	if a > 0 {
		return a
	}
	return -a
}
//line math.go2:11
func instantiate୦୦max୦int(vals ...int,) int {
	maxv := int(math.Inf(-1))
	for _, v := range vals {
		if v > maxv {
			maxv = v
		}
	}
	return maxv
}

//line math.go2:19
var _ = math.Abs
//line math.go2:19
var _ = testing.AllocsPerRun