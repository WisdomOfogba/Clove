package output

import "embed"

//go:embed dist/*
var EmbedFS embed.FS
