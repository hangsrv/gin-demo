package util

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var b *gorm.DB

//数据库连接
func init() {
	driverName := "mysql"
	host := "192.168.1.6"
	port := "3306"
	database := "hang"
	username := "root"
	password := "123456"
	charset := "utf8"
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=true&loc=Local",
		username,
		password,
		host,
		port,
		database,
		charset,
	)
	db, err := gorm.Open(driverName, args)
	if err != nil {
		panic(fmt.Sprintf("failed to connect database,err: %s", err.Error()))
	}

	//自动创建数据表
	//db.CreateTable(&model.Department{})
	//db.CreateTable(&model.Employee{})
	//db.CreateTable(&model.Owner{})
	//db.CreateTable(&model.Room{})
	//db.CreateTable(&model.RoomInfo{})
	b = db
}

func GetDb() *gorm.DB {
	return b
}
