package activity

import (
	"github.com/gin-gonic/gin"
	"groupSigin/models"
	"fmt"
)

func  GetInfo(ctx *gin.Context) {
	articleModel := new(models.Articles)
	res := articleModel.First(1)
	fmt.Println(res.Title)
	//id := ctx.GetInt("id")
	ctx.JSON(200, gin.H{
		"status":  12,
		"message": "ceshiddd123",
		"nick":    "ceshisssddd",
	})
}