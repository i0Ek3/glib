// Package union offers a union find set to Find/Merge integers with given number.
package union

type Union struct {
	size   int
	father []int
	// rank optimization
	rank []int
}

// New creates an Union instance
func New(n int) *Union {
	fa := make([]int, n)
	rk := make([]int, n)
	return &Union{
		size:   n,
		father: fa,
		rank:   rk,
	}
}

// Init initializes an Union instance
func (u *Union) Init(n int) {
	for i := 1; i <= n; i++ {
		u.father[i] = i
		u.rank[i] = 1
	}
}

func (u *Union) Find(x int) int {
	if u.father[x] == x {
		return u.father[x]
	}
	return u.Find(u.father[x])
}

func (u *Union) Merge(from, to int) {
	x, y := u.Find(from), u.Find(to)
	if u.rank[x] <= u.rank[y] {
		u.father[x] = y
	} else {
		u.father[y] = x
	}
	if u.rank[x] == u.rank[y] && x != y {
		u.rank[y]++
	}
}

func (u *Union) Same(x, y int) bool {
	return u.Find(x) == u.Find(y)
}
