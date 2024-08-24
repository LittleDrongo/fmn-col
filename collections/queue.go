package collections

import "sync"

type Queue struct {
	items []any
	mutex sync.Mutex
}

func NewQueue() *Queue {
	return &Queue{items: make([]any, 0)}
}

func (q *Queue) Put(item any) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

func (q *Queue) Get() any {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *Queue) Size() int {
	return len(q.items)
}
