package dcfg

import (
	"errors"
	"fmt"
	"log"
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

// FatalString is a wrapper arount GetString that fails with log.Fatal on error.
func FatalString(key string) string {
	val, err := GetString(key)
	if err != nil {
		log.Fatalf("get string %q: %v", key, err)
	}
	return val
}

// FatalBool is a wrapper arount GetBool that fails with log.Fatal on error.
func FatalBool(key string) bool {
	val, err := GetBool(key)
	if err != nil {
		log.Fatalf("get bool %q: %v", key, err)
	}
	return val
}

// FatalFloat is a wrapper arount GetFloat that fails with log.Fatal on error.
func FatalFloat(key string) float64 {
	val, err := GetFloat(key)
	if err != nil {
		log.Fatalf("get float %q: %v", key, err)
	}
	return val
}

// FatalInt is a wrapper arount GetInt that fails with log.Fatal on error.
func FatalInt(key string) int64 {
	val, err := GetInt(key)
	if err != nil {
		log.Fatalf("get int %q: %v", key, err)
	}
	return val
}

// MustString is a wrapper arount GetString that fails with panic on error.
func MustString(key string) string {
	val, err := GetString(key)
	if err != nil {
		panic(fmt.Errorf("get string %q: %v", key, err))
	}
	return val
}

// MustBool is a wrapper arount GetBool that fails with panic on error.
func MustBool(key string) bool {
	val, err := GetBool(key)
	if err != nil {
		panic(fmt.Errorf("get bool %q: %v", key, err))
	}
	return val
}

// MustFloat is a wrapper arount GetFloat that fails with panic on error.
func MustFloat(key string) float64 {
	val, err := GetFloat(key)
	if err != nil {
		panic(fmt.Errorf("get float %q: %v", key, err))
	}
	return val
}

// MustInt is a wrapper arount GetInt that fails with panic on error.
func MustInt(key string) int64 {
	val, err := GetInt(key)
	if err != nil {
		panic(fmt.Errorf("get int %q: %v", key, err))
	}
	return val
}
