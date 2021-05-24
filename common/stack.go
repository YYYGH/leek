package common

import "sync"

type Stack struct {
	data      []interface{}
	cap       int
	lastIndex int
	sync.Mutex
}

type StackFace interface {
	Put(val interface{})
	Get() interface{}
	Empty() bool
	Size() int
}

func NewStack() *Stack {
	return &Stack{
		data:      make([]interface{}, 0),
		cap:       0,
		lastIndex: -1,
	}
}

func (s *Stack) Put(val interface{}) {
	s.Lock()
	defer s.Unlock()
	if s.cap > s.lastIndex+1 {
		s.data[s.lastIndex+1] = val
	} else {
		s.data = append(s.data, val)
	}
	s.lastIndex++
	s.cap++
}

func (s *Stack) Get() (interface{}, bool) {
	s.Lock()
	defer s.Unlock()
	// 没有元素
	if s.lastIndex < 0 {
		return nil, false
	}
	item := s.data[s.lastIndex]
	s.data[s.lastIndex] = nil
	s.lastIndex--
	return item, true
}

func (s *Stack) Empty() bool {
	if s.lastIndex < 0 {
		return true
	}
	return false
}

func (s *Stack) Size() int {
	return s.lastIndex + 1
}
