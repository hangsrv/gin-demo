//HTML 模板渲染
package main

import (
	"encoding/gob"
	"fmt"
	"gin_demo/controller"
	"gin_demo/middleware"
	"gin_demo/model"
	"github.com/gin-gonic/gin"
)

func main() {
	gob.Register(model.Employee{})

	route := gin.Default()

	// 定义中间件
	route.Use(middleware.LoginMiddleware)

	//加载模板
	route.LoadHTMLGlob("./frontend/templates/**/*")

	//加载静态文件
	route.Static("/assets", "./frontend/assets")

	// 定义首页
	route.GET("/index", func(c *gin.Context) {
		controller.Index(c)
	})

	//定义路由
	loginRouter := route.Group("/log")
	{
		// 登陆
		loginRouter.GET("/toLogin", controller.ToLogin)
		loginRouter.POST("/doLogin", controller.DoLogin)
		// 退出
		loginRouter.GET("/doLogout", controller.DoLogout)
	}

	//定义路由
	ownerRouter := route.Group("/owner")
	{
		//列表页
		ownerRouter.GET("/list", controller.OwnerList)
		//添加页面
		ownerRouter.GET("/create", controller.OwnerCreate)
		//表单提交
		ownerRouter.POST("/createAdd", controller.OwnerCreateAdd)
		//修改页面
		ownerRouter.GET("/update", controller.OwnerUpdate)
		ownerRouter.POST("/updateAdd", controller.OwnerUpdateAdd)
		//删除操作
		ownerRouter.GET("/delete", controller.OwnerDelete)
	}

	//定义路由
	roomRouter := route.Group("/room")
	{
		//列表页
		roomRouter.GET("/list", controller.RoomList)
		//添加页面
		roomRouter.GET("/create", controller.RoomCreate)
		//表单提交
		roomRouter.POST("/createAdd", controller.RoomCreateAdd)
		//修改页面
		roomRouter.GET("/update", controller.RoomUpdate)
		roomRouter.POST("/updateAdd", controller.RoomUpdateAdd)
		roomRouter.GET("/updateOwner", controller.RoomOwnerUpdate)
		roomRouter.POST("/updateOwnerAdd", controller.RoomOwnerUpdateAdd)
		//删除操作
		roomRouter.GET("/delete", controller.RoomDelete)
	}

	//定义路由
	departmentRouter := route.Group("/department")
	{
		//列表页
		departmentRouter.GET("/list", controller.DepartmentList)
		//添加页面
		departmentRouter.GET("/create", controller.DepartmentCreate)
		//表单提交
		departmentRouter.POST("/createAdd", controller.DepartmentCreateAdd)
		//修改页面
		departmentRouter.GET("/update", controller.DepartmentUpdate)
		departmentRouter.POST("/updateAdd", controller.DepartmentUpdateAdd)
		//删除操作
		departmentRouter.GET("/delete", controller.DepartmentDelete)
	}

	//定义路由
	employeeRouter := route.Group("/employee")
	{
		//列表页
		employeeRouter.GET("/list", controller.EmployeeList)
		//添加页面
		employeeRouter.GET("/create", controller.EmployeeCreate)
		//表单提交
		employeeRouter.POST("/createAdd", controller.EmployeeCreateAdd)
		//修改页面
		employeeRouter.GET("/update", controller.EmployeeUpdate)
		employeeRouter.POST("/updateAdd", controller.EmployeeUpdateAdd)
		//删除操作
		employeeRouter.GET("/delete", controller.EmployeeDelete)
	}

	//定义路由
	roomInfoRouter := route.Group("/roomInfo")
	{
		//列表页
		roomInfoRouter.GET("/list", controller.RoomInfoList)
		//添加页面
		roomInfoRouter.GET("/create", controller.RoomInfoCreate)
		//表单提交
		roomInfoRouter.POST("/createAdd", controller.RoomInfoCreateAdd)
		//修改页面
		roomInfoRouter.GET("/update", controller.RoomInfoUpdate)
		roomInfoRouter.POST("/updateAdd", controller.RoomInfoUpdateAdd)
		//删除操作
		roomInfoRouter.GET("/delete", controller.RoomInfoDelete)
	}

	err := route.Run(":8080")
	if err != nil {
		panic(err.Error())
	}
	fmt.Println("start success~~")
}
