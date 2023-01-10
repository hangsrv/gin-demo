package controller

import (
	"fmt"
	"gin_demo/model"
	"gin_demo/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func RoomInfoList(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("id = ?", id).Find(&room)

	var roomInfo []*model.RoomInfo
	util.GetDb().Table((&model.RoomInfo{}).TableName()).Where("number = ?", room.Number).Order("year desc,month desc").Find(&roomInfo)
	for _, r := range roomInfo {
		r.Fee = float64(room.Area) + r.Electricity*0.8 + r.Water*0.7
	}

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "roomInfo/list.html", gin.H{
		"title":    "详情页",
		"active1":  "active",
		"active2":  "",
		"room":     room,
		"roomInfo": roomInfo,
		"e":        e,
	})
}

func RoomInfoCreate(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("id = ?", id).Find(&room)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "roomInfo/create.html", gin.H{
		"title":   "新增页",
		"active1": "",
		"active2": "active",
		"room":    room,
		"e":       e,
	})
}

func RoomInfoCreateAdd(c *gin.Context) {
	id := c.DefaultQuery("id", "")
	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("id = ?", id).Find(&room)

	var roomInfo model.RoomInfo
	if err := c.Bind(&roomInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	roomInfo.Fee = float64(room.Area) + roomInfo.Electricity*0.8 + roomInfo.Water*0.7
	roomInfo.Number = room.Number
	if err := util.GetDb().Table((&model.RoomInfo{}).TableName()).Create(&roomInfo).Error; err != nil {
		c.JSON(500, gin.H{"error ": "新增错误—" + err.Error()})
		return
	}

	c.Redirect(301, fmt.Sprintf("/roomInfo/list?id=%d", room.ID))
}

func RoomInfoUpdate(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var roomInfo model.RoomInfo
	util.GetDb().Table((&model.RoomInfo{}).TableName()).Where("id = ?", id).Find(&roomInfo)

	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("number = ?", roomInfo.Number).Find(&room)

	e, err := util.Get(c)
	if err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
		return
	}
	c.HTML(http.StatusOK, "roomInfo/update.html", gin.H{
		"title":    "新增页",
		"active1":  "",
		"active2":  "active",
		"room":     room,
		"roomInfo": roomInfo,
		"e":        e,
	})
}

func RoomInfoUpdateAdd(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var roomInfo model.RoomInfo
	if err := c.Bind(&roomInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	var r model.RoomInfo
	util.GetDb().Table((&model.RoomInfo{}).TableName()).Where("id = ?", id).Find(&r)
	roomInfo.Number = r.Number
	var room model.Room
	util.GetDb().Table((&model.Room{}).TableName()).Where("number = ?", roomInfo.Number).Find(&room)
	roomInfo.Fee = float64(room.Area) + roomInfo.Electricity*0.8 + roomInfo.Water*0.7
	if err := util.GetDb().Table((&model.RoomInfo{}).TableName()).Model(&roomInfo).Where("id = ?", id).Updates(&roomInfo).Error; err != nil {
		c.JSON(500, gin.H{"error ": "更新错误-" + err.Error()})
		return
	}

	c.Redirect(301, fmt.Sprintf("/roomInfo/list?id=%d", room.ID))
}

func RoomInfoDelete(c *gin.Context) {
	id := c.DefaultQuery("id", "")

	var roomInfo model.RoomInfo
	util.GetDb().Table((&model.RoomInfo{}).TableName()).Delete(roomInfo, "id = ?", id)

	c.Redirect(301, fmt.Sprintf("/roomInfo/list?id=%s", id))
}
