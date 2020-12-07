package main

import (
	"k8smanager/src/handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func addRouter(e *echo.Echo) {
	// e.POST("/users", handlers.SaveUser)
	// e.GET("/users/:id", handlers.GetUser)
	// e.GET("/show", handlers.Show)
	// e.POST("/save", handlers.Save)
	// e.POST("/save/file", handlers.SaveFile)

	e.GET("/pod", handlers.GetPod)
	e.GET("/pod/list", handlers.ListPod)

	e.GET("/deployment", handlers.GetDeployment)
	e.GET("/deployment/list", handlers.ListDeployment)

	e.GET("/namespace", handlers.GetNamespace)
	e.GET("/namespace/list", handlers.ListNamespace)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	addRouter(e)

	e.Logger.Fatal(e.Start(":1323"))

	// api.ConnectK8S("./config/config.yaml")
}
