package climbStairs

import (
	"groupSigin/models"
	"time"
	"math"
	"fmt"
	"groupSigin/service"
)
//活动规则
type ClimbStairsInfo struct {
	Id int64 `json:"id"`
	Rule string `json:"rule" binding:"required"`
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
//活动商品信息
type ClimbStairsItems struct {
	TopicId int64 `json:"topicId"`
	NeedNum int64 `json:"needNum"`
	TopicTitle string `json:"topicTitle"`
	TopicImg string `json:"topicImg"`
	ReceivedNum int64 `json:"receivedNum"`
	DiffNum int64 `json:"diffNum"`
}
//首页返回值
type ClimbStairsInit struct {
	Id int64 `json:"id"`
	Rule string `json:"rule" `
	Title string `json:"title"`
	Status int `json:"status"`
	BgColor string `json:"bgColor"`
	PageHeadImg string `json:"pageHeadImg"`
	StartTime string `json:"startTime"`
	EndTime string `json:"endTime"`
	//Info *ClimbStairsInfo `json:"-"`
	ItemInfo []*ClimbStairsItems `json:"itemInfo"`
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
func GetClimbStairsInfo(id int64) *ClimbStairsInfo{
	var climbStairs models.ClimbStairs
	res := climbStairs.GetInfoClimbStairsById(id)
	if res.Id==0 {
		return nil
	}
	info := &ClimbStairsInfo{
		Id:res.Id,
		Rule:res.Rule,
		Title:res.Title,
		Status:res.Status,
		BgColor:res.BgColor,
		PageHeadImg:res.PageHeadImg,
		StartTime:res.StartTime.Format("2006-01-02 15:04:05"),
		EndTime:res.EndTime.Format("2006-01-02 15:04:05"),
	}

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
	//fmt.Println(cs)
	res ,count := climbStairs.GetInfoClimbStairsList(cs.PageNum,cs.PageSize,maps)
	var list []*ClimbStairsInfo
	for _,value := range res{
		list=append(list,&ClimbStairsInfo{
			Id:value.Id,
			Rule:value.Rule,
			Title:value.Title,
			Status:value.Status,
			BgColor:value.BgColor,
			PageHeadImg:value.PageHeadImg,
			StartTime:value.StartTime.Format("2006-01-02 15:04:05"),
			EndTime:value.EndTime.Format("2006-01-02 15:04:05"),
		})
	}
	info["list"] = list
	info["totalCount"] = count
	info["pageSize"] = cs.PageSize
	info["pageNum"] = cs.PageNum
	info["totalPage"] = math.Ceil(float64(count)/float64(cs.PageSize))
	return info
}
/**
初始化活动
*/
func GetClimbStairsInit(id int64) *ClimbStairsInit{
	//活动信息
	climbStairsInfo := GetClimbStairsInfo(id)
	//商品信息
	activityId := climbStairsInfo.Id
	items := GetClimbStairsItemInfoByActId(activityId)
	if items==nil {
		return nil
	}
	var topic []int64
	for _,value:=range items{
		topic = append(topic,value.TopicId)
	}
	var climbStairsItemsList []*ClimbStairsItems
	topics :=  service.GetTopicInfoByIds(topic)
	//fmt.Println(topics[1001].TopicTitle)
	for _,value:=range items{
		fmt.Println(topics[int(value.TopicId)] )
		climbStairsItemsList = append(climbStairsItemsList,&ClimbStairsItems{
			TopicId:value.TopicId,
			NeedNum:value.NeedNum,
			TopicImg:topics[int(value.TopicId)].TopicImage,
			TopicTitle:topics[int(value.TopicId)].TopicTitle,
		})
	}
	indexInfo := &ClimbStairsInit{
		Id:climbStairsInfo.Id,
		Rule:climbStairsInfo.Rule,
		Title:climbStairsInfo.Title,
		Status:climbStairsInfo.Status,
		BgColor:climbStairsInfo.BgColor,
		PageHeadImg:climbStairsInfo.PageHeadImg,
		StartTime:climbStairsInfo.StartTime,
		EndTime:climbStairsInfo.EndTime,
		ItemInfo:climbStairsItemsList,
	}
	return indexInfo
}
