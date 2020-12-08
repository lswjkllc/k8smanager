package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"

	"github.com/labstack/echo"
)

func GetPod(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	name := c.QueryParam("name")

	ks := services.New()

	pod, err := ks.GetPod(namespace, name)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	data := buildPod(pod)

	return responseJson(c, Success, "", data)
}

func ListPod(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	ks := services.New()

	pods, err := ks.ListPod(namespace)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	items := pods.Items
	size := len(items)

	data := make([]models.Pod, size)
	for i, pod := range items {
		data[i] = buildPod(&pod)
	}

	return responseJson(c, Success, "", models.PodList{Data: data, Size: size})
}

func GetDeployment(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	name := c.QueryParam("name")

	ks := services.New()

	deployment, err := ks.GetDeployment(namespace, name)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	data := buildDeployment(deployment)

	return responseJson(c, Success, "", data)
}

func ListDeployment(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	ks := services.New()

	deployments, err := ks.ListDeployment(namespace)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	items := deployments.Items
	size := len(items)

	data := make([]models.Deployment, size)
	for i, deployment := range items {
		data[i] = buildDeployment(&deployment)
	}

	return responseJson(c, Success, "", models.DeploymentList{Data: data, Size: size})
}

func CreateDeployment(c echo.Context) error {
	dps := new(models.DeploymentParams)
	// 绑定
	if err := c.Bind(dps); err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	// 创建
	ks := services.New()
	deployment, err := ks.CreateDeployment(dps.Namespace, dps)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	data := buildDeployment(deployment)

	return responseJson(c, Success, "", data)
}

func GetNamespace(c echo.Context) error {
	name := c.QueryParam("name")

	ks := services.New()

	namespace, err := ks.GetNamespace(name)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	data := buildNamespace(namespace)

	return responseJson(c, Success, "", data)
}

func ListNamespace(c echo.Context) error {
	ks := services.New()

	namespaces, err := ks.ListNamespace()
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	items := namespaces.Items
	size := len(items)

	data := make([]models.Namespace, size)
	for i, namespace := range items {
		data[i] = buildNamespace(&namespace)
	}

	return responseJson(c, Success, "", models.NamespaceList{Data: data, Size: size})
}

func GetService(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	name := c.QueryParam("name")

	ks := services.New()

	service, err := ks.GetService(namespace, name)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	data := buildService(service)

	return responseJson(c, Success, "", data)
}

func ListService(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	ks := services.New()

	services, err := ks.ListService(namespace)
	if err != nil {
		return responseJson(c, Fail, err.Error(), nil)
	}

	items := services.Items
	size := len(items)

	data := make([]models.Service, size)
	for i, service := range items {
		data[i] = buildService(&service)
	}

	return responseJson(c, Success, "", models.ServiceList{Data: data, Size: size})
}
