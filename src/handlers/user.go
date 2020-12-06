package handlers

import (
	"io"
	"net/http"
	"os"

	"k8smanager/src/models"

	"github.com/labstack/echo"
)

// 保存用户
func SaveUser(c echo.Context) error {
	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, u)
	// or
	// return c.XML(http.StatusCreated, u)
}

// 获取用户
func GetUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}

// 展示用户
func Show(c echo.Context) error {
	team := c.QueryParam("team")
	member := c.QueryParam("member")
	return c.String(http.StatusOK, "team:"+team+", member:"+member)
}

// 保存用户
func Save(c echo.Context) error {
	name := c.FormValue("name")
	email := c.FormValue("email")
	return c.String(http.StatusOK, "name:"+name+", email:"+email)
}

// 保存文件
func SaveFile(c echo.Context) error {
	name := c.FormValue("name")

	avatar, err := c.FormFile("avatar")
	if err != nil {
		return err
	}

	src, err := avatar.Open()
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(avatar.Filename)
	if err != nil {
		return err
	}
	defer dst.Close()

	if _, err = io.Copy(dst, src); err != nil {
		return err
	}

	return c.HTML(http.StatusOK, "<b>Thank you! "+name+"<b>")
}
