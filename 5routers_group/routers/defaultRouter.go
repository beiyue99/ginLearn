package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {

	//前台路由
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", func(c *gin.Context) {
			c.HTML(http.StatusOK, "index.html", gin.H{
				"msg": "我是一个Msg",
			})
		})

		defaultRouters.GET("/news", func(c *gin.Context) {
			c.String(200, "新闻")
		})

	}
}
