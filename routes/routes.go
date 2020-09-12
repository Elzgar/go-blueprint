package routes

import (
	"go-blueprint/api/users"
	"net/http"

	"github.com/labstack/echo"
)

// Init function called in server.go at root project
func Init() *echo.Echo {
	e := echo.New()

	// default routes
	e.GET("/", func(context echo.Context) error {
		return context.String(http.StatusOK, echo.Version)
	})

	// example writing routes
	e.GET("/users", users.GetAll)
	e.POST("/auth", users.Login)

	return e
}
