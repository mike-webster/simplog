package router

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mike-webster/simplog/env"
)

func Start(cfg *env.Config) {
	g := gin.New()
	// allow requests to static assets
	g.Static("/js", "./web/js")
	g.Static("/css", "./web/css")

	// load templates
	g.LoadHTMLGlob("web/templates/*")

	// add routes/handlers
	g.GET("/search", searchHandler)
	g.GET("/", homeHandler)

	// start server
	g.Run(fmt.Sprint(cfg.Server.Host, ":", cfg.Server.Port))
}
