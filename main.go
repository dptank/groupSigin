package main

import (
	"net/http"
	"time"
	"groupSigin/routers"
	"github.com/go-ini/ini"
	"fmt"
	"os"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := ini.Load("conf/app.ini")
	if err != nil {
		fmt.Printf("Fail to read file: %v", err)
		os.Exit(1)
	}
	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()
	if mode == "develop" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	//路由初始化
	routersInit := router.InitRoute()

	s := &http.Server{
		Addr:           ":8000",
		Handler:        routersInit,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
