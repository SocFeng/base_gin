package api

import (
	"base_gin/src/views"
	"github.com/gin-gonic/gin"
)

func InitRoutes(r *gin.Engine) {

	// 配置路由分级
	user := r.Group("/user")
	{
		user.GET("/userName", views.UserName)
		// Gorm 插入数据
		user.GET("/insertOneUserData", views.InsertOneUserData)
		user.GET("/insertManyUserData", views.InsertManyUserData)
		user.GET("/insertUserDataReturnId", views.InsertUserDataReturnId)

		// Gorm 查询数据
		// 基础查询
		user.GET("/queryBaseUserData", views.QueryBaseUserData)
		// 条件查询
		user.GET("/queryWhereUserData", views.QueryWhereUserData)

		// 原生SQL语句操作
		user.GET("/queryUsersCountSQL", views.QueryUsersCountSQL)

		// 接收post数据
		user.POST("getPostData", views.GetPostData)

	}
	zzuser := r.Group("/zzuser")
	{
		zzuser.GET("/createOneZzUser", views.CreateOneZzUser)
		zzuser.GET("/createManyZzUser", views.CreateManyZzUser)
		zzuser.GET("/queryAllZzUser", views.QueryAllZzUser)
		zzuser.GET("/queryWhereZzUser", views.QueryWhereZzUser)
		zzuser.GET("/changeWhereZzUser", views.ChangeWhereZzUser)
		zzuser.GET("/deleteWhereZzUser", views.DeleteWhereZzUser)
		zzuser.GET("/queryJoinZzUser", views.QueryJoinZzUser)
		zzuser.GET("/txExecTimeUser", views.TxExecTimeUser)
	}
	r.GET("/", views.BseReturn)

}
