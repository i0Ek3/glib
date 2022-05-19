package singleton

import (
	"sync"
)

type PtrMutexInstance struct {
	sync.Mutex
	instanceP *int
}

func (in *PtrMutexInstance) ReturnPtrMutexInstance() *int {
	in.Lock()
	if in.instanceP == nil {
		i := 1
		in.instanceP = &i
	}
	in.Unlock()
	return in.instanceP
}

type IntMutexInstance struct {
	sync.Mutex
	instanceI int
}

func (in *IntMutexInstance) ReturnIntMutexInstance() int {
	in.Lock()
	if in.instanceI == 0 {
		i := 1
		in.instanceI = i
	}
	in.Unlock()
	return in.instanceI
}

type PtrOnceInstance struct {
	sync.Once
	instanceP *int
}

func (in *PtrOnceInstance) ReturnPtrOnceInstance() *int {
	in.Do(func() {
		if in.instanceP == nil {
			i := 1
			in.instanceP = &i
		}
	})
	return in.instanceP
}

type IntegerInstance struct {
	sync.Once
	instanceI int
}

func (in *IntegerInstance) ReturnIntOnceInstance() int {
	in.Do(func() {
		if in.instanceI == 0 {
			i := 1
			in.instanceI = i
		}
	})
	return in.instanceI
}
