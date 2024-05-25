package main

import (
	"fmt"
	"net/http"
	"text/template"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

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
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, "你好gin")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
			"score": 89,
			"msg":   "我是一个消息",
			"hobby": []string{"eat", "drink", "sleep"},
			"newsList": []interface{}{
				&Article{
					Title:   "新闻标题111",
					Content: "新闻详情111",
				},
				&Article{
					Title:   "新闻标题222",
					Content: "新闻详情222",
				},
			},
			"news": &Article{
				Title:   "新闻标题",
				Content: "新闻内容",
			},
			"date": 1629423555,
		})
	})
	r.GET("/news", func(c *gin.Context) {
		news := &Article{
			Title:   "新闻标题",
			Content: "新闻详情",
		}
		c.HTML(http.StatusOK, "news.html", gin.H{

			"title": "新闻页面",
			"news":  news,
		})
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
