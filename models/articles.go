package models

import "fmt"

type Articles struct {
	ID      int
	Title   string
	Author  string
	Content string
	Click   int
	// 避免时区问题，时间简单使用string
	// time.ParseInLocation("2006-01-02 15:04:05",time.Now().Format("2006-01-02 15:04:05"),time.Local)
	CreateTime string
	UpdateTime string
}

// 用id查询一条记录
func (article *Articles) First(id int) *Articles {
	fmt.Println(id)
	db.conn("write").Where(&Articles{ID: id}).First(article)
	//defer orm.Close()
	return article
}