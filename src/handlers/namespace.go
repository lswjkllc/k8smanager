package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func DeleteNamespace(c echo.Context) error {
	ks := services.New()
	name := c.Request().Header.Get("Namespace")

	err := ks.DeleteNamespace(name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", nil)
}

func GetNamespace(c echo.Context) error {
	ks := services.New()
	name := c.Request().Header.Get("Namespace")

	namespace, err := ks.GetNamespace(name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildNamespace(namespace)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListNamespace(c echo.Context) error {
	ks := services.New()

	namespaces, err := ks.ListNamespace()
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	items := namespaces.Items
	size := len(items)

	data := make([]models.Namespace, size)
	for i, namespace := range items {
		data[i] = buildNamespace(&namespace)
	}

	result := make(map[string]interface{})
	result["data"] = data
	result["size"] = size

	return us.ResponseJson(c, us.Success, "", result)
}
