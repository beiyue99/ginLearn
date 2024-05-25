package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Article struct {
	Title   string
	Content string
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		//c.String(http.StatusOK, "你好gin")
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "首页",
			"score": 89,
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
