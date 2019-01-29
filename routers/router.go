package router

import (
	"github.com/gin-gonic/gin"
	"groupSigin/middleware/jwt"
	"groupSigin/controllers/admin"
	"groupSigin/controllers/activity"
)

/**
路由
*/
func InitRoute() *gin.Engine{
	router := gin.Default()
	//后台路由
	admins := router.Group("/admin/")
	{
		//获取活动信息
		admins.POST("activity/info", admin.Info)
		//添加活动信息
		admins.POST("activity/add", admin.AddInfo)
		//修改活动信息
		admins.POST("activity/update", admin.UpdateInfo)
	}
	//前端路由
	authorized := router.Group("/activity",jwtauth.JWTAuth())
	authorized.POST("info", pinActivity.Info)

	return router
}