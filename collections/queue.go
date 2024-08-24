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

func (q *queue[T]) Get() *T {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return &item
}
