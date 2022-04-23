package main

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
)

// createRenderer creates a multitemplate.Render instance.
func createRenderer() (multitemplate.Renderer, error) {
	layouts, err := filepath.Glob("./template/layouts/*.tmpl")
	if err != nil {
		return nil, err
	}
	pages, err := filepath.Glob("./template/pages/*.tmpl")
	if err != nil {
		return nil, err
	}

	// Add pages.
	r := multitemplate.NewRenderer()

	for _, page := range pages {
		files := []string{page}
		files = append(files, layouts...)
		r.AddFromFiles(filepath.Base(page), files...)
	}

	return r, nil
}
