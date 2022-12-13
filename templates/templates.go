package templates

import (
	"embed"
	"net/http"

	"github.com/gofiber/template/html"
)

var (
	//go:embed *.html layouts/*.html
	embeds embed.FS
)

func SetTemplates() (*html.Engine, error) {
	engine := html.NewFileSystem(http.FS(embeds), ".html")

	if err := engine.Load(); err != nil {
		return nil, err
	}

	return engine, nil
}
