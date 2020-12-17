package calculator

import (
	"sync"
)

type Stack struct {
	items []interface{}
	lock  sync.RWMutex
}

func NewStack() *Stack {
	return &Stack{
		items: []interface{}{},
	}
}

func (s *Stack) Push(t interface{}) {
	s.lock.Lock()
	defer s.lock.Unlock()
	s.items = append(s.items, t)
}

func (s *Stack) Pop() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return nil
	} else {
		item := s.items[len(s.items)-1]
		s.items = s.items[0 : len(s.items)-1]
		return item
	}

}

func (s *Stack) Peek() interface{} {
	s.lock.Lock()
	defer s.lock.Unlock()

	if len(s.items) == 0 {
		return nil
	} else {
		item := s.items[len(s.items)-1]
		return item
	}

}


