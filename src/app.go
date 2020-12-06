package main

import (
	api "k8smanager/src/handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	e.GET("/user/:id", api.GetUser)
	e.Logger.Fatal(e.Start(":1323"))

	// api.ConnectK8S("./config/config.yaml")
}
