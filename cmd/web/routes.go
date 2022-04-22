package main

import "github.com/gin-gonic/gin"

// route routes the endpoints and return a gin engine.
func (app *application) routes() *gin.Engine {
	router := gin.Default()
	router.GET("/virtual-terminal", app.VirtualTerminal)
	return router
}
