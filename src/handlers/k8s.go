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

	return c.String(http.StatusOK, pod.Name)
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
		var mpod = models.Pod{Name: pod.Name, Namespace: pod.Namespace, Status: string(pod.Status.Phase)}
		data[i] = mpod
	}

	return c.JSON(http.StatusOK, models.PodList{Data: data, Size: size})
}
