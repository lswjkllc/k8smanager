package models

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
