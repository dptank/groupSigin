/*
@Time : 19/3/29 下午3:16 
@Author : gongjiapeng
@File : climbStairsItem.go
@Software: GoLand
*/
package models

import "time"

type ClimbStairsItem struct {
	Id int64 `gorm:"column:id"`
	ActivityId int64 `gorm:"column:activity_id"`
	TopicId int64 `gorm:"column:topic_id"`
	NeedNum int64 `gorm:"column:need_num"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}
/**
设置表名
*/
func (ClimbStairsItem) TableName() string{
	return "activity_climb_item"
}
/**
根据id获取详情
*/
func (csi *ClimbStairsItem) GetClimbStairsItemInfo(maps interface{}) *ClimbStairsItem{
	db.conn("read").Where(maps).First(csi)
	return csi
}
/**
根据actId获取详情
*/
func (csi *ClimbStairsItem) GetClimbStairsItemInfoByActId(actId int64) []*ClimbStairsItem{
	var climbStairsItem []*ClimbStairsItem
	err := db.conn("read").Where(&ClimbStairsItem{ActivityId:actId}).Find(&climbStairsItem).Error
	if err != nil {
		return nil
	}
	return climbStairsItem
}
/**
保存更新活动商品
*/
func (csi *ClimbStairsItem) SaveClimbStairsItem()error {
	err := db.conn("write").Save(csi).Error
	if err!=nil {
		return err
	}
	return nil
}
/**
分页获取商品图片列表信息
*/
func (csi *ClimbStairsItem) GetClimbStairsItemList(pageNum int64, pageSize int64, maps interface{})([]*ClimbStairsItem, int64){
	var climbStairsItem []*ClimbStairsItem
	var count int64
	offets := (pageNum-1)*pageSize
	err := db.conn("read").Where(maps).Offset(offets).Limit(pageSize).Find(&climbStairsItem).Error
	if err != nil {
		return nil ,0
	}
	errs := db.conn("read").Table(csi.TableName()).Where(maps).Count(&count).Error
	if errs!=nil {
		return nil, 0
	}
	return climbStairsItem ,count
}