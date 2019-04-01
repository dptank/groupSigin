/*
@Time : 19/3/26 下午3:49 
@Author : gongjiapeng
@File : ClimbStairs
@Software: GoLand
*/
package admin

import (
	"github.com/gin-gonic/gin"
	"groupSigin/service/climbStairs"
	app2 "groupSigin/pkg/app"
	"net/http"
	"groupSigin/pkg/ex"
	"groupSigin/pkg/ginlog"
)

/*
*活动添加
*/
func SaveClimbStairs(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	var cs climbStairs.ClimbStairsInfo
	err := ctx.BindJSON(&cs)
	if err != nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,err.Error())
		return
	}
	errs := climbStairs.SaveClimbStairs(&cs)
	if errs!=nil {
		app.Response(http.StatusBadRequest,ex.ERROR_ADD_FAIL,false,"")
		return
	}
	app.Response(http.StatusOK, ex.SUCCESS,true,"")
}
/**
获取详细信息
*/
func GetClimbStairsInfo(ctx *gin.Context)  {
	app := app2.Gin{C:ctx}
	var cs climbStairs.SelectStairs
	err := ctx.BindJSON(&cs)
	if err != nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,err.Error())
		return
	}
	//id ,_:= strconv.ParseInt(ctx.Query("id"),10,64)
	if cs.Id==0 {
		app.Response(http.StatusBadRequest,500,false,"信息不存在")
		return
	}
	ginlog.LogPrint("主键",cs.Id)
	res := climbStairs.GetClimbStairsInfo(cs.Id)
	app.Response(http.StatusOK,ex.SUCCESS,true,res)
}
/**
分页获取信息
*/
func GetClimbStairsList(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	var cs climbStairs.SelectStairs
	err := ctx.BindJSON(&cs)
	if err != nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,err.Error())
		return
	}
	if cs.PageNum==0 {
		cs.PageNum=1
	}
	if cs.PageSize==0 {
		cs.PageSize=10
	}
	res := climbStairs.GetClimbStairsList(&cs)
	app.Response(http.StatusOK,ex.SUCCESS,true,res)
}
