package models

import (
	"time"
)

type ClimbStairs struct {
	Id int64 `gorm:"primary_key"`
	Rule string `json:"rule" `
	Title string `json:"title" `
	Status int `json:"status"`
	BgColor string `json:"bg_color" `
	PageHeadImg string `json:"page_head_img"`
	StartTime time.Time `json:"start_time"`
	EndTime time.Time `json:"end_time"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
/**
设置表名
*/
func (ClimbStairs) TableName() string {
	return "activity_climb_stairs"
}
/**
根据id获取详情
*/
func (cs *ClimbStairs)GetInfoClimbStairsById(id int64) *ClimbStairs{
	db.conn("read").Where(&ClimbStairs{Id:id}).First(cs)
	return cs
}
/**
添加活动规则
*/
func (cs *ClimbStairs)AddClimbStairsInfo()error {
	if err := db.conn("write").Create(cs).Error;err!=nil  {
		return err
	}
	return nil
}
/**
修改活动信息
*/
func (cs *ClimbStairs) UpdateClimbStairsInfo()error  {
	if err := db.conn("write").Save(cs).Error;err!=nil  {
		return err
	}
	return nil
}
/**
获取列表
*/
func (cs *ClimbStairs) GetInfoClimbStairsList(pageNum int64, pageSize int64, maps interface{})([]*ClimbStairs, int64) {
	var climbStairs []*ClimbStairs
	var count int64
	offset := (pageNum-1)*pageSize
	err := db.conn("read").Where(maps).Offset(offset).Limit(pageSize).Find(&climbStairs).Error
	if err!=nil {
		return nil, 0
	}
	errs := db.conn("read").Table(cs.TableName()).Where(maps).Count(&count).Error
	if errs!=nil {
		return nil, 0
	}
	return climbStairs ,count
}