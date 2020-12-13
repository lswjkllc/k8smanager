package handlers

import (
	ms "k8smanager/src/models"
	ss "k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func GetService(c echo.Context) error {
	bp := new(ms.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	service, err := ks.GetService(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildServiceParams(service)

	return us.ResponseJson(c, us.Success, "", data)
}

func GetServiceTypes(c echo.Context) error {
	types := []ms.ServiceType{
		ms.ServiceTypeClusterIP,
		ms.ServiceTypeNodePort,
		ms.ServiceTypeLoadBalancer,
		ms.ServiceTypeExternalName,
	}

	data := map[string]interface{}{"types": types}

	return us.ResponseJson(c, us.Success, "", data)
}

func DeleteService(c echo.Context) error {
	bp := new(ms.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	err := ks.DeleteService(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", nil)
}

func CreateService(c echo.Context) error {
	sp := new(ms.ServiceParams)

	if err := c.Bind(sp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	service, err := ks.CreateService(namespace, sp)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildService(service)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListService(c echo.Context) error {
	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	serviceList, err := ks.ListService(namespace)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	items := serviceList.Items
	size := len(items)

	data := make([]ms.Service, size)
	for i, service := range items {
		data[i] = buildService(&service)
	}

	result := make(map[string]interface{})
	result["data"] = data
	result["size"] = size

	return us.ResponseJson(c, us.Success, "", result)
}
