package models

import (
	"github.com/go-ini/ini"
	"fmt"
	"os"
	"github.com/jinzhu/gorm"
	"time"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)
type dbconnect struct {
	dbs map[string]*gorm.DB
}
//数据库读写
const (
	readDb = "read"
	writeDb = "write"
)
var db dbconnect
//数据库链接
func (conns *dbconnect) conn(str string) *gorm.DB{
	if orm, ok := conns.dbs[str]; ok {
		return orm
	}
	if conns.dbs == nil {
		conns.dbs =  make(map[string]*gorm.DB)
	}
	fmt.Println( conns.dbs[str])

	// load配置
	cfg, err := ini.Load("conf/database.ini", "conf/app.ini")
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()
	// 主机
	mysqlHost := "mysql."+str+".host"
	host := cfg.Section(mode).Key(mysqlHost).String()
	// 端口
	mysqlPort := "mysql."+str+".port"
	port := cfg.Section(mode).Key(mysqlPort).String()
	// 用户名
	mysqlUsername := "mysql."+str+".username"
	username := cfg.Section(mode).Key(mysqlUsername).String()
	// 密码
	mysqlPassword := "mysql."+str+".password"
	password := cfg.Section(mode).Key(mysqlPassword).String()
	// 数据库名称
	mysqlDbname := "mysql."+str+".dbname"
	dbname := cfg.Section(mode).Key(mysqlDbname).String()
	// 最大空闲连接数
	mysqlMaxIdleConns := "mysql."+str+".max_idle_conns"
	maxIdleConns, err := cfg.Section(mode).Key(mysqlMaxIdleConns).Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	// 最大打开的连接数
	mysqlMaxOpenConns := "mysql."+str+".max_open_conns"
	maxOpenConns, err := cfg.Section(mode).Key(mysqlMaxOpenConns).Int()
	if err != nil {
		fmt.Printf("%v", err)
		os.Exit(1)
	}
	dsn := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbname + "?charset=utf8&parseTime=true&loc=Local"
	fmt.Println(dsn)
	orm, err := gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("Fail to open mysql: %v", err)
		os.Exit(1)
	}
	orm.DB().SetMaxIdleConns(maxIdleConns)
	orm.DB().SetMaxOpenConns(maxOpenConns)
	orm.DB().SetConnMaxLifetime(time.Hour)
	conns.dbs[str] = orm
	return orm
}
/**
关闭链接
*/
func (conn *dbconnect) closeConn(str string)  {
	dbConn := conn.conn(str)
	if dbConn!=nil {
		dbConn.Close()
	}
}
/**
初始化
*/
func init() {
	db.conn(writeDb)
}
