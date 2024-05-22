package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

func main() {
	r := gin.Default()
	//配置模板的文件
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.String(200, "值：%v", "首页")
	})
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			//H实际就是封装的 map[string]interface{}
			"message": "pong",
		})
	})
	r.GET("/json1", func(c *gin.Context) {
		a := &Article{
			Title:   "我是标题",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSON(200, a)
	})
	r.GET("/jsonp", func(c *gin.Context) {
		a := &Article{
			Title:   "我是标题-jsonp",
			Desc:    "描述",
			Content: "测试内容",
		}
		c.JSONP(200, a)
	})
	//http://localhost:8080/jsonp?callback=xxx
	//xxx({"title":"我是标题-jsonp","desc":"描述","content":"测试内容"});

	r.GET("/news", func(c *gin.Context) {
		c.HTML(http.StatusOK, "news.html", gin.H{
			"title": "我是news后台的数据",
		})
	})
	r.GET("/goods", func(c *gin.Context) {
		c.HTML(http.StatusOK, "goods.html", gin.H{
			"title": "我是goods后台的数据",
			"price": 20,
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
