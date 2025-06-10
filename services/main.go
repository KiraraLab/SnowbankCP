package main

import (
	"fmt"
	"os"

	"github.com/common-nighthawk/go-figure"
)

func main() {

	splashFigure := figure.NewFigure("SnowCP powered by Go", "", true)
	splashFigure.Print()

	// Initialize configuration
	env, err := LoadEnv()
	if err != nil {
		fmt.Println("Error initializing env:", err)
		os.Exit(1)
	}

	
}
