package main

import (
	"fmt"
	"github.com/cheekybits/genny/generic"
	"sync"
)

type Key generic.Type
type Value generic.Type

type ValueHashtable struct {
	item map[int]Value
	lock sync.RWMutex
}

func hash(k Key) int{
	key := fmt.Sprintf("%s", k)
	h := 0
	for i:=0; i<len(key); i++{
		h = 31*h + int(key[i])
	}
	return h
}

func (ht *ValueHashtable) Put(k Key, v Value){
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	if ht.item == nil{
		ht.item = make(map[int]Value)
	}
	ht.item[i] = v
}

func(ht *ValueHashtable) Remove(k Key){
	ht.lock.Lock()
	defer ht.lock.Unlock()
	i := hash(k)
	delete (ht.item, i)
}

func (ht *ValueHashtable) Get(k Key) Value{
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	i := hash(k)
	return ht.item[i]
}

func (ht *ValueHashtable) Size() int{
	ht.lock.RLock()
	defer ht.lock.RUnlock()
	return len(ht.item)
}