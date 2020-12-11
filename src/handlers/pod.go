package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func DeletePod(c echo.Context) error {
	bp := new(models.BaseParams)

	if err := c.Bind(bp); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	err := ks.DeletePod(namespace, bp.Name)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", nil)
}

func UpdatePod(c echo.Context) error {
	pps := new(models.PodParams)

	if err := c.Bind(pps); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	pod, err := ks.ApplyPod(namespace, pps, true)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildPod(pod)

	return us.ResponseJson(c, us.Success, "", data)
}

func CreatePod(c echo.Context) error {
	pps := new(models.PodParams)

	if err := c.Bind(pps); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := services.New()
	namespace := c.Request().Header.Get("Namespace")

	pod, err := ks.ApplyPod(namespace, pps, false)
	if err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	data := buildPod(pod)

	return us.ResponseJson(c, us.Success, "", data)
}

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

	data := buildPodParams(pod)

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
