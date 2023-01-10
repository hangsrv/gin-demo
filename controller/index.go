package controller

import (
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Index(c *gin.Context) {
	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "index.html", gin.H{
		"title": "网站首页",
		"e":     e,
	})
}
