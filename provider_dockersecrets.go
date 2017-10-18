package dcfg

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// DockerSecretsProvider provides values from files populated by the docker
// secrets mechanism.
type DockerSecretsProvider struct {
	SecretsDir string
}

// NewDockerSecretsProvider initializes a new DockerSecretsProvider instance
func NewDockerSecretsProvider() *DockerSecretsProvider {
	return &DockerSecretsProvider{
		SecretsDir: "/run/secrets/",
	}
}

// Get retrieves a value from a docker secrets file. If no file for the given
// key exists, ErrKeyMissing is returned.
func (p DockerSecretsProvider) Get(key string) (string, error) {
	fh, err := os.Open(filepath.Join(p.SecretsDir, key))
	if os.IsNotExist(err) {
		return "", ErrKeyMissing
	} else if err != nil {
		return "", err
	}

	data, err := ioutil.ReadAll(fh)
	if err != nil {
		return "", err
	}

	str := strings.TrimSpace(string(data))
	return str, nil
}
