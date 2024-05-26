package middlewares

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func InitMiddleware(c *gin.Context) {
	fmt.Println("当前时间:", time.Now())
	fmt.Println("请求地址:", c.Request.URL)
	c.Set("username", "张三")
}
