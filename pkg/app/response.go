package app

import (
	"github.com/gin-gonic/gin"
	"groupSigin/pkg/ex"
)

type Gin struct {
	C *gin.Context
}
/**
请求响应
*/
func (g *Gin) Response(httpCode, errCode int,res bool,data interface{}) {
	errmsg := make(map[string]interface{})
	errmsg["ec"] = httpCode
	errmsg["em"] = ex.GetMsg(errCode)
	g.C.JSON(httpCode, gin.H{
		"err": errmsg,
		"success": res,
		"data": data,
	})
	return
}