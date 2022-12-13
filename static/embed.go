package static

import (
	"embed"
	"net/http"
)

var (
	//go:embed scripts/* styles/*
	embeds embed.FS

	Dir = http.FS(embeds)
)
