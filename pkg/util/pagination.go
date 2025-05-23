package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"github.com/qiangge2020/go-gin-blog/pkg/setting"
)

// 从gin.Context对象的URL参数中获取名为"page"的值，并根据该值计算出相应的偏移量
func GetPage(c *gin.Context) int {
	result := 0
	page := com.StrTo(c.Query("page")).MustInt()
	if page > 0 {
		result = (page - 1) * setting.AppSetting.PageSize
	}
	return result
}
