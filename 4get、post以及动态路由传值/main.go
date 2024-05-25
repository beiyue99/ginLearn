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
type UserInfo struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
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

	// http://localhost:8080/ticle?id=12  获取get传值
	r.GET("/ticle", func(c *gin.Context) {
		id := c.DefaultQuery("id", "1")
		c.JSON(http.StatusOK, gin.H{
			"msg": "新闻详情",
			"id":  id,
		})
	})

	//获取post传值演示
	r.GET("/user", func(c *gin.Context) {
		c.HTML(http.StatusOK, "user.html", gin.H{})
	})
	//获取表单post过来的数据
	r.POST("/doAddUser", func(c *gin.Context) {
		username := c.PostForm("username")
		password := c.PostForm("password")
		age := c.DefaultPostForm("age", "20")
		//通过c.json返回数据
		c.JSON(http.StatusOK, gin.H{
			"Username": username,
			"Password": password,
			"age":      age,
		})
	})

	//获取GET POST传递的数据绑定到结构体
	r.GET("/getUser", func(c *gin.Context) {
		user := &UserInfo{}
		if err := c.ShouldBind(&user); err == nil {
			c.JSON(http.StatusOK, user)
		} else {
			c.JSON(http.StatusOK, gin.H{
				"err": err.Error(),
			})
		}
	})

	//动态路由传值
	r.GET("/list/:cid", func(c *gin.Context) {
		cid := c.Param("cid")
		c.String(http.StatusOK, "%v", cid)
	})

	if err := r.Run(); err != nil {
		panic(err)
	}
}
