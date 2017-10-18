package dcfg

import (
	"errors"
	"strconv"
)

var cache = newCache()

// Defaults is a global instance of an MemoryProvider that can be used to
// provide default values when last in the Providers list.
var Defaults = NewMemoryProvider()

// Providers is asked in order for values until the first provider produces one.
var Providers = []Provider{
	NewEnvProvider(),
	NewDockerSecretsProvider(),
	Defaults,
}

// ErrKeyMissing indicated that for the given key no value could be retrieved.
var ErrKeyMissing = errors.New("key not set")

// Provider is used to retrieve a value for a given key. If the provider has no
// value for a key, it should return ErrMissingKey. If the provider fails to
// retrieve a key, it can return an error, which will stop the retrieval and
// returns the error to the caller.
type Provider interface {
	Get(key string) (string, error)
}

// GetString returns the value associated with the key as a string.
func GetString(key string) (string, error) {
	return cache.Get(key, func(key string) (string, error) {
		for _, p := range Providers {
			str, err := p.Get(key)
			if err != nil && err != ErrKeyMissing {
				return "", err
			}

			if err == nil {
				return str, nil
			}
		}

		return "", ErrKeyMissing
	})
}

// GetBool returns the value associated with the key as a boolean.
func GetBool(key string) (bool, error) {
	str, err := GetString(key)
	if err != nil {
		return false, err
	}

	boolean, err := strconv.ParseBool(str)
	return boolean, err
}

// GetFloat returns the value associated with the key as a float.
func GetFloat(key string) (float64, error) {
	str, err := GetString(key)
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseFloat(str, 64)
	return num, err
}

// GetInt returns the value associated with the key as an integer.
func GetInt(key string) (int64, error) {
	str, err := GetString(key)
	if err != nil {
		return 0, err
	}

	num, err := strconv.ParseInt(str, 0, 64)
	return num, err
}
