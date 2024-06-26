package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type IndexController struct{}

func (con IndexController) Index(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println(username)

	v, ok := username.(string)
	if ok {
		c.String(200, "用户列表--"+v)
	} else {
		c.String(200, "用户列表获取失败")
	}

}
