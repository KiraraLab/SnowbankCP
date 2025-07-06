package webpages

import (
	"embed"
)

//go:embed build/client/*
var WebPageFS embed.FS
