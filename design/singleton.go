package design

import (
	"sync"
)

type Instance struct {
	sync.Once
	instanceP *int
	instanceI int
}

func (in *Instance) ReturnPointerInstance() *int {
	in.Do(func() {
		if in.instanceP == nil {
			i := 1
			in.instanceP = &i
		}
	})
	return in.instanceP
}

func (in *Instance) ReturnIntInstance() int {
	in.Do(func() {
		if in.instanceI == 0 {
			i := 1
			in.instanceI = i
		}
	})
	return in.instanceI
}
