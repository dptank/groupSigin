package jwtauth
import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	app2 "groupSigin/pkg/app"
	"net/http"
)
/**
token解析
*/
func JWTAuth() gin.HandlerFunc{
	return func(ctx *gin.Context) {
		app := app2.Gin{C:ctx}
		tokenString := ctx.Request.Header.Get("token")
		if tokenString=="" {
			app.Response(http.StatusBadRequest,500,false,"token为空")
			ctx.Abort()
			return
		}
		token, err := jwt.Parse(tokenString,secret())
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			ctx.Set("claims", claims)
		} else {
			app.Response(http.StatusBadRequest,500,false,err)
			ctx.Abort()
			return
		}
	}
}
/**
token 秘钥
*/
func secret()jwt.Keyfunc{
	return func(token *jwt.Token) (interface{}, error) {
		return []byte(""),nil
	}
}

