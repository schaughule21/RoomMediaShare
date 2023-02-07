package api

import (
	"RoomMediaShare/server/middleware"

	"github.com/labstack/echo"
)

// Init API Binding
func Init(e *echo.Echo) {

	o := e.Group("/o")
	r := e.Group("/r")
	c := r.Group("/c")
	// Bind Middeleware

	middleware.Init(e, o, r, c)

}
