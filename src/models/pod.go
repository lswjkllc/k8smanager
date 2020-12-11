package models

type Pod struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Status    string `json:"status" xml:"status" form:"status" query:"status"`
	Age       int64  `json:"age" xml:"age" form:"age" query:"age"`
}

type PodParams struct {
	Name         string               `json:"name" xml:"name" form:"name" query:"name"`
	Namespace    string               `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Image        string               `json:"image" xml:"image" form:"image" query:"image"`
	Resources    ResourceRequirements `json:"resources" xml:"resources" form:"resources" query:"resources"`
	Replicas     int32                `json:"replicas" xml:"replicas" form:"replicas" query:"replicas"`
	Env          []EnvVar             `json:"env" xml:"env" form:"env" query:"env"`
	VolumeMounts []VolumeMount        `json:"volumeMounts" xml:"volumeMounts" form:"volumeMounts" query:"volumeMounts"`
}
