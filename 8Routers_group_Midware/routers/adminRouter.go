package routers

import (
	"gin-learn/gin-learn/6CustomController/controllers/admin"
	"gin-learn/gin-learn/8Routers_group_Midware/middlewares"

	"github.com/gin-gonic/gin"
)

func AdminRouterInit(r *gin.Engine) {
	//在路由分组中配置中间件
	//如配置了多个中间件，那么他们之间可以数据共享
	adminRouters := r.Group("/admin", middlewares.InitMiddleware)
	{
		//也可以在控制器里面共享
		adminRouters.GET("/", admin.IndexController{}.Index)
		adminRouters.GET("/user", admin.UserController{}.Index)
		adminRouters.GET("/user/add", admin.UserController{}.Add)
		adminRouters.GET("/user/edit", admin.UserController{}.Edit)

		adminRouters.GET("/article", admin.ArticleController{}.Index)
		adminRouters.GET("/article/add", admin.ArticleController{}.Add)
		adminRouters.GET("/article/edit", admin.UserController{}.Edit)

	}
}
