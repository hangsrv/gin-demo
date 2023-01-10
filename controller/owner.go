package controller

import (
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func OwnerList(c *gin.Context) {
	var owners []model.Owner
	result := util.GetDb().Table((&model.Owner{}).TableName()).Order("id desc").Find(&owners)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": result.Error.Error()})
		return
	}
	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "owner/list.html", gin.H{
		"title":   "网站首页",
		"active1": "active",
		"active2": "",
		"owner":   owners,
		"e":       e,
	})
}

func OwnerCreate(c *gin.Context) {
	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "owner/create.html", gin.H{
		"title":   "新增页",
		"active1": "",
		"active2": "active",
		"e":       e,
	})
}

func OwnerCreateAdd(c *gin.Context) {
	var data model.Owner
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Owner{}).TableName()).Create(&data)
	if result.Error != nil {
		c.JSON(500, gin.H{"error ": "新增错误—" + result.Error.Error()})
		return
	}

	c.Redirect(301, "/owner/list")
}

func OwnerUpdate(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	//定义结果集数组
	var owner model.Owner
	util.GetDb().Table((&model.Owner{}).TableName()).Where("id = ?", id).Find(&owner)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "owner/update.html", gin.H{
		"title":   "修改页",
		"active1": "active",
		"active2": "",
		"owner":   owner,
		"e":       e,
	})
}

func OwnerUpdateAdd(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var owner model.Owner
	if err := c.Bind(&owner); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Owner{}).TableName()).Model(&owner).Where("id = ?", id).Updates(&owner)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "更新错误-" + result.Error.Error()})
	}

	c.Redirect(301, "/owner/list")
}

func OwnerDelete(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var user model.Owner
	util.GetDb().Table((&model.Owner{}).TableName()).Delete(user, "id = ?", id)

	c.Redirect(301, "/owner/list")
}
