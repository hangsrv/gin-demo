package controller

import (
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ToLogin(c *gin.Context) {
	c.HTML(http.StatusOK, "login/index.html", nil)
}

func DoLogout(c *gin.Context) {
	if err := util.Remove(c); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "login/index.html", nil)
}

func DoLogin(c *gin.Context) {
	var data model.Employee
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if data.Username == "" || data.Password == "" {
		c.HTML(http.StatusOK, "login/index.html", nil)
		return
	}
	count := 0
	util.GetDb().Table(data.TableName()).Where("username = ? && password = ?", data.Username, data.Password).Find(&data).Count(&count)
	if count > 0 {
		if err := util.Set(c, data); err != nil {
			http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
			return
		}
		c.Redirect(301, "/index")
	} else {
		c.HTML(http.StatusOK, "login/index.html", nil)
	}
}
