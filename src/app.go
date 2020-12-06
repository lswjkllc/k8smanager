package main

import (
	"k8smanager/src/handlers"
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func addRouter(e *echo.Echo) {
	e.POST("/users", handlers.SaveUser)
	e.GET("/users/:id", handlers.GetUser)
	e.GET("/show", handlers.Show)
	e.POST("/save", handlers.Save)
	e.POST("/save/file", handlers.SaveFile)
}

func addMiddleware(e *echo.Echo) {
	// Root level middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Group level middleware
	g := e.Group("/admin")
	g.Use(middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		if username == "joe" && password == "secret" {
			return true, nil
		}
		return false, nil
	}))

	// Route level middleware
	track := func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			println("request to /users")
			return next(c)
		}
	}
	e.GET("/users", func(c echo.Context) error {
		return c.String(http.StatusOK, "/users")
	}, track)
}

func main() {
	e := echo.New()

	addMiddleware(e)
	// addRouter(e)

	e.Logger.Fatal(e.Start(":1323"))

	// api.ConnectK8S("./config/config.yaml")
}
