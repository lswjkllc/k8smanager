package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func GetPod(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	pod, err := ks.GetPod(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildPod(pod)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListPod(c echo.Context) error {
	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	pods, err := ks.ListPod(namespace)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	items := pods.Items
	size := len(items)

	data := make([]models.Pod, size)
	for i, pod := range items {
		data[i] = buildPod(&pod)
	}

	result := make(map[string]interface{})
	result["data"] = data
	result["size"] = size

	return us.ResponseJson(c, us.Success, "", result)
}

func GetDeployment(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	deployment, err := ks.GetDeployment(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeploymentParams(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListDeployment(c echo.Context) error {
	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	deployments, err := ks.ListDeployment(namespace)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	items := deployments.Items
	size := len(items)

	data := make([]models.Deployment, size)
	for i, deployment := range items {
		data[i] = buildDeployment(&deployment)
	}

	result := make(map[string]interface{})
	result["data"] = data
	result["size"] = size

	return us.ResponseJson(c, us.Success, "", result)
}

func CreateDeployment(c echo.Context) error {
	dps := new(models.DeploymentParams)

	if err := c.Bind(dps); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	deployment, err := ks.CreateDeployment(namespace, dps)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeployment(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func UpdateDeployment(c echo.Context) error {
	dps := new(models.DeploymentParams)

	if err := c.Bind(dps); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	deployment, err := ks.UpdateDeployment(namespace, dps)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeployment(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func DeleteDeployment(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	err := ks.DeleteDeployment(namespace, bp.Name)
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
