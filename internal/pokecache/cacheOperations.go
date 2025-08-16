package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) *Cache{
	cache := &Cache{
		cacheMap: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
	}

	go cache.reapLoop(interval)
	return cache
}

func (c *Cache) Add(key string, val []byte){
	c.mu.Lock()
	c.cacheMap[key] = cacheEntry{ 
		createdAt: time.Now().UTC(),
		val: val,
	}
	c.mu.Unlock()
}


func (c *Cache) Get(key string) ([]byte, bool){
	c.mu.Lock()
	defer c.mu.Unlock()
	if i, ok := c.cacheMap[key]; ok{
		return i.val, true
	}
	return nil, false
}


func (c *Cache) reapLoop(interval time.Duration){
	ticker := time.NewTicker(interval)
	for range ticker.C{
		c.reap(time.Now().UTC(), interval)
	}
}

func (c *Cache) reap(now time.Time, last time.Duration){
	c.mu.Lock()
	defer c.mu.Unlock()
	for i, el := range c.cacheMap{
		if el.createdAt.Before(now.Add(-last)) {
			delete(c.cacheMap, i)
		}
	}
}