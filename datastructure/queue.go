package main
import (
	"sync"
	"github.com/cheekybits/genny/generic"
)

type Item generic.Type

type ItemQueue struct {
	items []Item
	lock sync.RWMutex
}

func (s *ItemQueue) New() *ItemQueue{
	s.items = []Item{}
	return s
}

func (s *ItemQueue) Enqueue(t Item){
	s.lock.Lock()
	s.items = append(s.items, t)
	s.lock.Unlock()
}

func (s *ItemQueue) Dequeue() *Item{
	s.lock.Lock()
	item := s.items[0]
	s.items = s.items[1:len(s.items)]
	s.lock.Unlock()
	return &item
}

func (s *ItemQueue) Front() *Item{
	s.lock.RLock()
	item := s.items[0]
	s.lock.RUnlock()
	return &item
}

func (s *ItemQueue) IsEmpty() bool{
	return len(s.items) == 0
}

func (s *ItemQueue) Size() int{
	return len(s.items)
}