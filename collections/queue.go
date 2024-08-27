package collections

import "sync"

type queue[T any] struct {
	items []T
	mutex sync.Mutex
}

func NewQueue[T any]() *queue[T] {
	return &queue[T]{items: make([]T, 0)}
}

func (q *queue[T]) Put(item T) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

func (q *queue[T]) HasNext() bool {
	return len(q.items) > 0
}

func (q *queue[T]) Get() (T, bool) {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	var zero T

	if len(q.items) == 0 {
		return zero, false
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item, true
}

func (q *queue[T]) Size() int {
	return len(q.items)
}

func (q *queue[T]) Look() (T, bool) {
	if q.HasNext() {
		return q.items[0], true
	}
	var zero T
	return zero, false
}
