package dcfg

import (
	"os"
)

// EnvProvider provides values from the environment.
type EnvProvider struct{}

// NewEnvProvider initializes a new EnvProvider instance
func NewEnvProvider() *EnvProvider {
	return &EnvProvider{}
}

// Get retrieves a value from environment variables. If no
// environment variable is set or it is empty, ErrKeyMissing is
// returned
func (p *EnvProvider) Get(key string) (string, error) {
	str := os.Getenv(key)
	if str == "" {
		return "", ErrKeyMissing
	}
	return str, nil
}
