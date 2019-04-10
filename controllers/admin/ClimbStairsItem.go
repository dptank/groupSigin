/*
@Time : 19/3/26 下午3:49 
@Author : gongjiapeng
@File : ClimbStairsItem
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
	"fmt"
)

/*
*活动商品添加
*/
func SaveClimbStairsItem(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	var cls climbStairs.ClimbStairsItemInfo
	err := ctx.BindJSON(&cls)
	if err != nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,"",err.Error())
		return
	}
	errs := climbStairs.SaveClimbStairsItem(&cls)
	//fmt.Println(errs.Error())
	if errs!=nil {
		ginlog.LogPrint(errs.Error())
		app.Response(http.StatusBadRequest,ex.ERROR_ADD_FAIL,false,"",errs.Error())
		return
	}
	app.Response(http.StatusOK, ex.SUCCESS,true,"","")
}
/**
获取详细信息
*/
func GetClimbStairsItemInfo(ctx *gin.Context)  {
	app := app2.Gin{C:ctx}
	var sl climbStairs.SelectStairsItem
	err := ctx.BindJSON(&sl)
	if err != nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,"",err.Error())
		return
	}
	if sl.Id<=0 {
		app.Response(http.StatusBadRequest,500,false,"","信息不存在")
		return
	}
	ginlog.LogPrint("主键",sl.Id)
	res := climbStairs.GetClimbStairsItemInfo(sl.Id)
	app.Response(http.StatusOK,ex.SUCCESS,true,res,"")
}
/**
分页获取信息
*/
func GetClimbStairsItemList(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	var sl climbStairs.SelectStairsItem
	err := ctx.BindJSON(&sl)
	if err != nil {
		app.Response(http.StatusBadRequest,ex.INVALID_PARAMS,false,"",err.Error())
		return
	}
	if sl.PageNum==0 {
		sl.PageNum=1
	}
	fmt.Println(sl.PageNum)
	if sl.PageSize==0 {
		sl.PageSize=10
	}
	res := climbStairs.GetClimbStairsItemList(&sl)
	app.Response(http.StatusOK,ex.SUCCESS,true,res,"")
}
