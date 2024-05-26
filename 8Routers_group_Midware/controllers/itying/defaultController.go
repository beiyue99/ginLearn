package itying

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type DefaultController struct {
}

func (con DefaultController) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"msg": "我是一个Msg",
	})
}
func (con DefaultController) News(c *gin.Context) {
	c.String(200, "News")
}
