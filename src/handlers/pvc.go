package handlers

import (
	ms "k8smanager/src/models"
	ss "k8smanager/src/services"
	us "k8smanager/src/utils"

	"github.com/labstack/echo"
)

func CreatePVC(c echo.Context) error {
	pvcs := new(ms.PVCParams)

	if err := c.Bind(pvcs); err != nil {
		return us.ResponseJson(c, us.Fail, err.Error(), nil)
	}

	ks := ss.New()
	namespace := c.Request().Header.Get("Namespace")

	pv, err := ks.CreatePVC(namespace, pvcs)
	if err != nil {
		return us.ResponseCode(c, us.Fail, err.Error(), nil)
	}

	return us.ResponseJson(c, us.Success, "", nil)
}
