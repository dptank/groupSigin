package models

import (
	"time"
	"fmt"
)

type PinActivity struct {
	ID  int `gorm:"primary_key"`
	Title string
	CountLimit int
	OwnerPrice int
	MemberPrice int
	PriceType int
	StartTime int64
	EndTime int64
	Status int
	Img string
	Stock int
	CreatedAt time.Time
	UpdatedAt time.Time
}
/**
设置表名
*/
func (PinActivity) TableName() string {
	return "pin_activity"
}
/**
根据id获取详情
*/
func (pa *PinActivity) GetActivityInfo(activityId int) *PinActivity {
	fmt.Println(activityId)
	db.conn("read").Where(&PinActivity{ID: activityId}).First(pa)
	return pa
}
/**
添加活动
*/
func (pa *PinActivity) PinActivityAdd() error {
	if err := db.conn("write").Create(pa).Error; err != nil {
		return err
	}
	return nil
}