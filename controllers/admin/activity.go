package activity

import (
	"github.com/gin-gonic/gin"
	"groupSigin/service/pinActivityService"
	"strconv"
	"groupSigin/pkg/app"
	"net/http"
	"groupSigin/pkg/ex"
	"fmt"
)
/**
根据id获取详情
*/
func Info(ctx *gin.Context) {
	app := app.Gin{C:ctx}
	id ,_ := strconv.Atoi(ctx.Query("id"))
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
/**
添加活动信息
*/
func AddInfo(ctx *gin.Context)  {
	app := app.Gin{C:ctx}
	var activityInfo pinActivityService.ActivityInfo
	err := ctx.BindJSON(&activityInfo)
	fmt.Println()
	if err!=nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,"")
		return
	}
	//保存活动信息
	errs := pinActivityService.AddActivityInfo(&activityInfo)
	if errs!=nil {
		app.Response(http.StatusBadRequest,ex.ERROR_ADD_FAIL,false,"添加失败")
		return
	}
	app.Response(http.StatusOK,ex.SUCCESS,true,"")
}