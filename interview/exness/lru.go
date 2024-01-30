package main

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
)

var (
	ErrWrongArg      = errors.New("ErrWrongArg")
	ErrInvalidEntiry = errors.New("ErrInvalidEntiry")
)

type LRUCache interface {
	Get(key int) (int, bool, error)
	Put(key int, value int) error
}

func NewLRUCache(limit int) (LRUCache, error) {
	if limit < 1 {
		return nil, ErrWrongArg
	}
	return &lruCache{
		limit:   limit,
		list:    list.New().Init(),
		storage: make(map[int]*list.Element, limit),
	}, nil
}

type lruCache struct {
	sync.RWMutex
	storage map[int]*list.Element
	list    *list.List
	limit   int
}

type entity struct {
	key int
	val int
}

func (c *lruCache) Get(key int) (int, bool, error) {
	c.RLock()
	el, ok := c.storage[key]
	c.RUnlock()
	if !ok {
		return 0, false, nil
	}
	c.list.MoveToFront(el)

	ent, ok := el.Value.(entity)
	if !ok {
		return 0, false, ErrInvalidEntiry
	}

	return ent.val, true, nil
}

func (c *lruCache) Put(key int, value int) error {
	//TODO: check already in storage
	ent := entity{key: key, val: value}
	el := c.list.PushFront(ent)
	c.Lock()
	defer c.Unlock()
	c.storage[key] = el

	return c.checkLimit()
}

func (c *lruCache) checkLimit() error {
	if len(c.storage) <= c.limit {
		return nil
	}
	el := c.list.Back()
	ent, ok := el.Value.(entity)
	c.list.Remove(el)
	if !ok {
		return ErrInvalidEntiry
	}
	delete(c.storage, ent.key)
	return nil
}

func main() {

	cache, err := NewLRUCache(10)
	checkErr(err)

	val, ok, err := cache.Get(123)
	checkErr(err)
	checkFalse(ok, "not exist 123")
	checkInt(0, val)
	fmt.Println("SUCCESS 1")

	err = cache.Put(0, 0)
	checkErr(err)
	val, ok, err = cache.Get(0)
	checkErr(err)
	checkTrue(ok, "not exist 123")
	checkInt(0, val)
	fmt.Println("SUCCESS 2")

	for i := 1; i < 10; i++ {
		err = cache.Put(i, i)
		checkErr(err)
	}
	for i := 0; i < 10; i++ {
		val, ok, err = cache.Get(i)
		checkErr(err)
		checkTrue(ok, fmt.Sprintf("not exist %d", i))
		checkInt(i, val)
	}
	fmt.Println("SUCCESS 3")

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func checkTrue(ok bool, msg string) {
	if !ok {
		panic(msg)
	}
}

func checkFalse(ok bool, msg string) {
	if ok {
		panic(msg)
	}
}

func checkInt(expect, actual int) {
	if expect != actual {
		panic(fmt.Sprintf("%d != %d", expect, actual))
	}
}
