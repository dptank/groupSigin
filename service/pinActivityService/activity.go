package pinActivityService

import (
	"groupSigin/models"
	"fmt"
	"time"
)
type ActivityInfo struct {
	Id int `json:"id"`
	Title string `json:"title" binding:"required"`
	CountLimit int `json:"countLimit" binding:"required"`
	OwnerPrice int `json:"ownerPrice" binding:"required"`
	MemberPrice int `json:"memberPrice" binding:"required"`
	PriceType int `json:"priceType" binding:"required"`
	StartTime int64 `json:"startTime" binding:"required"`
	EndTime int64 `json:"endTime" binding:"required"`
	Status int `json:"status"`
	Img string `json:"img" binding:"required"`
	Stock int `json:"stock" binding:"required"`
}
/**
获取活动详情
*/
func GetInfo(id int) (res ActivityInfo){
	var ob models.PinActivity
	var result ActivityInfo
	info := ob.GetActivityInfo(id)

	result.Id = info.ID
	result.Img = info.Img
	result.Title = info.Title
	result.CountLimit = info.CountLimit
	result.OwnerPrice = info.OwnerPrice
	result.MemberPrice = info.MemberPrice
	result.StartTime = info.StartTime
	result.EndTime = info.EndTime
	result.Stock = info.Stock
	result.Status = info.Status
	fmt.Println(result)
	//fmt.Println(info)
	return result
}
/**
添加活动信息
*/
func AddActivityInfo(pa *ActivityInfo) error{
	var ob models.PinActivity
	ob.Title = pa.Title
	ob.CountLimit = pa.CountLimit
	ob.OwnerPrice = pa.OwnerPrice
	ob.MemberPrice = pa.MemberPrice
	ob.PriceType = pa.PriceType
	ob.StartTime = pa.StartTime
	ob.EndTime = pa.EndTime
	ob.Status = pa.Status
	ob.Stock = pa.Stock
	ob.Img = pa.Img
	nowData := time.Now().Local()
	ob.CreatedAt = nowData
	ob.UpdatedAt = nowData
	//添加
	err := ob.PinActivityAdd()
	return err
}
