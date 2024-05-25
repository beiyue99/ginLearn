package routers

import (
	"gin-learn/gin-learn/6CustomController/controllers/api"

	"github.com/gin-gonic/gin"
)

func ApiRouterInit(r *gin.Engine) {

	apiRouters := r.Group("/api")
	{

		apiRouters.GET("/", api.ApiController{}.Index)
		apiRouters.GET("/userlist", api.ApiController{}.Userlist)
		apiRouters.GET("/plist", api.ApiController{}.Plist)

	}
}
