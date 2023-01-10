package controller

import (
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DepartmentList(c *gin.Context) {
	var department []model.Department
	if err := util.GetDb().Table((&model.Department{}).TableName()).Order("id desc").Find(&department).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "department/list.html", gin.H{
		"title":      "网站首页",
		"active1":    "active",
		"active2":    "",
		"department": department,
		"e":          e,
	})
}

func DepartmentCreate(c *gin.Context) {
	var employee []model.Employee
	if err := util.GetDb().Table((&model.Employee{}).TableName()).Order("id desc").Find(&employee).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "department/create.html", gin.H{
		"title":    "新增页",
		"active1":  "",
		"active2":  "active",
		"e":        e,
		"employee": employee,
	})
}

func DepartmentCreateAdd(c *gin.Context) {
	var department model.Department
	if err := c.Bind(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Department{}).TableName()).Create(&department)
	if result.Error != nil {
		c.JSON(500, gin.H{"error ": "新增错误—" + result.Error.Error()})
		return
	}

	c.Redirect(301, "/department/list")
}

func DepartmentUpdate(c *gin.Context) {
	var employee []model.Employee
	if err := util.GetDb().Table((&model.Employee{}).TableName()).Order("id desc").Find(&employee).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.DefaultQuery("id", "")
	//定义结果集数组
	var department model.Department
	util.GetDb().Table((&model.Department{}).TableName()).Where("id = ?", id).Find(&department)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "department/update.html", gin.H{
		"title":      "修改页",
		"active1":    "active",
		"active2":    "",
		"employee":   employee,
		"department": department,
		"e":          e,
	})
}

func DepartmentUpdateAdd(c *gin.Context) {
	var department model.Department
	if err := c.Bind(&department); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.DefaultQuery("id", "")
	result := util.GetDb().Table((&model.Department{}).TableName()).Model(&department).Where("id = ?", id).Updates(&department)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "更新错误-" + result.Error.Error()})
	}

	c.Redirect(301, "/department/list")
}

func DepartmentDelete(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var user model.Department
	util.GetDb().Table((&model.Department{}).TableName()).Delete(user, "id = ?", id)

	c.Redirect(301, "/department/list")
}
