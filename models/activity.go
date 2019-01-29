package models

import (
	"time"
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
/**
修改活动信息
*/
func (pa *PinActivity) PinActivityUpdate() error {
	if err := db.conn("write").Save(pa).Error; err != nil {
		return err
	}
	return nil
}
//func (u *PinActivity) AfterCreate() (err error) {
//	if (u.ID > 10) {
//		err = errors.New("user id is already greater than 1000")
//	}
//	return
//}
