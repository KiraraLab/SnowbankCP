package webpages

import (
	"embed"
)

//go:embed build/*
var WebPageFiles embed.FS
