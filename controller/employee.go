package controller

import (
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func EmployeeList(c *gin.Context) {

	var employee []model.Employee
	result := util.GetDb().Table((&model.Employee{}).TableName()).Order("id desc").Find(&employee)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "employee/list.html", gin.H{
		"title":    "网站首页",
		"active1":  "active",
		"active2":  "",
		"employee": employee,
		"e":        e,
	})
}

func EmployeeCreate(c *gin.Context) {
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
	c.HTML(http.StatusOK, "employee/create.html", gin.H{
		"title":      "新增页",
		"active1":    "",
		"active2":    "active",
		"department": department,
		"e":          e,
	})
}

func EmployeeCreateAdd(c *gin.Context) {
	var employee model.Employee
	if err := c.Bind(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Employee{}).TableName()).Create(&employee)
	if result.Error != nil {
		c.JSON(500, gin.H{"error ": "新增错误—" + result.Error.Error()})
		return
	}

	c.Redirect(301, "/employee/list")
}

func EmployeeUpdate(c *gin.Context) {
	var department []model.Department
	if err := util.GetDb().Table((&model.Department{}).TableName()).Order("id desc").Find(&department).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.DefaultQuery("id", "")

	//定义结果集数组
	var employee model.Employee
	util.GetDb().Table((&model.Employee{}).TableName()).Where("id = ?", id).Find(&employee)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "employee/update.html", gin.H{
		"title":      "修改页",
		"active1":    "active",
		"active2":    "",
		"employee":   employee,
		"department": department,
		"e":          e,
	})
}

func EmployeeUpdateAdd(c *gin.Context) {
	var employee model.Employee
	if err := c.Bind(&employee); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	id := c.DefaultQuery("id", "")
	result := util.GetDb().Table((&model.Employee{}).TableName()).Model(&employee).Where("id = ?", id).Updates(&employee)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "更新错误-" + result.Error.Error()})
	}

	c.Redirect(301, "/employee/list")
}

func EmployeeDelete(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var user model.Employee
	util.GetDb().Table((&model.Employee{}).TableName()).Delete(user, "id = ?", id)

	c.Redirect(301, "/employee/list")
}
