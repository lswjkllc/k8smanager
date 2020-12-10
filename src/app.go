package main

import (
	"encoding/json"
	"k8smanager/src/handlers"
	"net/http"

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
	e.POST("/deployment/create", handlers.CreateDeployment)
	e.POST("/deployment/delete", handlers.DeleteDeployment)
	e.POST("/deployment/update", handlers.UpdateDeployment)

	e.GET("/namespace", handlers.GetNamespace)
	e.GET("/namespace/list", handlers.ListNamespace)

	e.GET("/service", handlers.GetService)
	e.GET("/service/list", handlers.ListService)
}

//中间件函数
func authorization(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		var authJson map[string]string
		authStr := c.Request().Header.Get("authorization")
		err := json.Unmarshal([]byte(authStr), &authJson)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized)
		}

		return next(c)
	}
}

func main() {
	e := echo.New()
	// 注册中间件
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(authorization)
	// 添加路由
	addRouter(e)
	// 开启服务
	e.Logger.Fatal(e.Start(":1323"))
}
