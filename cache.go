package dcfg

import (
	"sync"
)

type memcache struct {
	mu   sync.RWMutex
	data map[string]string
}

func newCache() *memcache {
	return &memcache{
		data: map[string]string{},
	}
}

// Set stores the value into the memcache instance.
func (c *memcache) Set(key string, value string) {
	c.mu.Lock()
	c.data[key] = value
	c.mu.Unlock()
}

// Get retrieves the value from cache. If the value is not yet cached the
// compute function is used to create a value.
func (c *memcache) Get(key string, compute func(string) (string, error)) (string, error) {
	var value string
	var err error
	var ok bool

	c.mu.RLock()
	value, ok = c.data[key]
	c.mu.RUnlock()

	if !ok {
		value, err = compute(key)
		if err == nil {
			// only cache success
			c.mu.Lock()
			c.data[key] = value
			c.mu.Unlock()
		}
	}

	return value, err
}
