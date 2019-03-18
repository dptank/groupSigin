package test

import (
	"github.com/gin-gonic/gin"
	"net/http"
	app2 "groupSigin/pkg/app"
	"groupSigin/pkg/ginlog"
	"fmt"
	"github.com/Shopify/sarama"
	"github.com/bsm/sarama-cluster"
	"os"
	"os/signal"
	"groupSigin/pkg/gokafka"
)

type TestLogs struct {
	UserName string `json:"user_name"`
	UserAddr string `json:"user_addr"`
	UserTel int64 `json:"user_tel"`
}
/**
基础用法
*/
func TestLog(ctx *gin.Context)  {
	app := app2.Gin{C:ctx}
	test := TestLogs{UserName:"cesi",UserAddr:"ceshi2",UserTel:18734922837}
	//test := "我的日志"
	//a ,_ := json.Marshal(test)
	ginlog.LogPrint("TestLogs的值：",test)
	ginlog.LogPrint("TestLogs的值：","我的测试")
	app.Response(http.StatusBadRequest,200,true,"日志存入成功")
	return
}
/**
生产者
*/
func TestKafka(ctx *gin.Context){
	app := app2.Gin{C:ctx}
	//fmt.Printf("producer_test\n")
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Version = sarama.V2_1_0_0
	producer ,err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"},config)
	if err !=nil {
		fmt.Printf("生产者创建失败 producer error :%s\n", err.Error())
		return
	}
	defer producer.Close()
	msg := &sarama.ProducerMessage{
		Topic:"wwwww",
		Key:sarama.StringEncoder("peng"),
	}
	value:="这个一个信息"
	data := make(map[string]interface{})
	//for {
		//fmt.Println(&value)
		msg.Value=sarama.ByteEncoder(value)
		//fmt.Printf("input [%s]\n", value)
		partition,offset,err := producer.SendMessage(msg)
		data["partition"] = partition
		data["offset"] = offset
		if err!=nil {
			fmt.Println("发送失败",err.Error())
		}
	//}
	app.Response(http.StatusBadRequest,200,true,data)
	return
}
/**
消费者
*/
func TestConsoumer(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	fmt.Printf("consumer_test\n")
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	brokers := []string{"127.0.0.1:9092"}
	topics := []string{"express", "wwwww"}
	consumer ,err := cluster.NewConsumer(brokers,"peng", topics, config)
	if err != nil {
		panic(err)
	}
	defer consumer.Close()
	signals := make(chan os.Signal,1)
	signal.Notify(signals,os.Interrupt)

	for {
		select {
		case msg, ok := <-consumer.Messages():
			fmt.Println("你好",ok)
			fmt.Println("你好a",msg)
			if ok {
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				//fmt.Println(os.Stdout,"信息：",msg.Value)
				//fmt.Fprintf(os.Stdout, "%s/%d/%d\t%s\t%s\n", msg.Topic, msg.Partition, msg.Offset, msg.Key, msg.Value)
				consumer.MarkOffset(msg, "")	// mark message as processed
			}
		case <-signals:
			return
		}
	}
	app.Response(http.StatusBadRequest,200,true,"")
	return
}
/**
*卡夫卡生产者
*/
func KafkaTest(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	topics := "wwwww"
	msg := "测试nihaoledd"
	res := gokafka.SendToKafka(topics,msg)
	app.Response(http.StatusBadRequest,200,true,res)
	return
}
/**
*卡夫卡消费者
*/
func KafkaConsoumer(ctx *gin.Context) {
	app := app2.Gin{C:ctx}
	topics := []string{"express", "wwwww"}
	groupId := "peng"
	consumer :=  gokafka.GetConsoumer(topics,groupId)
	defer consumer.Close()
	signals := make(chan os.Signal,1)
	signal.Notify(signals,os.Interrupt)
	for {
		select {
		case msg, ok := <-consumer.Messages():
			if ok {
				//ginlog.LogPrint("offset  partition timestamp  value ",msg.Offset,msg.Partition,msg.Timestamp.String(), string(msg.Value))
				fmt.Printf("msg offset: %d, partition: %d, timestamp: %s, value: %s\n",
					msg.Offset, msg.Partition, msg.Timestamp.String(), string(msg.Value))
				consumer.MarkOffset(msg, "")	// mark message as processed
			}
		case <-signals:
			return
		}
	}
	app.Response(http.StatusBadRequest,200,true,"")
	return
}