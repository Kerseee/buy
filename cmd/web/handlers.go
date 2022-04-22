package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VertualTerminal responds welcome message for now.
func (app *application) VirtualTerminal(c *gin.Context) {
	c.String(http.StatusOK, "Welcome to virtual terminal!")
}
