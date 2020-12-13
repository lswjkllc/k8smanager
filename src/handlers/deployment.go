package handlers

import (
	ms "k8smanager/src/models"
	ss "k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func GetDeployment(c echo.Context) error {
	bp := new(ms.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	deployment, err := ks.GetDeployment(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeploymentParams(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListDeployment(c echo.Context) error {
	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	deployments, err := ks.ListDeployment(namespace)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	items := deployments.Items
	size := len(items)

	data := make([]ms.Deployment, size)
	for i, deployment := range items {
		data[i] = buildDeployment(&deployment)
	}

	result := make(map[string]interface{})
	result["data"] = data
	result["size"] = size

	return us.ResponseJson(c, us.Success, "", result)
}

func CreateDeployment(c echo.Context) error {
	dps := new(ms.DeploymentParams)

	if err := c.Bind(dps); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	deployment, err := ks.ApplyDeployment(namespace, dps, false)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeployment(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func UpdateDeployment(c echo.Context) error {
	dps := new(ms.DeploymentParams)

	if err := c.Bind(dps); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	deployment, err := ks.ApplyDeployment(namespace, dps, true)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeployment(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func DeleteDeployment(c echo.Context) error {
	bp := new(ms.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	err := ks.DeleteDeployment(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", nil)
}
