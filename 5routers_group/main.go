package main

import (
	"fmt"
	"gin-learn/gin-learn/5routers_group/routers"
	"text/template"

	"github.com/gin-gonic/gin"
)

func Println(str1 string, str2 string) string {
	fmt.Println(str1, str2)
	return str1 + "---" + str2
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

	routers.AdminRouterInit(r)
	routers.ApiRouterInit(r)
	routers.DefaultRouterInit(r)

	if err := r.Run(); err != nil {
		panic(err)
	}
}
