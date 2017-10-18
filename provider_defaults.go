package dcfg

// MemoryProvider provides values from an memory store.
type MemoryProvider struct {
	values *memcache
}

// NewMemoryProvider initializes a new DefaultsProvider instance
func NewMemoryProvider() *MemoryProvider {
	return &MemoryProvider{
		values: newCache(),
	}
}

// Set stores a value for the given key that the provider
// returns when queried.
func (p *MemoryProvider) Set(key string, value string) {
	p.values.Set(key, value)
}

// Get returns a previously set default value for the given key or
// ErrKeyMissing when no default is set.
func (p MemoryProvider) Get(key string) (string, error) {
	return p.values.Get(key, func(string) (string, error) {
		return "", ErrKeyMissing
	})
}
