package v1

import "github.com/gin-gonic/gin"

func Api(g *gin.RouterGroup) {
	v1 := g.Group("/v1")
	{
		v1.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		
	}
}
