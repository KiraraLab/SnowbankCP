package main

import (
	"fmt"
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

	fmt.Println("SnowCP Environment Configurations:")
	fmt.Printf("Host: %s\n", envStruct.Host)
	fmt.Printf("Port: %d\n", envStruct.Port)
	fmt.Printf("Root: %s\n", envStruct.Root)
	fmt.Printf("Is Debug Mode: %t\n", envStruct.Debug)

	return envStruct, nil
}
