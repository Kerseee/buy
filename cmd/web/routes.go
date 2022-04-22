package main

// route routes the endpoints.
func (app *application) route() {
	app.router.GET("/virtual-terminal", app.VirtualTerminal)
}
