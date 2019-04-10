/*
@Time : 19/4/2 下午2:25 
@Author : gongjiapeng
@File : ClimbStairs.go
@Software: GoLand
*/
package activity

import (
	"github.com/gin-gonic/gin"
	app2 "groupSigin/pkg/app"
	"net/http"
	"groupSigin/pkg/ex"
	"groupSigin/service/climbStairs"
)
/**
爬楼活动首页
*/
type queryIndex struct {
	Id int64 `json:"id"`
}
func GetClimbStairsIndex(ctx *gin.Context) {
	app:=app2.Gin{C:ctx}
	var queryIndex queryIndex
	err := ctx.BindJSON(&queryIndex)
	if err!=nil {
		app.Response(http.StatusBadRequest,ex.ERROR_ID_FAIL,false,"",err.Error())
		return
	}
	//初始化信息
	res := climbStairs.GetClimbStairsInit(queryIndex.Id)
	app.Response(http.StatusOK,ex.SUCCESS,true,res,"")
}
