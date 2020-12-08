package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"
	"net/http"

	"github.com/labstack/echo"
)

func GetPod(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	name := c.QueryParam("name")

	ks := services.New()

	pod, err := ks.GetPod(namespace, name)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	data := BuildPod(pod)

	return c.JSON(http.StatusOK, data)
}

func ListPod(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	ks := services.New()

	pods, err := ks.ListPod(namespace)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	items := pods.Items
	size := len(items)

	data := make([]models.Pod, size)
	for i, pod := range items {
		data[i] = BuildPod(&pod)
	}

	return c.JSON(http.StatusOK, models.PodList{Data: data, Size: size})
}

func GetDeployment(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	name := c.QueryParam("name")

	ks := services.New()

	deployment, err := ks.GetDeployment(namespace, name)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	data := BuildDeployment(deployment)

	return c.JSON(http.StatusOK, data)
}

func ListDeployment(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	ks := services.New()

	deployments, err := ks.ListDeployment(namespace)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	items := deployments.Items
	size := len(items)

	data := make([]models.Deployment, size)
	for i, deployment := range items {
		data[i] = BuildDeployment(&deployment)
	}

	return c.JSON(http.StatusOK, models.DeploymentList{Data: data, Size: size})
}

func GetNamespace(c echo.Context) error {
	name := c.QueryParam("name")

	ks := services.New()

	namespace, err := ks.GetNamespace(name)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	data := BuildNamespace(namespace)

	return c.JSON(http.StatusOK, data)
}

func ListNamespace(c echo.Context) error {
	ks := services.New()

	namespaces, err := ks.ListNamespace()
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	items := namespaces.Items
	size := len(items)

	data := make([]models.Namespace, size)
	for i, namespace := range items {
		data[i] = BuildNamespace(&namespace)
	}

	return c.JSON(http.StatusOK, models.NamespaceList{Data: data, Size: size})
}

func GetService(c echo.Context) error {
	namespace := c.QueryParam("namespace")
	name := c.QueryParam("name")

	ks := services.New()

	service, err := ks.GetService(namespace, name)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	data := BuildService(service)

	return c.JSON(http.StatusOK, data)
}

func ListService(c echo.Context) error {
	namespace := c.QueryParam("namespace")

	ks := services.New()

	services, err := ks.ListService(namespace)
	if err != nil {
		return c.String(http.StatusOK, err.Error())
	}

	items := services.Items
	size := len(items)

	data := make([]models.Service, size)
	for i, service := range items {
		data[i] = BuildService(&service)
	}

	return c.JSON(http.StatusOK, models.ServiceList{Data: data, Size: size})
}
