package main

import (
	"strconv"

	"github.com/joho/godotenv"
)

type EnvConfig struct {
	Host      string
	Port      int
	Root      string
	AdminUser string
	AdminPass string
	Debug     bool
}

// Env holds the configuration for the application.
func LoadEnv() (EnvConfig, error) {

	var envMap map[string]string
	envMap, err := godotenv.Read("../.env.example")

	envStruct := EnvConfig{}

	envStruct.Host = envMap["SNOWCP_HOST"]
	port, err := strconv.Atoi(envMap["SNOWCP_PORT"])
	if err != nil {
		envStruct.Port = 8080 // Default port if not set
	} else {
		envStruct.Port = port
	}
	envStruct.Root = envMap["SNOWCP_ROOT"]
	envStruct.AdminUser = envMap["SNOWCP_ADMIN_USER"]
	envStruct.AdminPass = envMap["SNOWCP_ADMIN_PASS"]
	envStruct.Debug = envMap["SNOWCP_DEBUG"] == "true"

	return envStruct, nil
}
