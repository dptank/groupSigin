package router

import (
	"github.com/gin-gonic/gin"
	"groupSigin/controllers/admin"
)

/**
路由
*/
func InitRoute() *gin.Engine{
	router := gin.Default()
	//后台路由
	admin := router.Group("/admin/")
	{
		//获取信息
		admin.POST("activity/info", activity.Info)
		admin.POST("activity/one", activity.GetInfo)
		admin.POST("activity/add", activity.AddInfo)
	}
	return router
}