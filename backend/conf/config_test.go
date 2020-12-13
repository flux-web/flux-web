package conf

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	envFluxUser     = "FLUX_USER"
	envPollInterval = "POLL_INTERVAL"
	envPollTimeOut  = "POLL_TIMEOUT"
	testFluxUser    = "dummy-user"
)

func TestSetEnvConfig(t *testing.T) {
	err := os.Setenv(envFluxUser, testFluxUser)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv(envPollInterval, "420")
	if err != nil {
		t.Fatal(err)
	}
	err = os.Setenv(envPollTimeOut, "963")
	if err != nil {
		t.Fatal(err)
	}

	cfg := Get()
	assert.Equal(t, testFluxUser, cfg.FluxUser)
	assert.Equal(t, 963, cfg.PollTimeout)
	assert.Equal(t, 420, cfg.PollInterval)

}

func TestDefaultConfig(t *testing.T) {
	err := os.Unsetenv(envFluxUser)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Unsetenv(envPollInterval)
	if err != nil {
		t.Fatal(err)
	}
	err = os.Unsetenv(envPollTimeOut)
	if err != nil {
		t.Fatal(err)
	}

	cfg := Get()
	assert.Equal(t, defaultFluxUser, cfg.FluxUser)
	assert.Equal(t, defaultPollTimeout, cfg.PollTimeout)
	assert.Equal(t, defaultPollInterval, cfg.PollInterval)
}
