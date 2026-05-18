package templates

import (
	"embed"
)

var (
	//go:embed *.html partials/*.html
	Embeds embed.FS
)
