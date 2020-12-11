package main

import (
	"encoding/json"
	"errors"
	hs "k8smanager/src/handlers"
	ss "k8smanager/src/services"
	us "k8smanager/src/utils"
	"strings"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func addRouter(e *echo.Echo) {
	e.GET("/pod", hs.GetPod)
	e.GET("/pod/list", hs.ListPod)
	e.POST("/pod/create", hs.CreatePod)
	e.POST("/pod/update", hs.UpdatePod)
	e.POST("/pod/delete", hs.DeletePod)

	e.GET("/deployment", hs.GetDeployment)
	e.GET("/deployment/list", hs.ListDeployment)
	e.POST("/deployment/create", hs.CreateDeployment)
	e.POST("/deployment/update", hs.UpdateDeployment)
	e.POST("/deployment/delete", hs.DeleteDeployment)

	e.GET("/namespace", hs.GetNamespace)
	e.GET("/namespace/list", hs.ListNamespace)

	e.GET("/service", hs.GetService)
	e.GET("/service/list", hs.ListService)
}

func checkName(name string) (string, error) {
	ks := ss.New()
	_, err := ks.GetNamespace(name)
	if err != nil {
		_, err := ks.CreateNamespace(name)
		return name, err
	}
	return name, err
}

func aiCheck(data string) (string, error) {
	name := strings.Split(data, "@")[0]
	return checkName(name)
}

func ssoCheck(data string) (string, error) {
	var ssoData map[string]string
	err := json.Unmarshal([]byte(data), &ssoData)
	if err != nil {
		return "", err
	}
	name := ssoData["name"]
	return checkName(name)
}

// 登陆验证
func login(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		header := c.Request().Header
		namespace := ""
		err := errors.New("Unauthorized")

		if data := header.Get("authorization"); data != "" {
			c.Logger().Print("ai check...")
			namespace, err = aiCheck(data)
		} else if data := header.Get("sso"); data != "" {
			c.Logger().Print("sso check...")
			namespace, err = ssoCheck(data)
		}

		if err != nil {
			return us.ResponseJson(c, us.Fail, err.Error(), nil)
		}
		// 向 c 增加 namespace
		header["Namespace"] = []string{namespace}

		return next(c)
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
