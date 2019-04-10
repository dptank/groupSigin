/*
@Time : 19/4/2 下午6:17 
@Author : gongjiapeng
@File : commService.go
@Software: GoLand
*/
package service

import (
	"groupSigin/pkg/configs"
	"encoding/json"
	"groupSigin/controllers"
	"bytes"
	"log"
)

/**
获取topic详细信息
*/
type Topic struct {
	TopicIds []int64 `json:"topicIds"`
}
type TopicInfo struct {
	Data Lists `json:"data"`
	controllers.Res
}
type Lists struct {
	List map[int]*TopicList `json:"list"`
}
type TopicList struct {
	TopicImage string `json:"topicImage"`
	TopicTitle string `json:"topicTitle"`
}
func GetTopicInfoByIds(topicId []int64) map[int]*TopicList{
	url := configs.GetConfig("topic_url")
	topicIds := &Topic{TopicIds:topicId}
	jsonTopicIds ,err:=json.Marshal(topicIds)
	if err!=nil {
		return nil
	}
	body := bytes.NewReader(jsonTopicIds)
	content ,err := controllers.HttpPost(url,body)
	if err != nil {
		log.Println("Post failed:", err)
		return nil
	}
	res := &TopicInfo{}
	err = json.Unmarshal(content, res)
	//fmt.Println(res)
	return res.Data.List
}