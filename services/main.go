package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
	"github.com/kiraralab/snowbankcp/api"
)

func main() {

	splashFigure := figure.NewFigure("SnowbankCP + Go", "", true)
	splashFigure.Print()

	// Initialize configuration
	env, err := LoadEnv()
	if err != nil {
		fmt.Println("Error initializing env:", err)
		os.Exit(1)
	}

	fmt.Println("SnowbankCP Environment Configurations:")
	fmt.Printf("Host: %s\n", env.Host)
	fmt.Printf("Port: %d\n", env.Port)
	fmt.Printf("Root: %s\n", env.Root)
	fmt.Printf("Is Debug Mode: %t\n", env.Debug)

	api.StartServer()
}
