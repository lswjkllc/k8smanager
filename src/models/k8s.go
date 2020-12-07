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

type Service struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Type      string `json:"type" xml:"type" form:"type" query:"type"`
	Age       int64  `json:"age" xml:"age" form:"age" query:"age"`
}

type ServiceList struct {
	Data []Service `json:"data"`
	Size int       `json:"size"`
}
