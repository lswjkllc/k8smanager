package handlers

import (
	"k8smanager/src/models"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	v1 "k8s.io/api/core/v1"
)

func BuildPod(pod *v1.Pod) models.Pod {
	// 获取元信息
	pmeta := pod.ObjectMeta
	// 计算时长
	age := time.Now().Unix() - pmeta.CreationTimestamp.Unix()

	return models.Pod{
		Name: pmeta.Name, Namespace: pmeta.Namespace,
		Status: string(pod.Status.Phase), Age: age}
}

func BuildDeployment(deployment *appsv1.Deployment) models.Deployment {
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

func BuildNamespace(namespace *v1.Namespace) models.Namespace {
	// 获取元信息
	nmeta := namespace.ObjectMeta
	// 计算运行时长（秒）
	age := time.Now().Unix() - nmeta.CreationTimestamp.Unix()

	return models.Namespace{
		Name: nmeta.Name, Age: age,
		Status: string(namespace.Status.Phase)}
}

func BuildService(service *v1.Service) models.Service {

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
