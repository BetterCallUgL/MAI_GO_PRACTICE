package lrucache

import (
	"container/list"
)

type item struct {
	k int
	v int
}

type cache struct {
	*list.List
	cap int
}

func (c cache) Get(key int) (int, bool) {
	for e := c.Front(); e != nil; e = e.Next() {
		w := e.Value.(*item)
		if w.k == key {
			c.List.MoveToFront(e)
			return w.v, true
		}
	}
	return 0, false
}

func (c cache) Set(key, value int) {
	found := false
	for e := c.Front(); e != nil; e = e.Next() {
		w := e.Value.(*item)
		if w.k == key {
			w.v = value
			found = true
			c.List.MoveToFront(e)
			break
		}
	}
	if !found {
		c.List.PushFront(&item{key, value})
		if c.Len() > c.cap {
			c.Remove(c.Back())
		}
	}
}

func (c cache) Clear() {
	c.Init()
	c.cap = 0
}

func (c cache) Range(f func(key, value int) bool) {
	for e := c.Back(); e != nil; e = e.Prev() {
		w := e.Value.(*item)
		if f(w.k, w.v) == false {
			break
		}
	}
}

func New(cap int) Cache {
	return cache{list.New(), cap}
}
