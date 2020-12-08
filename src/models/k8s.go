package models

type Pod struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Status    string `json:"status" xml:"status" form:"status" query:"status"`
	Age       int64  `json:"age" xml:"age" form:"age" query:"age"`
}

type PodList struct {
	Data []Pod `json:"data"`
	Size int   `json:"size"`
}

type DeploymentStatus struct {
	Replicas            int32 `json:"replicas" xml:"replicas" form:"replicas" query:"replicas"`
	UpdatedReplicas     int32 `json:"updateRplicas" xml:"updateRplicas" form:"updateRplicas" query:"updateRplicas"`
	ReadyReplicas       int32 `json:"readyReplicas" xml:"readyReplicas" form:"readyReplicas" query:"readyReplicas"`
	AvailableReplicas   int32 `json:"availableReplicas" xml:"availableReplicas" form:"availableReplicas" query:"availableReplicas"`
	UnavailableReplicas int32 `json:"unavailableReplicas" xml:"unavailableReplicas" form:"unavailableReplicas" query:"unavailableReplicas"`
}

type Deployment struct {
	Name      string           `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string           `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Status    DeploymentStatus `json:"status" xml:"status" form:"status" query:"status"`
	Age       int64            `json:"age" xml:"age" form:"age" query:"age"`
}

type DeploymentList struct {
	Data []Deployment `json:"data"`
	Size int          `json:"size"`
}

type Namespace struct {
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Age    int64  `json:"age" xml:"age" form:"age" query:"age"`
}

type NamespaceList struct {
	Data []Namespace `json:"data"`
	Size int         `json:"size"`
}
type IntOrString struct {
	Type   int64  `json:"type" xml:"type" form:"type" query:"type"`
	IntVal int32  `json:"intVal" xml:"intVal" form:"intVal" query:"intVal"`
	StrVal string `json:"strVal" xml:"strVal" form:"strVal" query:"strVal"`
}

type ServicePort struct {
	Name       string      `json:"name" xml:"name" form:"name" query:"name"`
	Port       int32       `json:"port" xml:"port" form:"port" query:"port"`
	TargetPort IntOrString `json:"targetPort" xml:"targetPort" form:"targetPort" query:"targetPort"`
	NodePort   int32       `json:"nodePort" xml:"nodePort" form:"nodePort" query:"nodePort"`
}

type Service struct {
	Name        string        `json:"name" xml:"name" form:"name" query:"name"`
	Namespace   string        `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Type        string        `json:"type" xml:"type" form:"type" query:"type"`
	ClusterIP   string        `json:"clusterIP" xml:"clusterIP" form:"clusterIP" query:"clusterIP"`
	ExternalIPs []string      `json:"externalIPs" xml:"externalIPs" form:"externalIPs" query:"externalIPs"`
	Ports       []ServicePort `json:"ports" xml:"ports" form:"ports" query:"ports"`
	Age         int64         `json:"age" xml:"age" form:"age" query:"age"`
}

type ServiceList struct {
	Data []Service `json:"data"`
	Size int       `json:"size"`
}

// type ResourceList struct {
// 	Cpu    string `json:"cpu" xml:"cpu" form:"cpu" query:"cpu"`
// 	Memory string `json:"memory" xml:"memory" form:"memory" query:"memory"`
// }

type ResourceName string

const (
	ResourceCPU              ResourceName = "cpu"
	ResourceMemory           ResourceName = "memory"
	ResourceStorage          ResourceName = "storage"
	ResourceEphemeralStorage ResourceName = "ephemeral-storage"
)

type ResourceList map[ResourceName]string

type ResourceRequirements struct {
	Limits   ResourceList `json:"limits" xml:"limits" form:"limits" query:"limits"`
	Requests ResourceList `json:"requests" xml:"requests" form:"requests" query:"requests"`
}
type EnvVar struct {
	Name  string `json:"name" xml:"name" form:"name" query:"name"`
	Value string `json:"value" xml:"value" form:"value" query:"value"`
}

type DeploymentParams struct {
	Name      string               `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string               `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Image     string               `json:"image" xml:"image" form:"image" query:"image"`
	Resources ResourceRequirements `json:"resources" xml:"resources" form:"resources" query:"resources"`
	Replicas  int32                `json:"replicas" xml:"replicas" form:"replicas" query:"replicas"`
	Env       []EnvVar             `json:"env" xml:"env" form:"env" query:"env"`
}
