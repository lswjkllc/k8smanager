package models

type Namespace struct {
	Name   string `json:"name" xml:"name" form:"name" query:"name"`
	Status string `json:"status" xml:"status" form:"status" query:"status"`
	Age    int64  `json:"age" xml:"age" form:"age" query:"age"`
}
