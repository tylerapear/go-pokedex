package pokecache

import (
	"fmt"
	"time"
	"sync"
)

type cacheEntry struct {
	createdAt	time.Time
	val 		[]byte
}

type Cache struct {
	mu 				sync.Mutex
	interval 		time.Duration
	cacheEntries 	map[string]cacheEntry
}

func NewCache(interval time.Duration) Cache {
	cache := Cache{
		interval: interval,
		cacheEntries: make(map[string]cacheEntry),
	}
	go cache.reapLoop()
	return cache
}

func (c Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.cacheEntries[key] = cacheEntry{
		createdAt: time.Now(),
		val: val,
	}

	//fmt.Printf("Cache added \nkey: %s\nval: %s\n", key, string(val))
}

func (c Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	// Resolve unused fmt import
	foo := fmt.Sprintf("")
	foo = foo + "bar"

	entry, exists := c.cacheEntries[key]
	if !exists {
		return nil, false
	}

	return entry.val, true
}

func (c Cache) GetAll()	 map[string]cacheEntry {
	c.mu.Lock()
	defer c.mu.Unlock()

	return c.cacheEntries
}

func (c cacheEntry) GetVal() []byte {
	return c.val
}

func (c cacheEntry) GetCreatedAt() time.Time {
	return c.createdAt
}

func (c Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		for key, entry := range c.cacheEntries {
			if time.Since(entry.createdAt) > c.interval {
				c.mu.Lock()
				delete(c.cacheEntries, key)
				c.mu.Unlock()
			}
		}
	}
}