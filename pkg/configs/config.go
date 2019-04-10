/*
@Time : 19/3/29 下午10:08 
@Author : gongjiapeng
@File : config.go
@Software: GoLand
*/
package configs
import (
	"github.com/go-ini/ini"
	"groupSigin/pkg/ginlog"
)
/**
获取配置信息
*/
func GetConfig(keys string) string{
	cfg, err := ini.Load("conf/api.ini", "conf/app.ini")
	if err!=nil {
		panic(err)
	}
	defer func() {
		if err:=recover();err!=nil {
			ginlog.LogPrint("初始化配置信息错误",err)
		}
	}()
	mode := cfg.Section("").Key("app_mode").String()
	confValue :=cfg.Section(mode).Key(keys).String()
	return confValue
}