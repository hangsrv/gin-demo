package middleware

import (
	"gin_demo/util"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func LoginMiddleware(c *gin.Context) {
	uri := c.Request.RequestURI
	if strings.Contains(uri, "log") || strings.Contains(uri, "asset") {
		c.Next()
		return
	}

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}

	if e == nil {
		c.Redirect(301, "/log/toLogin")
	}
	c.Next()
}
