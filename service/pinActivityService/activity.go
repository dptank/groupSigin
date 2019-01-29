package pinActivityService

import (
	"groupSigin/models"
	"time"
	"github.com/pkg/errors"
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
/**
活动信息修改
*/
func UpdateActivityInfo(pa *ActivityInfo) error {
	var ob models.PinActivity
	//查询信息
	res := ob.GetActivityInfo(pa.Id)
	if res.ID==0 {
		return errors.New("信息不存在")
	}
	res.Title = pa.Title
	res.CountLimit = pa.CountLimit
	res.OwnerPrice = pa.OwnerPrice
	res.MemberPrice = pa.MemberPrice
	res.PriceType = pa.PriceType
	res.StartTime = pa.StartTime
	res.EndTime = pa.EndTime
	res.Status = pa.Status
	res.Stock = pa.Stock
	res.Img = pa.Img
	nowData := time.Now().Local()
	res.UpdatedAt = nowData
	//添加
	err := ob.PinActivityUpdate()
	return err
}
