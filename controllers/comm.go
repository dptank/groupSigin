package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"fmt"
	"io/ioutil"
	"io"
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
/**
http get 请求
*/
func HttpGet(url string) (r *http.Response, e error) {
	resp ,err := http.Get(url)
	if err !=nil {
		fmt.Println(err.Error())
	}
	return resp ,nil
}
/**
http post 请求
*/
func HttpPost(url string,data io.Reader) (res []byte, e error) {
	resp ,err := http.Post(url,"application/json;charset=UTF-8",data)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	return body, nil
}