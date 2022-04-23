package main

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/gin-contrib/multitemplate"
)

// createRenderer creates a multitemplate.Render instance.
func createRenderer() (multitemplate.Renderer, error) {
	layouts, err := filepath.Glob("./template/layouts/*.html.tmpl")
	if err != nil {
		return nil, err
	}
	pages, err := filepath.Glob("./template/pages/*.html.tmpl")
	if err != nil {
		return nil, err
	}

	// Add pages.
	r := multitemplate.NewRenderer()

	for _, page := range pages {
		files := []string{page}
		files = append(files, layouts...)
		name := strings.Split(filepath.Base(page), ".")[0]
		fmt.Println(name)
		r.AddFromFiles(name, files...)
	}

	return r, nil
}
