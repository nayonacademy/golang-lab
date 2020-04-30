package main

import (
	"sync"
)

type ItemStack struct {
	items []Item
	lock sync.RWMutex
}

func (s *ItemStack) New() *ItemStack{
	s.items = []Item{}
	return s
}

func (s *ItemStack) Push(t Item){
	s.lock.Lock()
	s.items = append(s.items, s)
	s.lock.Unlock()
}

func (s *ItemStack) Pop() *Item{
	s.lock.Lock()
	item := s.items[len(s.items)-1]
	s.items = s.items[0:len(s.items)-1]
	s.lock.Unlock()
	return &item
}