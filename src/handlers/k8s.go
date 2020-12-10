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

	pod, err := ks.GetPod(bp.Namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildPod(pod)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListPod(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()

	pods, err := ks.ListPod(bp.Namespace)
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

	deployment, err := ks.GetDeployment(bp.Namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildDeploymentParams(deployment)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListDeployment(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()

	deployments, err := ks.ListDeployment(bp.Namespace)
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

	deployment, err := ks.CreateDeployment(dps.Namespace, dps)
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

	deployment, err := ks.UpdateDeployment(dps.Namespace, dps)
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

	err := ks.DeleteDeployment(bp.Namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", nil)
}

func GetNamespace(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()

	namespace, err := ks.GetNamespace(bp.Name)
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

	service, err := ks.GetService(bp.Namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildService(service)

	return us.ResponseJson(c, us.Success, "", data)
}

func ListService(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()

	services, err := ks.ListService(bp.Namespace)
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
