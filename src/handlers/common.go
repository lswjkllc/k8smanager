package handlers

import (
	"encoding/json"
	"k8smanager/src/models"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func buildPod(pod *v1.Pod) models.Pod {
	// 获取元信息
	pmeta := pod.ObjectMeta
	// 计算时长
	age := time.Now().Unix() - pmeta.CreationTimestamp.Unix()

	return models.Pod{
		Name: pmeta.Name, Namespace: pmeta.Namespace,
		Status: string(pod.Status.Phase), Age: age}
}

func buildPodParams(pod *v1.Pod) models.PodParams {
	// 获取元数据
	pmeta := pod.ObjectMeta
	// 获取 Spec
	pspec := pod.Spec
	// 获取 第一个 Container
	pstcon := pspec.Containers[0]

	// 重构 Env: pstcon.Env
	var envs []models.EnvVar
	envParams, _ := json.Marshal(pstcon.Env)
	json.Unmarshal(envParams, &envs)
	// 重构 Resources: pstcon.Resources
	presources := pstcon.Resources
	// limits
	var limits models.ResourceList
	limitParams, _ := json.Marshal(presources.Limits)
	json.Unmarshal(limitParams, &limits)
	// requests
	var requests models.ResourceList
	requestsParams, _ := json.Marshal(presources.Requests)
	json.Unmarshal(requestsParams, &requests)

	resources := models.ResourceRequirements{
		Limits:   limits,
		Requests: requests}

	return models.PodParams{
		Name:      pmeta.Name,
		Image:     pstcon.Image,
		Env:       envs,
		Resources: resources,
		NodeName:  pspec.NodeName,
	}
}

func buildDeploymentParams(deployment *appsv1.Deployment) models.DeploymentParams {
	// 获取元数据
	dmeta := deployment.ObjectMeta
	// 获取 Spec
	dspec := deployment.Spec
	// 获取 Template
	dspecTemp := dspec.Template
	// PodTemp spec
	dst := dspecTemp.Spec
	// 获取 第一个 Container
	dstcon := dst.Containers[0]

	// 重构 Env: dstcon.Env
	var envs []models.EnvVar
	envParams, _ := json.Marshal(dstcon.Env)
	json.Unmarshal(envParams, &envs)
	// 重构 Resources: dstcon.Resources
	dresources := dstcon.Resources
	// limits
	var limits models.ResourceList
	limitParams, _ := json.Marshal(dresources.Limits)
	json.Unmarshal(limitParams, &limits)
	// requests
	var requests models.ResourceList
	requestsParams, _ := json.Marshal(dresources.Requests)
	json.Unmarshal(requestsParams, &requests)

	resources := models.ResourceRequirements{
		Limits:   limits,
		Requests: requests}

	return models.DeploymentParams{
		Name:      dmeta.Name,
		Image:     dstcon.Image,
		Replicas:  *dspec.Replicas,
		Env:       envs,
		Resources: resources,
		NodeName:  dst.NodeName,
	}
}

func buildDeployment(deployment *appsv1.Deployment) models.Deployment {
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

	return models.Deployment{
		Name: dmeta.Name, Namespace: dmeta.Namespace,
		Status: status, Age: age}
}

func buildNamespace(namespace *v1.Namespace) models.Namespace {
	// 获取元信息
	nmeta := namespace.ObjectMeta
	// 计算运行时长（秒）
	age := time.Now().Unix() - nmeta.CreationTimestamp.Unix()

	return models.Namespace{
		Name: nmeta.Name, Age: age,
		Status: string(namespace.Status.Phase)}
}

func buildService(service *v1.Service) models.Service {

	// 获取元信息
	smeta := service.ObjectMeta
	// 计算时长
	age := time.Now().Unix() - smeta.CreationTimestamp.Unix()
	// 获取Spec
	spec := service.Spec
	// 获取Ports
	sports := spec.Ports
	size := len(sports)
	ports := make([]models.ServicePort, size)
	for i, port := range sports {
		ptargetPort := port.TargetPort
		targetPort := models.IntOrString{
			Type:   int64(ptargetPort.Type),
			IntVal: ptargetPort.IntVal,
			StrVal: ptargetPort.StrVal}

		ports[i] = models.ServicePort{
			Name: port.Name, Port: port.Port,
			TargetPort: targetPort,
			NodePort:   port.NodePort}
	}

	return models.Service{
		Name: smeta.Name, Age: age,
		Namespace:   smeta.Namespace,
		Type:        string(spec.Type),
		ClusterIP:   string(spec.ClusterIP),
		ExternalIPs: spec.ExternalIPs,
		Ports:       ports}
}
