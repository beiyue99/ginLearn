package routers

import "github.com/gin-gonic/gin"

func ApiRouterInit(r *gin.Engine) {

	apiRouters := r.Group("/api")
	{
		apiRouters.GET("/", func(c *gin.Context) {
			c.String(200, "我是一个api")
		})
		apiRouters.GET("/userlist", func(c *gin.Context) {
			c.String(200, "我是一个api接口---userlist")
		})
		apiRouters.GET("/plist", func(c *gin.Context) {
			c.String(200, "我是一个api接口---plist")
		})
	}
}
