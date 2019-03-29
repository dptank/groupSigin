package router

import (
	"github.com/gin-gonic/gin"
	"groupSigin/middleware/jwt"
	"groupSigin/controllers/admin"
	"groupSigin/controllers/activity"
	"groupSigin/controllers/test"
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
		//送礼活动
		admins.POST("climbStairs/save", admin.SaveClimbStairs)
		admins.POST("climbStairs/info", admin.GetClimbStairsInfo)
		admins.POST("climbStairs/list", admin.GetClimbStairsList)
		//送礼活动商品
		admins.POST("climbStairs/item/save", admin.SaveClimbStairsItem)
		admins.POST("climbStairs/item/info", admin.GetClimbStairsItemInfo)
		admins.POST("climbStairs/item/list", admin.GetClimbStairsItemList)
	}
	//送礼活动前端
	//前端路由
	authorized := router.Group("/activity",jwtauth.JWTAuth())
	{
		authorized.POST("info", pinActivity.Info)
	}
	//日志测试
	log := router.Group("log")
	{
		log.POST("test",test.TestLog)
	}
	//kafka测试
	kafka := router.Group("kafka")
	{
		kafka.POST("produce",test.KafkaTest)
		kafka.POST("consoumer",test.KafkaConsoumer)
	}
	return router
}