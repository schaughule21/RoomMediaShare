package middleware

import (
	"labstack/echo/middleware"

	"github.com/labstack/echo"
)

// Init middleware
func Init(e *echo.Echo, o *echo.Group, r *echo.Group, c *echo.Group) {

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:  []string{"http://localhost", "http://localhost:8080", "http://localhost:8081", "http://localhost:8088", "http://localhost:3031", "*"},
		AllowMethods:  []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS, echo.PATCH, echo.HEAD},
		ExposeHeaders: []string{echo.HeaderAuthorization},
	}))
}
