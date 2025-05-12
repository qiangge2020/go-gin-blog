package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/qiangge2020/go-gin-blog/controllers/apiv1"
)

func Api(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		// 注册
		v1.POST("/refresh", apiv1.Refresh)
		// 登入
		v1.POST("/login", apiv1.Login)
	}
}
