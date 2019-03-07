package pinActivity

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"groupSigin/service/pinActivityService"
	"net/http"
	app2 "groupSigin/pkg/app"
	"groupSigin/pkg/ex"
)
/**
获取活动详情
*/
func Info(ctx *gin.Context) {
	//gredis.Set("a","ceshi",10000000)
	//c ,err:= gredis.Get("a")
	//if err!=nil {
	//	fmt.Println(err.Error()+"ceshi111")
	//}
	//fmt.Println(c)

	app := app2.Gin{C:ctx}
	id ,_ := strconv.Atoi(ctx.Query("id"))
	//获取token解密信息
	//tokenInfo := controllers.GetTokenInfo(ctx)
	//userId ,_:= strconv.Atoi(tokenInfo.UserId)
	//fmt.Println(userId)
	//判断信息
	info := pinActivityService.GetInfo(id)
	if info.Id== 0 {
		app.Response(http.StatusBadRequest,500,false,"信息不存在")
		return
	}
	row := make(map[string]interface{})
	row["id"] = info.Id
	row["img"] = info.Img
	row["title"] = info.Title
	row["countLimit"] = info.CountLimit
	row["ownerPrice"] = info.OwnerPrice
	row["memberPrice"] = info.MemberPrice
	row["startTime"] = info.StartTime
	row["endTime"] = info.EndTime
	row["status"] = info.Status
	row["stock"] = info.Stock
	app.Response(http.StatusOK, ex.SUCCESS,true,row)
}