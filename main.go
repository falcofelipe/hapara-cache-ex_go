package main

import (
	"fmt"

	"github.com/falcofelipe/hapara_cache_ex_go/LRUCache"
)

func main() {
	cache := LRUCache.New(2);

	cache.Put(1, LRUCache.NewCacheEntry(1, 2))
	fmt.Printf("Get 1: %d\n", cache.Get(1)) // returns 1
	fmt.Printf("Get 1: %d\n", cache.Get(1)) // returns 1
	fmt.Printf("Get 1: %d\n", cache.Get(1)) // returns 1
	// cache.Put(2, LRUCache.NewCacheEntry(2, 2))
	// fmt.Printf("Get 1: %d\n", cache.Get(1)) // returns 1
	// cache.Put(3, LRUCache.NewCacheEntry(3, 1)) // evicts key 2
	// fmt.Printf("Get 2: %d\n", cache.Get(2)) // returns -1 (not found)
	// cache.Put(4, LRUCache.NewCacheEntry(4, 2)) // evicts key 1
	// fmt.Printf("Get 1: %d\n", cache.Get(1)) // returns -1 (not found)
	// fmt.Printf("Get 3: %d\n", cache.Get(3)) // returns 3
	// fmt.Printf("Get 4: %d\n", cache.Get(4)) // returns 4
	// fmt.Printf("Delete 3: %d\n", cache.Delete(3)) // returns 3
	// fmt.Printf("Get 3: %d\n", cache.Get(3)) // returns -1 (not found)
}