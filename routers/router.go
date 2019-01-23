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
		//获取活动信息
		admin.POST("activity/info", activity.Info)
		//添加活动信息
		admin.POST("activity/add", activity.AddInfo)
		//修改活动信息
		admin.POST("activity/update", activity.UpdateInfo)
	}
	return router
}