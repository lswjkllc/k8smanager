package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func GetService(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	service, err := ks.GetService(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildService(service)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListService(c echo.Context) error {
	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	services, err := ks.ListService(namespace)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	items := services.Items
	size := len(items)

	data := make([]models.Service, size)
	for i, service := range items {
		data[i] = buildService(&service)
	}

	result := make(map[string]interface{})
	result["data"] = data
	result["size"] = size

	return us.ResponseJson(c, us.Success, "", result)
}
