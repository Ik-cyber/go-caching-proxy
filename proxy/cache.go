package proxy

import (
	"sync"
	"time"

	"github.com/Ik-cyber/caching-proxy/models"
)

type CacheItem struct {
	Response []byte
	StoredAt time.Time
	TTL      int
}

type Cache struct {
	store map[string]CacheItem
	mu    sync.RWMutex
}

func NewCache(cfg *models.Config) *Cache {
	c := &Cache{
		store: make(map[string]CacheItem),
	}
	// Start cleanup routine
	go c.cleanupExpiredItems()
	return c
}

// Set stores a new response in the cache
func (c *Cache) Set(key string, response []byte, ttl int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.store[key] = CacheItem{
		Response: response,
		StoredAt: time.Now(),
		TTL:      ttl,
	}
}

// Get retrieves a cached response if it exists and is not expired
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, exists := c.store[key]
	if !exists || time.Since(item.StoredAt).Seconds() > float64(item.TTL) {
		return nil, false
	}
	return item.Response, true
}

// cleanupExpiredItems periodically removes expired cache entries
func (c *Cache) cleanupExpiredItems() {
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		c.mu.Lock()
		for key, item := range c.store {
			if time.Since(item.StoredAt).Seconds() > float64(item.TTL) {
				delete(c.store, key)
			}
		}
		c.mu.Unlock()
	}
}
