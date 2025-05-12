package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/qiangge2020/go-gin-blog/pkg/setting"
	"github.com/qiangge2020/go-gin-blog/pkg/upload"
	v1 "github.com/qiangge2020/go-gin-blog/routers/v1"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	r.StaticFS("/upload", http.Dir(upload.GetImageFullPath()))
	gin.SetMode(setting.ServerSetting.RunMode)
	apiv1 := r.Group("/api")
	v1.Api(apiv1)
	return r
}
