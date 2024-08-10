package utils

import (
	"errors"
	"os"

	"gopkg.in/yaml.v3"
)

type LBStrategy int

type Config struct {
	Port            int      `yaml:"lb_port"`
	MaxAttemptLimit int      `yaml:"max_attempt_limit"`
	Backends        []string `yaml:"backends"`
	Strategy        string   `yaml:"strategy"`
}

const (
	RoundRobin LBStrategy = iota
	LeastConnected
)

const MAX_LB_ATTEMPTS int = 3


func GetLBStrategy(strategy string) LBStrategy {
	switch strategy {
	case "least-connection":
		return LeastConnected
	default:
		return RoundRobin
	}
}

func GetLBConfig() (*Config, error) {
	var config Config

	configFile, err := os.ReadFile("config.yaml")

	if err != nil {
		return nil, err
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		return nil, err
	}

	if len(config.Backends) == 0 {
		return nil, errors.New("backend hosts expected, none provided")
	}

	if config.Port == 0 {
		return nil, errors.New("load balancer port not found")
	}


	return &config, nil

}