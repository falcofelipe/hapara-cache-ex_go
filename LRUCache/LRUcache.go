package LRUCache

import (
	"fmt"
)

/***** CONSTRUCTORS *****/

type LRUCache struct {
	/* NOTE: Please use New() to instantiate LRUCaches instead */

	/* Required Args */
	capacity int

	/* Optional Args */
	activeCache  map[int]int
	lastUsedKeys []int
}

func New(capacity int) *LRUCache {
	cache := new(LRUCache)
	if capacity <= 0 {
		panic("Please ensure that the number (capacity) passed to New() is > 0.")
	}
	cache.capacity = capacity
	cache.activeCache = make(map[int]int)
	cache.lastUsedKeys = make([]int, 0, capacity)

	fmt.Printf("New LRUCache object created with capacity of %d\n", capacity);

	return cache
}


/***** MAIN METHODS *****/

func (cache *LRUCache) Put(key, value int) {
	if key < 0 || value < 0 {
		panic("LRUCache only supports operations with values >= 0");
	}
	fmt.Printf("PUT: key %d value %d\n", key, value);
	
	lastKeyRemoved := cache.addToRecentKeys(key);

	_, keyExists := cache.activeCache[key]

	if !keyExists && len(cache.activeCache) >= cache.capacity {
		delete(cache.activeCache, lastKeyRemoved)
		fmt.Printf("Removed key %d from the Active Cache\n", lastKeyRemoved);
	}
	
	cache.activeCache[key] = value
	fmt.Println("Active Cache:", cache.activeCache);
}
func (cache LRUCache) Get(key int) int {
	fmt.Printf("GET: key %d\n", key);

	value, keyExists := cache.activeCache[key]
	if !keyExists {
		return -1
	}
	cache.addToRecentKeys(key)
	return value
}
func (cache LRUCache) Delete(key int) int {
	fmt.Printf("DELETE: key %d\n", key);

	value, keyExists := cache.activeCache[key]
	if !keyExists {
		return -1
	}
	delete(cache.activeCache, key)
	cache.deleteFromRecentKeys(key)
	return value
}


/***** SUPPORT METHODS *****/

func (cache *LRUCache) addToRecentKeys(key int) int {
	keyIndex := searchSlice(cache.lastUsedKeys, key)
	if keyIndex > -1 {
		cache.lastUsedKeys = removeOrdered(cache.lastUsedKeys, keyIndex)
	}

	lastKeyRemoved := cache.trimKeysSlice(cache.capacity - 1)

	cache.lastUsedKeys = append(cache.lastUsedKeys, key)
	fmt.Println("Recent Keys:", cache.lastUsedKeys)

	return lastKeyRemoved
}
func (cache *LRUCache) trimKeysSlice(maxCapacity int) int {
	lastKeyRemoved := -1

	for len(cache.lastUsedKeys) > maxCapacity {
		lastKeyRemoved = cache.lastUsedKeys[0]
		cache.lastUsedKeys = cache.lastUsedKeys[1:]
	}
	return lastKeyRemoved
}
func (cache *LRUCache) deleteFromRecentKeys (key int) {
	keyIndex := searchSlice(cache.lastUsedKeys, key)
	if keyIndex > -1 {
		cache.lastUsedKeys = removeOrdered(cache.lastUsedKeys, keyIndex)
		fmt.Println("Recent Keys:", cache.lastUsedKeys)
	} else {
		fmt.Printf("Key %d not found in recent keys\n", key)
	}
}
func removeOrdered(slice []int, s int) []int {
    return append(slice[:s], slice[s+1:]...)
}
func searchSlice(slice []int, value int) int {
	for i := range slice {
		if slice[i] == value {
			return i
		}
	}
	return -1
}