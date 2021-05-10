package router

import (
	"github.com/gin-gonic/gin"
	"github.com/mike-webster/simplog/app"
	"github.com/mike-webster/simplog/env"
)

func homeHandler(c *gin.Context) {
	c.HTML(200, "home.tmpl", nil)
}

func searchHandler(c *gin.Context) {
	config, err := env.GetConfig()
	if err != nil {
		c.Status(500)
		return
	}

	text := c.Query("text")
	if len(text) < 1 {
		c.Status(400)
		return
	}

	res, err := app.Search(config, text)
	if err != nil {
		c.AbortWithError(500, err)
		return
	}

	ret := struct {
		Items []string `json:"items"`
	}{Items: res}

	c.JSON(200, ret)
}
