package config

import (
	"encoding/json"
	"fmt"
	"interview-service/internal/logger"
	"os"
)

type (
	GrpcConfig struct {
		ServerHost   string `json:"server_host" default:"localhost"`
		UnsecurePort string `json:"unsecure_port" default:"8080"`
	}

	EnvConfig struct {
		JwtSecret string
		Admin     string
	}
)

func LoadConfigFromFile(path string) *GrpcConfig {
	var cfg GrpcConfig
	val, err := os.ReadFile(path)
	if err != nil {
		logger.LogError(fmt.Errorf("error reading file; %v", err), true)
	}

	if unMarErr := json.Unmarshal(val, &cfg); unMarErr != nil {
		logger.LogError(fmt.Errorf("error during unmarshal; %v", err), true)
	}

	return &cfg
}

func LoadEnv() *EnvConfig {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		logger.LogError(fmt.Errorf("failed to read JWT_SECRET from environment"), true)
	}

	username := os.Getenv("ADMIN")
	if username == "" {
		logger.LogError(fmt.Errorf("failed to read USER_NAME from environment"), true)
	}

	return &EnvConfig{
		JwtSecret: jwtSecret,
		Admin:     username,
	}
}
