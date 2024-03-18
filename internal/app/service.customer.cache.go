package app

import (
	uuid "github.com/satori/go.uuid"
	"time"
)

const ttl = time.Hour * 24 * 30

type cachedValue struct {
	value     uuid.UUID
	timestamp time.Time
}

type CustomerIDMemoryCache struct {
	cache map[string]cachedValue
}

func NewCustomerIDMemoryCache() *CustomerIDMemoryCache {
	cache := &CustomerIDMemoryCache{cache: make(map[string]cachedValue)}
	go cache.invalidator()
	return cache
}

func (c *CustomerIDMemoryCache) GetUserIDFromCustomerID(customer string) (uuid.UUID, error) {
	if cached := c.get(customer); cached != uuid.Nil {
		return cached, nil
	}
	return uuid.Nil, nil
}

func (c *CustomerIDMemoryCache) AddCustomerIDForUser(customer string, user uuid.UUID) error {
	if user == uuid.Nil {
		return nil
	}
	c.cache[customer] = cachedValue{value: user, timestamp: time.Now()}
	return nil
}

func (c *CustomerIDMemoryCache) get(customer string) uuid.UUID {
	if cached, ok := c.cache[customer]; ok {
		return cached.value
	}
	return uuid.Nil
}

func (c *CustomerIDMemoryCache) invalidator() {
	ticker := time.NewTicker(ttl)
	defer ticker.Stop()
	for range ticker.C {
		for key, item := range c.cache {
			if time.Since(item.timestamp) >= ttl {
				delete(c.cache, key)
			}
		}
	}
}
