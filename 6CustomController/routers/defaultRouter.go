package routers

import (
	"gin-learn/gin-learn/6CustomController/controllers/itying"

	"github.com/gin-gonic/gin"
)

func DefaultRouterInit(r *gin.Engine) {

	//前台路由
	defaultRouters := r.Group("/")
	{
		defaultRouters.GET("/", itying.DefaultController{}.Index)
		defaultRouters.GET("/news", itying.DefaultController{}.News)

	}
}
