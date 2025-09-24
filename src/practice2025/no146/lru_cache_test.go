package no146

import (
	"fmt"
	"testing"
)

func TestLRUCache(t *testing.T) {

	//capacity := 2
	//lruCache := Constructor(capacity)
	//lruCache.Put(1, 1)
	//lruCache.Put(2, 2)
	//fmt.Printf("%v\n", lruCache.Get(1))
	//lruCache.Put(3, 3)
	//fmt.Printf("%v\n", lruCache.Get(2))
	//lruCache.Put(4, 4)
	//fmt.Printf("%v\n", lruCache.Get(1))
	//fmt.Printf("%v\n", lruCache.Get(3))
	//fmt.Printf("%v\n", lruCache.Get(4))

	//capacity := 1
	//
	//lruCache := Constructor(capacity)
	//lruCache.Put(2, 1)
	//fmt.Printf("%v\n", lruCache.Get(2))

	//capacity := 2
	//lruCache := Constructor(capacity)
	//lruCache.Put(2, 1)
	//lruCache.Put(1, 1)
	//lruCache.Put(2, 3)
	//lruCache.Put(4, 1)
	//fmt.Printf("%v\n", lruCache.Get(1))
	//fmt.Printf("%v\n", lruCache.Get(2))

	capacity := 2
	lruCache := Constructor(capacity)
	fmt.Printf("%v\n", lruCache.Get(2))
	lruCache.Put(2, 6)
	fmt.Printf("%v\n", lruCache.Get(1))
	lruCache.Put(1, 5)
	lruCache.Put(1, 2)
	fmt.Printf("%v\n", lruCache.Get(1))
	fmt.Printf("%v\n", lruCache.Get(2))
}
