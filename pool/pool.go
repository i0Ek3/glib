package pool

// Pool defines a goroutine pool, which contains
// task channel and task number channel
type Pool struct {
	work chan func()
	todo chan struct{}
}

// New creates a goroutine pool with given size
func New(size int) *Pool {
	return &Pool{
		work: make(chan func()),
		todo: make(chan struct{}, size),
	}
}

// NewTask adds new task to work, if work is busy
// then create a new goroutine to call worker and run the task
func (p *Pool) NewTask(task func()) {
	select {
	case p.work <- task:
	case p.todo <- struct{}{}:
		go p.worker(task)
	}
}

// worker runs task
func (p *Pool) worker(task func()) {
	defer func() {
		<-p.todo
	}()

	for {
		task()
		task = <-p.work
	}
}
