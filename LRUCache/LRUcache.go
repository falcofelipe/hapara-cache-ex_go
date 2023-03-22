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
	activeCache  map[int]CacheEntry
	lastUsedKeys []*CacheAccess
}

func New(capacity int) *LRUCache {
	cache := new(LRUCache)
	if capacity <= 0 {
		panic("Please ensure that the number (capacity) passed to New() is > 0.")
	}
	cache.capacity = capacity
	cache.activeCache = make(map[int]CacheEntry)
	cache.lastUsedKeys = make([]*CacheAccess, 0, capacity)

	fmt.Printf("New LRUCache object created with capacity of %d\n", capacity);

	return cache
}

type CacheEntry struct {
	value int
	expireAfterUses int
}
func NewCacheEntry(value, expireAfterUses int) CacheEntry {
	cacheEntry := new(CacheEntry)
	if value < 0 || expireAfterUses <= 0 {
		panic("Please ensure that the value is >=0  and that expireAfterUses is > 0.")
	}
	cacheEntry.value = value
	cacheEntry.expireAfterUses = expireAfterUses

	return *cacheEntry
}


type CacheAccess struct {
	key int
	usesBeforeExpired int
}

/***** MAIN METHODS *****/

func (cache *LRUCache) Put(key int, cacheEntry CacheEntry) {
	value := cacheEntry.value
	expireAfterUses := cacheEntry.expireAfterUses

	if key < 0 || cacheEntry.value < 0 {
		panic("LRUCache only supports operations with values >= 0");
	}
	fmt.Printf("PUT: key %d value %d\n", key, value);
	
	lastKeyRemoved := cache.addToRecentKeys(&CacheAccess{key: key, usesBeforeExpired: expireAfterUses});

	_, keyExists := cache.activeCache[key]

	if !keyExists && len(cache.activeCache) >= cache.capacity {
		// Possibly check if lastKeyRemoved > -1, otherwise decide what to do
		delete(cache.activeCache, lastKeyRemoved)
		fmt.Printf("Removed key %d from the Active Cache\n", lastKeyRemoved);
	}
	
	cache.activeCache[key] = cacheEntry
	fmt.Println("Active Cache:", cache.activeCache);
}
func (cache LRUCache) Get(key int) int {
	fmt.Printf("GET: key %d\n", key);

	cacheEntry, keyExists := cache.activeCache[key]
	if !keyExists {
		return -1
	}
	removeKey := cache.accessKey(key)
	if removeKey > -1 {
		delete(cache.activeCache, removeKey)
	}
	return cacheEntry.value
}
func (cache LRUCache) Delete(key int) int {
	fmt.Printf("DELETE: key %d\n", key);

	cacheEntry, keyExists := cache.activeCache[key]
	if !keyExists {
		return -1
	}
	delete(cache.activeCache, key)
	cache.deleteFromRecentKeys(key)
	return cacheEntry.value
}


/***** SUPPORT METHODS *****/

func (cache *LRUCache) addToRecentKeys(access *CacheAccess) int {
	lastKeyRemoved := -1
	keyIndex := searchLastUsedKeys(cache.lastUsedKeys, access.key)
	if keyIndex > -1 {
		lastKeyRemoved = cache.accessKey(access.key)
	}

	if access.usesBeforeExpired >= 0 {
		lastKeyRemoved = cache.trimLastUsedKeysSlice(cache.capacity - 1)
		cache.lastUsedKeys = append(cache.lastUsedKeys, access)
		fmt.Println("Recent Keys:", &cache.lastUsedKeys)
		return lastKeyRemoved
	} else {
		fmt.Printf("The key %d has been accessed more times than its limit.\n", access.key)
		return keyIndex
	}
}
func (cache *LRUCache) accessKey(key int) int {
	var lastKeyAccess *CacheAccess
	var lastKeyRemoved int
	keyIndex := searchLastUsedKeys(cache.lastUsedKeys, key)

	if keyIndex > -1 {
		lastKeyAccess = cache.lastUsedKeys[keyIndex];
		cache.lastUsedKeys = removeOrderedFromLastUsedKeys(cache.lastUsedKeys, keyIndex)
	} else {
		lastKeyAccess = &CacheAccess{key: key, usesBeforeExpired: cache.activeCache[key].expireAfterUses};
	}

	lastKeyAccess.usesBeforeExpired -= 1
	fmt.Printf("Key %d accessed. Remaining accesses: %d\n", key, lastKeyAccess.usesBeforeExpired)

	if lastKeyAccess.usesBeforeExpired > 0 {
		lastKeyRemoved = cache.trimLastUsedKeysSlice(cache.capacity - 1)
		cache.lastUsedKeys = append(cache.lastUsedKeys, lastKeyAccess)
		index := len(cache.lastUsedKeys) - 1
		fmt.Printf("Access entry added to lastKeysUsed: {%d, %d}\n", cache.lastUsedKeys[index].key, cache.lastUsedKeys[index].usesBeforeExpired)
	} else {
		lastKeyRemoved = lastKeyAccess.key
		fmt.Printf("Access entry removed from lastKeysUsed: {%d, %d}\n", lastKeyAccess.key, lastKeyAccess.usesBeforeExpired)
	}

	return lastKeyRemoved
}
func (cache *LRUCache) deleteFromRecentKeys (key int) {
	keyIndex := searchLastUsedKeys(cache.lastUsedKeys, key)
	if keyIndex > -1 {
		cache.lastUsedKeys = removeOrderedFromLastUsedKeys(cache.lastUsedKeys, keyIndex)
		fmt.Println("Recent Keys:", &cache.lastUsedKeys)
	} else {
		fmt.Printf("Key %d not found in recent keys\n", key)
	}
}

func (cache *LRUCache) trimLastUsedKeysSlice(maxCapacity int) int {
	lastKeyRemoved := -1

	for len(cache.lastUsedKeys) > maxCapacity {
		lastKeyRemoved = cache.lastUsedKeys[0].key
		cache.lastUsedKeys = cache.lastUsedKeys[1:]
	}
	return lastKeyRemoved
}
func removeOrderedFromLastUsedKeys(slice []*CacheAccess, index int) []*CacheAccess {
	if index >= len(slice) || index == -1 {
		fmt.Println("Please make sure that the index being passed to removeOrderedFromLastUsedKeys is < the slice length and >= 0");
		return slice;
	}

	/* This if{} check might not be needed depending on how append() deals with inaccessible indices */
	if index == len(slice) - 1 {
		return slice[:index]
	}
    return append(slice[:index], slice[index+1:]...)
}
func searchLastUsedKeys(slice []*CacheAccess, key int) int {
	for i := range slice {
		if slice[i].key == key {
			return i
		}
	}
	return -1
}