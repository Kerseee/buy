package main

import "github.com/gin-gonic/gin"

// route routes the endpoints and return a gin engine.
func (app *application) routes() (*gin.Engine, error) {
	// Create a multiRenderer
	r, err := createRenderer()
	if err != nil {
		return nil, err
	}
	// Create gin engine.
	router := gin.Default()
	router.HTMLRender = r
	router.GET("/virtual-terminal", app.VirtualTerminal)
	return router, nil
}
