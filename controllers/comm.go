package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
)

type UserInfo struct {
	UserId string `json:"userId"`
	OpenId string `json:"openId"`
	Username string `json:"username"`
	UnionId string `json:"unionId"`
	Platform string `json:"platform"`
	Env string `json:"env"`
}
/**
获取token解密信息
*/
func GetTokenInfo(ctx *gin.Context) *UserInfo {
	var info UserInfo
	claims := ctx.MustGet("claims").(jwt.MapClaims)
	info.UserId = claims["userId"].(string)
	info.OpenId = claims["openId"].(string)
	info.Username = claims["username"].(string)
	info.UnionId = claims["unionId"].(string)
	info.Env = claims["env"].(string)
	info.Platform = claims["platform"].(string)
	return &info
}