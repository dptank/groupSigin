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
	authorized := router.Group("/api/",jwtauth.JWTAuth())
	{
		authorized.POST("climbStairs/index", activity.GetClimbStairsIndex)
	}
	//日志测试
	log := router.Group("log")
	{
		log.POST("test",test.TestLog)
		log.POST("topic",test.GetTopic)
	}
	//kafka测试
	kafka := router.Group("kafka")
	{
		kafka.POST("produce",test.KafkaTest)
		kafka.POST("consoumer",test.KafkaConsoumer)
	}
	return router
}