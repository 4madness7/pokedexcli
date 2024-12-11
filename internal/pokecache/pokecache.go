package pokecache

import (
	"sync"
	"time"
)

func NewCache(interval time.Duration) Cache {
    cache := Cache{
        entries: map[string]cacheEntry{},
        mu: &sync.Mutex{},
    }
    go cache.reaploop(interval)
    return cache
}
