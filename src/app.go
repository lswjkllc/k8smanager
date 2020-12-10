package main

import (
	"errors"
	"fmt"
	hs "k8smanager/src/handlers"
	ss "k8smanager/src/services"
	us "k8smanager/src/utils"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func addRouter(e *echo.Echo) {
	// e.POST("/users", hs.SaveUser)
	// e.GET("/users/:id", hs.GetUser)
	// e.GET("/show", hs.Show)
	// e.POST("/save", hs.Save)
	// e.POST("/save/file", hs.SaveFile)

	e.GET("/pod", hs.GetPod)
	e.GET("/pod/list", hs.ListPod)

	e.GET("/deployment", hs.GetDeployment)
	e.GET("/deployment/list", hs.ListDeployment)
	e.POST("/deployment/create", hs.CreateDeployment)
	e.POST("/deployment/delete", hs.DeleteDeployment)
	e.POST("/deployment/update", hs.UpdateDeployment)

	e.GET("/namespace", hs.GetNamespace)
	e.GET("/namespace/list", hs.ListNamespace)

	e.GET("/service", hs.GetService)
	e.GET("/service/list", hs.ListService)
}

func aiCheck(c echo.Context, data string) error {
	name := strings.Split(data, "@")[0]
	fmt.Println("name: ", name)

	ks := ss.New()
	_, err := ks.GetNamespace(name)
	if err != nil {
		_, err := ks.CreateNamespace(name)
		return err
	}
	return err
}

// 登陆验证
func login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header
		err := errors.New("Unauthorized")

		if data := header.Get("ai-userid"); data != "" {
			err = aiCheck(c, data)
		}
		// else if data := header.Get("tuya_employee"); data != "" {
		// 	flag = ssoCheck(c, data)
		// }
		// else if data := header.Get("authorization"); data != "" {
		// 	flag = simpleCheck(c, data)
		// }

		if err != nil {
			return us.ResponseJson(c, us.Fail, err.Error(), nil)
		}
		return next(c)

		// var authJson map[string]string
		// authStr := c.Request().Header.Get("authorization")
		// err := json.Unmarshal([]byte(authStr), &authJson)
		// if err != nil {
		// 	return us.ResponseJson(c, us.Fail, "Unauthorized", nil)
		// }

		// return next(c)
	}
}

func main() {
	e := echo.New()
	// 注册中间件
	e.Use(middleware.Recover())
	e.Use(middleware.Logger())
	e.Use(login)
	// 添加路由
	addRouter(e)
	// 添加监听器
	// 开启服务
	e.Logger.Fatal(e.Start(":1323"))
}
