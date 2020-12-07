package handlers

import (
	"k8smanager/src/models"
	"k8smanager/src/services"
	"net/http"
	"time"

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
	// 获取元信息
	pmeta := pod.ObjectMeta
	// 计算时长
	age := time.Now().Unix() - pmeta.CreationTimestamp.Unix()

	data := models.Pod{
		Name: pmeta.Name, Namespace: pmeta.Namespace,
		Status: string(pod.Status.Phase), Age: age}

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
		// 获取元信息
		pmeta := pod.ObjectMeta
		// 计算时长
		age := time.Now().Unix() - pmeta.CreationTimestamp.Unix()
		var mpod = models.Pod{
			Name: pmeta.Name, Namespace: pmeta.Namespace,
			Status: string(pod.Status.Phase), Age: age}
		data[i] = mpod
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
	// 获取元信息
	dmeta := deployment.ObjectMeta
	// 计算运行时长（秒）
	age := time.Now().Unix() - dmeta.CreationTimestamp.Unix()
	// 获取相关状态
	dstatus := deployment.Status
	status := models.DeploymentStatus{
		Replicas:            dstatus.Replicas,
		UpdatedReplicas:     dstatus.UpdatedReplicas,
		ReadyReplicas:       dstatus.ReadyReplicas,
		AvailableReplicas:   dstatus.AvailableReplicas,
		UnavailableReplicas: dstatus.UnavailableReplicas}

	data := models.Deployment{
		Name: dmeta.Name, Namespace: dmeta.Namespace,
		Status: status, Age: age}

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
		// 获取元信息
		dmeta := deployment.ObjectMeta
		// 计算运行时长（秒）
		age := time.Now().Unix() - dmeta.CreationTimestamp.Unix()
		// 获取相关状态
		dstatus := deployment.Status
		status := models.DeploymentStatus{
			Replicas:            dstatus.Replicas,
			UpdatedReplicas:     dstatus.UpdatedReplicas,
			ReadyReplicas:       dstatus.ReadyReplicas,
			AvailableReplicas:   dstatus.AvailableReplicas,
			UnavailableReplicas: dstatus.UnavailableReplicas}
		var mdeployment = models.Deployment{
			Name: dmeta.Name, Namespace: dmeta.Namespace,
			Status: status, Age: age}
		data[i] = mdeployment
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
	// 获取元信息
	nmeta := namespace.ObjectMeta
	// 计算时长
	age := time.Now().Unix() - nmeta.CreationTimestamp.Unix()

	data := models.Namespace{
		Name: nmeta.Name, Age: age,
		Status: string(namespace.Status.Phase)}

	return c.JSON(http.StatusOK, data)
}
