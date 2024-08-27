package collections

import "sync"

type anyQueue struct {
	items []any
	mutex sync.Mutex
}

func NewAnyQueue() *anyQueue {
	return &anyQueue{items: make([]any, 0)}
}

func (q *anyQueue) Put(item any) {
	q.mutex.Lock()
	defer q.mutex.Unlock()
	q.items = append(q.items, item)
}

func (q *anyQueue) Get() any {
	q.mutex.Lock()
	defer q.mutex.Unlock()

	if len(q.items) == 0 {
		return nil
	}

	item := q.items[0]
	q.items = q.items[1:]
	return item
}

func (q *anyQueue) Size() int {
	return len(q.items)
}

func (q *anyQueue) HasNext() bool {
	return len(q.items) > 0
}

func (q *anyQueue) Look() any {
	if len(q.items) > 0 {
		return q.items[0]
	}
	return nil
}
