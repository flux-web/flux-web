package conf

import (
	"os"
	"strconv"

	"github.com/astaxie/beego/logs"
)

const (
	defaultFluxUser     = "Flux-Web"
	defaultPollTimeout  = 900 // seconds
	defaultPollInterval = 300 // milliseconds
)

var l = logs.GetLogger()

type Config struct {
	FluxUser     string
	PollInterval int
	PollTimeout  int
}

func Get() Config {
	c := Config{}

	fluxUser := defaultFluxUser
	if fluxEnvUser := os.Getenv("FLUX_USER"); fluxEnvUser != "" {
		fluxUser = fluxEnvUser
	}

	pollInterval := defaultPollInterval
	if pollEnvInterval := os.Getenv("POLL_INTERVAL"); pollEnvInterval != "" {
		pollEnvInterval, err := strconv.Atoi(pollEnvInterval)
		if err != nil {
			l.Panic(err.Error())
		} else {
			pollInterval = pollEnvInterval
		}
	}

	pollTimeout := defaultPollTimeout
	if pollEnvTimeout := os.Getenv("POLL_TIMEOUT"); pollEnvTimeout != "" {
		pollEnvTimeout, err := strconv.Atoi(pollEnvTimeout)
		if err != nil {
			l.Panic(err.Error())
		} else {
			pollTimeout = pollEnvTimeout
		}
	}

	c.FluxUser = fluxUser
	c.PollInterval = pollInterval
	c.PollTimeout = pollTimeout

	return c
}
