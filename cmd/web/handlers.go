package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// VertualTerminal responds welcome message for now.
func (app *application) VirtualTerminal(c *gin.Context) {
	c.HTML(http.StatusOK, "terminal.tmpl", gin.H{
		"msg": "Successfully render template!",
	})
}
