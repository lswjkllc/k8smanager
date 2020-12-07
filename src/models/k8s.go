package models

type Pod struct {
	Name      string `json:"name" xml:"name" form:"name" query:"name"`
	Namespace string `json:"namespace" xml:"namespace" form:"namespace" query:"namespace"`
	Status    string `json:"status" xml:"status" form:"status" query:"status"`
}

type PodList struct {
	Data []Pod `json:"data"`
	Size int   `json:"size"`
}
