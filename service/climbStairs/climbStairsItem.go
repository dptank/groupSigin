/*
@Time : 19/3/29 下午4:00 
@Author : gongjiapeng
@File : climbStairsItem.go
@Software: GoLand
*/
package climbStairs

import (
	"groupSigin/models"
	"time"
	"math"
	"errors"
)
/**
活动商品数据结构
*/
type ClimbStairsItemInfo struct {
	Id int64 `json:"id"`
	ActivityId int64 `json:"activityId" binding:"required"`
	TopicId int64 `json:"topicId" binding:"required"`
	NeedNum int64 `json:"needNum" binding:"required"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}
type SelectStairsItem struct {
	Id int64 `json:"id"`
	PageNum int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
	Offset int64
}
/**
保存商品信息
*/
func SaveClimbStairsItem(ob *ClimbStairsItemInfo) error{
	var cls models.ClimbStairsItem
	nowData := time.Now().Local()
	//修改只能修改数量
	if ob.Id!=0 {
		res := cls.GetClimbStairsItemInfo(ob.Id)
		if res.Id==0 {
			return errors.New("信息不存在！")
		}
		res.NeedNum = ob.NeedNum
		res.UpdatedAt = nowData
		return res.SaveClimbStairsItem()
	}
	maps := make(map[string]interface{})
	maps["activity_id"] =  ob.ActivityId
	maps["topic_id"] =  ob.TopicId
	info := cls.GetClimbStairsItemInfo(maps)
	if info.Id!=0 {
		return errors.New("该商品已经存在该活动中，请勿重复添加！")
	}
	cls.ActivityId = ob.ActivityId
	cls.NeedNum = ob.NeedNum
	cls.TopicId = ob.TopicId
	cls.CreatedAt = nowData
	cls.UpdatedAt = nowData
	res := cls.SaveClimbStairsItem()
	return res
}
/**
获取详细信息
*/
func GetClimbStairsItemInfo(id int64)  *ClimbStairsItemInfo {
	var cls models.ClimbStairsItem
	maps := make(map[string]interface{})
	maps["id"] =  id
	res := cls.GetClimbStairsItemInfo(maps)
	if res.Id==0 {
		return nil
	}
	info := &ClimbStairsItemInfo{
		Id:res.Id,
		NeedNum:res.NeedNum,
		TopicId:res.TopicId,
		ActivityId:res.ActivityId,
		CreatedAt:res.CreatedAt.Format("2006-01-02 15:04:05"),
		UpdatedAt:res.UpdatedAt.Format("2006-01-02 15:04:05"),
	}
	return info
}
/**
根据活动id获取活动商品
*/
func GetClimbStairsItemInfoByActId(activityId int64) []*ClimbStairsItemInfo {
	var cls models.ClimbStairsItem
	res := cls.GetClimbStairsItemInfoByActId(activityId)
	if res==nil {
		return nil
	}
	var list []*ClimbStairsItemInfo
	for _,value := range res{
		list=append(list,&ClimbStairsItemInfo{
			Id:value.Id,
			NeedNum:value.NeedNum,
			TopicId:value.TopicId,
			ActivityId:value.ActivityId,
			CreatedAt:value.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:value.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return list
}
/**
获取商品列表信息
*/
func GetClimbStairsItemList(sl *SelectStairsItem) map[string]interface{}{
	var cls models.ClimbStairsItem
	info := make(map[string]interface{})
	maps := make(map[string]interface{})
	res ,count :=cls.GetClimbStairsItemList(sl.PageNum,sl.PageSize,maps)
	var list []*ClimbStairsItemInfo
	for _,value := range res{
		list=append(list,&ClimbStairsItemInfo{
			Id:value.Id,
			NeedNum:value.NeedNum,
			TopicId:value.TopicId,
			ActivityId:value.ActivityId,
			CreatedAt:value.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:value.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	info["list"] = list
	info["totalCount"] = count
	info["pageSize"] = sl.PageSize
	info["pageNum"] = sl.PageNum
	info["totalPage"] = math.Ceil(float64(count)/float64(sl.PageSize))
	return info
}
