package climbStairs

import (
	"groupSigin/models"
	"time"
	"fmt"
	"math"
)

type ClimbStairsInfo struct {
	Id int64 `json:"id"`
	Rule string `json:"Rule" binding:"required"`
	Title string `json:"title" binding:"required"`
	Status int `json:"status"`
	BgColor string `json:"bgColor" binding:"required"`
	PageHeadImg string `json:"pageHeadImg" binding:"required"`
	StartTime string `json:"startTime" binding:"required"`
	EndTime string `json:"endTime" binding:"required"`
}
type SelectStairs struct {
	Id int64 `json:"id"`
	PageNum int64 `json:"pageNum"`
	PageSize int64 `json:"pageSize"`
	Offset int64
}
/**
添加信息
*/
func SaveClimbStairs(ob *ClimbStairsInfo) error {
	var climbStairs models.ClimbStairs
	//时间转换
	nowData := time.Now().Local()
	starTime := ob.StartTime
	endTime := ob.EndTime
	timeLayout := "2006-01-02 15:04:05"
	loc, _ := time.LoadLocation("Local")
	st, _ := time.ParseInLocation(timeLayout, starTime, loc)
	et,_:=time.ParseInLocation(timeLayout, endTime, loc)

	climbStairs.BgColor = ob.BgColor
	climbStairs.PageHeadImg = ob.PageHeadImg
	climbStairs.Rule = ob.Rule
	climbStairs.EndTime = st
	climbStairs.StartTime = et
	climbStairs.Status = ob.Status
	climbStairs.Title = ob.Title
	climbStairs.CreatedAt = nowData
	climbStairs.UpdatedAt = nowData
	if ob.Id!=0 {
		climbStairs.Id = ob.Id
		err := climbStairs.UpdateClimbStairsInfo()
		return err
	}
	//添加
	res := climbStairs.AddClimbStairsInfo()
	return res
}
/**
获取详细信息
*/
func GetClimbStairsInfo(id int64) map[string]interface{}{
	info := make(map[string]interface{})
	var climbStairs models.ClimbStairs
	res := climbStairs.GetInfoClimbStairsById(id)
	if res.Id==0 {
		return nil
	}
	info["id"] = res.Id
	info["rule"] = res.Rule
	info["title"] = res.Title
	info["status"] = res.Status
	info["bgColor"] = res.BgColor
	info["pageHeadImg"] = res.PageHeadImg
	info["startTime"] = res.StartTime
	info["endTime"] = res.EndTime
	return info
}
/**
列表信息
*/
func GetClimbStairsList(cs *SelectStairs) map[string]interface{}{
	info := make(map[string]interface{})
	var climbStairs models.ClimbStairs
	maps := make(map[string]interface{})
	maps["status"] = 1
	fmt.Println(cs)
	res ,count := climbStairs.GetInfoClimbStairsList(cs.PageNum,cs.PageSize,maps)
	info["list"] = res
	info["totalCount"] = count
	info["pageSize"] = cs.PageSize
	info["pageNum"] = cs.PageNum
	info["totalPage"] = math.Ceil(float64(count)/float64(cs.PageSize))
	return info
}
