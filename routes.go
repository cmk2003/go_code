package main

import (
	"ginEssential/controller/student"
	"ginEssential/controller/user"
	"ginEssential/middleware"
	"github.com/gin-gonic/gin"
)

func CollectRoute(r *gin.Engine) *gin.Engine {
	r.POST("api/auth/register", user.Register)
	r.POST("api/auth/login", user.Login)

	//r.GET("api/auth/info", middleware.AuthMiddleware(), controller.Info)
	// 创建一个需要令牌验证的路由组
	authGroup := r.Group("/api/user")
	// 为这个组添加AuthMiddleware中间件
	authGroup.Use(middleware.AuthMiddleware())

	authGroup.GET("/info", user.Info)
	authGroup.GET("/list", user.List)
	authGroup.GET("/getStudentList", student.List)
	authGroup.POST("/addStudent", student.Add)
	return r
}
