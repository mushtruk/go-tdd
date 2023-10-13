package sync_counter

import "sync"

type Counter struct {
	value int
	mutex sync.Mutex
}

func NewCounter() *Counter {
	return &Counter{}
}

func (counter *Counter) Increment() {
	counter.mutex.Lock()
	defer counter.mutex.Unlock()
	counter.value++
}

func (counter *Counter) Value() int {
	return counter.value
}
