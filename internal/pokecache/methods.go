package pokecache

import (
	"time"
)

func (c *Cache) Add(key string, data []byte) {
    c.mu.Lock()

    c.entries[key] = cacheEntry{
        val: data,
        createdAt: time.Now(),
    }

    c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool){
    c.mu.Lock()
    entry, exists := c.entries[key]
    c.mu.Unlock()
    if exists {
        return entry.val, exists
    }
    return nil, exists
}

func (c *Cache) reaploop(interval time.Duration) {
    ticker := time.NewTicker(interval)
    for {
        tick := <- ticker.C
        c.mu.Lock()
        for key, entry := range c.entries {
            if tick.Sub(entry.createdAt) > interval {
                delete(c.entries, key)
            }
        }
        c.mu.Unlock()
    }
}
