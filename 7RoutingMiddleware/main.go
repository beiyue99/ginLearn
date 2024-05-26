package main

import (
	"fmt"
	"text/template"
	"time"

	"github.com/gin-gonic/gin"
)

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "---" + str2
}
func initMiddleWare(c *gin.Context) {

	start := time.Now().UnixNano()
	fmt.Println("1---我是一个中间件")
	c.Next() //不会往下执行，直接跳转执行该请求的剩余处理程序
	//还有个c.Abort方法，会终止接下俩的处理程序，本函数体会执行完
	fmt.Println("2---我是一个中间件")
	end := time.Now().UnixNano()
	fmt.Println("时间差：", end-start) //查看程序执行时间
}

//再配一个全局中间件

func initMiddleWarePro(c *gin.Context) {
	fmt.Println("this is 全局中间件")
}

func main() {
	r := gin.Default()
	//自定义模板函数，要在加载模板之前写
	r.SetFuncMap(template.FuncMap{
		"Println": Println,
	})
	r.LoadHTMLGlob("templates/*")
	//配置静态文件路由
	r.Static("/static", "./static")

	r.Use(initMiddleWarePro) //如果有多个全局中间件，可以写在后面

	//第二个参数是回调函数，可以有多个回调函数，最后一个回调函数是路由的处理程序，之前的都可成为中间件
	//也可以在外部定义这个函数，然后把函数地址传进来
	r.GET("/",
		initMiddleWare,
		// func(c *gin.Context) {
		//     fmt.Println("aaa")
		// },
		func(c *gin.Context) {
			fmt.Println("这是一个首页")
			c.String(200, "gin首页")
		})

	r.GET("/news", func(c *gin.Context) {
		c.String(200, "新闻页面")
	})

	r.GET("/login", initMiddleWare, func(c *gin.Context) {
		c.String(200, "登录页面")
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
