package main

import (
	"k8smanager/src/handlers"

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

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	addRouter(e)

	e.Logger.Fatal(e.Start(":1323"))

	// api.ConnectK8S("./config/config.yaml")
}
