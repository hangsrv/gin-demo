package middleware

import (
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

func LoginMiddleware(c *gin.Context) {
	uri := c.Request.RequestURI
	if strings.Contains(uri, "log") {
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
	return
}
