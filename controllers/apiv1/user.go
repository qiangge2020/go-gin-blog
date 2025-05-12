package apiv1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiangge2020/go-gin-blog/models"
	"github.com/qiangge2020/go-gin-blog/pkg/app"
	"github.com/qiangge2020/go-gin-blog/pkg/e"
	"github.com/qiangge2020/go-gin-blog/pkg/util"
	"net/http"
)

func Refresh(c *gin.Context) {
	appG := app.Gin{c}
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	var err error
	pwd, err = util.HashPassword(pwd)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
		return
	}
	user := models.User{Name: name, Password: pwd}
	err = models.AddUser(user)
	if err != nil {
		appG.Response(http.StatusOK, e.ERROR, err)
		return
	}

	appG.Response(http.StatusOK, e.SUCCESS, nil)
}

func Login(c *gin.Context) {
	appG := app.Gin{c}
	name := c.PostForm("name")
	pwd := c.PostForm("pwd")
	var err error
	if util.CheckPasswordHash(pwd) {

	}
}
