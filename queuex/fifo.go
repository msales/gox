package queuex

import "sync"

// FIFO is simple implementation of FIFO (first-in-first-out) queue with locks.
type FIFO[T any] struct {
	m sync.Mutex

	vals []T
}

// NewFIFO initializes queue with values.
func NewFIFO[T any](vals ...T) *FIFO[T] {
	return &FIFO[T]{vals: vals}
}

// Push adds elements at the end of the queue.
func (f *FIFO[T]) Push(v T) *FIFO[T] {
	f.m.Lock()
	defer f.m.Unlock()

	f.vals = append(f.vals, v)
	return f
}

// Pop returns first element from the queue.
// Bool is also returned to indicate if queue has run out of the values.
func (f *FIFO[T]) Pop() (T, bool) {
	f.m.Lock()
	defer f.m.Unlock()

	if len(f.vals) == 0 {
		var t T
		return t, false
	}

	v := f.vals[0]
	f.vals = f.vals[1:]
	return v, true
}

// Len returns length of queue.
func (f *FIFO[T]) Len() int {
	f.m.Lock()
	defer f.m.Unlock()

	return len(f.vals)
}
