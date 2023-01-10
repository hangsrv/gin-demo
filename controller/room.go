package controller

import (
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoomList(c *gin.Context) {
	var rooms []model.Room
	if err := util.GetDb().Table((&model.Room{}).TableName()).Where("owner_name != ''").Order("id desc").Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "room/list.html", gin.H{
		"title":   "网站首页",
		"active1": "active",
		"active2": "",
		"active3": "",
		"room":    rooms,
		"e":       e,
	})
}

func RoomCreate(c *gin.Context) {
	var owner []model.Owner
	if err := util.GetDb().Table((&model.Owner{}).TableName()).Order("id desc").Find(&owner).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var rooms []model.Room
	if err := util.GetDb().Table((&model.Room{}).TableName()).Where("owner_name = ''").Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "room/create.html", gin.H{
		"title":   "新增页",
		"active1": "",
		"active2": "active",
		"active3": "",
		"owner":   owner,
		"rooms":   rooms,
		"e":       e,
	})
}

func RoomCreateAdd(c *gin.Context) {
	var data model.Room
	if err := c.Bind(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Room{}).TableName()).Create(&data)
	if result.Error != nil {
		c.JSON(500, gin.H{"error ": "新增错误—" + result.Error.Error()})
		return
	}

	c.Redirect(301, "/room/list")
}

func RoomUpdate(c *gin.Context) {
	var owner []model.Owner
	if err := util.GetDb().Table((&model.Owner{}).TableName()).Order("id desc").Find(&owner).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var rooms []model.Room
	if err := util.GetDb().Table((&model.Room{}).TableName()).Where("owner_name = ''").Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.DefaultQuery("id", "")
	//定义结果集数组
	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("id = ?", id).Find(&room)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "room/update.html", gin.H{
		"title":   "修改页",
		"active1": "active",
		"active2": "",
		"active3": "",
		"owner":   owner,
		"room":    room,
		"rooms":   rooms,
		"e":       e,
	})
}

func RoomUpdateAdd(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var room model.Room
	if err := c.Bind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Room{}).TableName()).Model(&room).Where("id = ?", id).Updates(&room)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "更新错误-" + result.Error.Error()})
	}

	c.Redirect(301, "/room/list")
}

func RoomOwnerUpdate(c *gin.Context) {
	var owner []model.Owner
	if err := util.GetDb().Table((&model.Owner{}).TableName()).Order("id desc").Find(&owner).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	var rooms []model.Room
	if err := util.GetDb().Table((&model.Room{}).TableName()).Where("owner_name = ''").Find(&rooms).Error; err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	id := c.DefaultQuery("id", "")
	//定义结果集数组
	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("id = ?", id).Find(&room)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "room/updateOwner.html", gin.H{
		"title":   "修改页",
		"active1": "",
		"active2": "",
		"active3": "active",
		"owner":   owner,
		"room":    room,
		"rooms":   rooms,
		"e":       e,
	})
}

func RoomOwnerUpdateAdd(c *gin.Context) {
	var room model.Room
	if err := c.Bind(&room); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result := util.GetDb().Table((&model.Room{}).TableName()).Model(&room).Where("number = ?", room.Number).Updates(&room)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "更新错误-" + result.Error.Error()})
	}

	c.Redirect(301, "/room/list")
}

func RoomDelete(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	data := make(map[string]interface{})
	data["owner_name"] = ""
	util.GetDb().Table((&model.Room{}).TableName()).Where("id = ?", id).Updates(data)
	c.Redirect(301, "/room/list")
}
