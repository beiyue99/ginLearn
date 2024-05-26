package admin

import "github.com/gin-gonic/gin"

type ArticleController struct{ BaseController }

func (con ArticleController) Index(c *gin.Context) {
	// c.String(200, "文章列表---")
	con.success(c)
}
func (con ArticleController) Add(c *gin.Context) {
	c.String(200, "文章---add")
}
func (con ArticleController) Edit(c *gin.Context) {
	c.String(200, "文章---edit")
}
