package gredis

import (
	"github.com/gomodule/redigo/redis"
	"github.com/go-ini/ini"
	"os"
	"time"
	"encoding/json"
	"fmt"
)
//连接池
var RedisConn *redis.Pool
/**
redis 链接池
*/
func SetUpRedis() error {
	cfg, err := ini.Load("conf/database.ini", "conf/app.ini")
	if err != nil {
		os.Exit(1)
	}
	// 运行模式
	mode := cfg.Section("").Key("app_mode").String()
	//redis 基础信息
	host := cfg.Section(mode).Key("redis.host").String()
	port := cfg.Section(mode).Key("redis.port").String()
	pwd := cfg.Section(mode).Key("redis.password").String()
	maxIdleConns, _ := cfg.Section(mode).Key("redis.max_idle_conns").Int()
	maxOpenConns, _ := cfg.Section(mode).Key("redis.max_open_conns").Int()
	maxTimeOut ,_:= cfg.Section(mode).Key("redis.max_idle_timeout").Int()
	RedisConn = &redis.Pool{
		MaxIdle:     maxIdleConns,
		MaxActive:   maxOpenConns,
		IdleTimeout: time.Duration(maxTimeOut)* time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", host+":"+port)
			if err != nil {
				return nil, err
			}
			if pwd != "" {
				if _,err := c.Do("AUTH",pwd);err !=nil {

					c.Close()
					return nil,err
				}
				fmt.Println(c)
			}
			return c,err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}
	return  nil
}
/**
set 值
*/
func Set(key string, data interface{}, time int) error{

	conn := RedisConn.Get()
	defer conn.Close()
	value,err := json.Marshal(data)
	//fmt.Println(value)
	if err !=nil {
		return err
	}
	_,err  = conn.Do("SET",key,value)
	if err !=nil {
		return err
	}
	_, err = conn.Do("EXPIRE", key, time)
	if err !=nil {
		return err
	}
	return nil
}
/*
判断是否存在
*/
func Exists(key string) bool{
	conn := RedisConn.Get()
	defer conn.Close()
	exists ,err := redis.Bool(conn.Do("EXISTS",key))
	if err != nil {
		return false
	}
	return exists
}
/**
获取key值
*/
func Get(key string) ([]byte, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	reply, err := redis.Bytes(conn.Do("GET", key))
	if err != nil {
		return nil, err
	}
	return reply, nil
}
/**
删除key
*/
func Delete(key string) (bool, error) {
	conn := RedisConn.Get()
	defer conn.Close()
	return redis.Bool(conn.Do("DEL", key))
}
