package main

import "strconv"

var (

	// Name is the name of the application
	Name = "SnowCP"

	// Description is a brief description of the application
	Description = "A modern, lightweight control panel for game servers, built with Go."

	// Version is the current version of SnowCP
	MajerVersion = 0
	MinorVersion = 1
	PatchVersion = 0

	// Version is the full version string of SnowCP
	Version = strconv.Itoa(MajerVersion) + "." + strconv.Itoa(MinorVersion) + "." + strconv.Itoa(PatchVersion)

	Organization = "KiraraLab"
)
