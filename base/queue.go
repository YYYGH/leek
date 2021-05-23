package base

import "sync"

type Queue struct {
	data []interface{}
	size int
	sync.Mutex
}

type QueueFace interface {
	Push(v interface{})
	Pop()
	Empty() bool
	Size() int
}

func NewQueue() *Queue {
	return &Queue{
		data: make([]interface{}, 0),
		size: 0,
	}
}

func (q *Queue) Push(val interface{}) {
	q.Lock()
	defer q.Unlock()
	q.data = append(q.data, val)
	q.size++
}

func (q *Queue) Pop() (interface{}, bool) {
	q.Lock()
	defer q.Unlock()
	// 没有元素
	if q.size <= 0 {
		return nil, false
	}
	item := q.data[0]
	q.data[0] = nil
	q.data = q.data[1:]
	q.size--
	return item, true
}

func (q *Queue) Empty() bool {
	if q.size <= 0 {
		return true
	}
	return false
}

func (q *Queue) Size() int {
	return q.size + 1
}
