package gokafka

import (
	"github.com/Shopify/sarama"
	"time"
	"strings"
	"groupSigin/pkg/ginlog"
	"github.com/go-ini/ini"
	"github.com/bsm/sarama-cluster"
)

/**
*kafka 发送消息
*topics  主题
*str 消息体
*/
func SendToKafka(topics string,str string) bool {
	//初始化配置
	cfg, err := ini.Load("conf/kafka.ini", "conf/app.ini")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err:=recover();err!=nil {
			ginlog.LogPrint("初始化配置失败",err)
		}
	}()
	//运行模式
	mode := cfg.Section("").Key("app_mode").String()
	//host
	host := strings.Split(cfg.Section(mode).Key("kafka.host").String(),",")
	//端口
	port := cfg.Section(mode).Key("kafka.port").String()
	//超时时间
	//timeSecond ,err:= cfg.Section(mode).Key("kafka.produce.timeout").Int64()
	//机器列表
	var brokerList []string
	for i:=0;i<len(host) ;i++  {
		brokerList = append(brokerList,host[i]+":"+port)
	}
	//fmt.Println("时间eddd：",(time.Duration(timeSecond)))
	//过期时间10秒
	timeOut := time.Second*10
	producer := getProducer(brokerList,timeOut)
	if producer == nil{
		ginlog.LogPrint("no find producer")
		return false
	}
	msg := &sarama.ProducerMessage{
		Topic: topics,
		Value: sarama.ByteEncoder(str),
	}
	ginlog.LogPrint("send message:",msg)
	producer.Input() <- msg
	return true
}

/**
*获取生产者
*/
func getProducer(brokerList []string,timeout time.Duration) sarama.AsyncProducer{
	config := sarama.NewConfig()
	config.Producer.Return.Successes= true
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	config.Producer.Return.Errors = true
	config.Producer.Timeout = timeout
	config.Version = sarama.V2_1_0_0
	//创建生产者
	producer ,err := sarama.NewAsyncProducer(brokerList,config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err:=recover();err!=nil {
			ginlog.LogPrint("创建生产者失败",err)
		}
	}()
	return producer
}
/**
*获取消费者实类
*/
func GetConsoumer(topics []string,consumerGroup string) *cluster.Consumer{
	//初始化配置
	cfg, err := ini.Load("conf/kafka.ini", "conf/app.ini")
	if err != nil {
		panic(err)
	}
	defer func() {
		if err:=recover();err!=nil {
			ginlog.LogPrint("初始化配置失败",err)
		}
	}()
	//运行模式
	mode := cfg.Section("").Key("app_mode").String()
	//host
	host := strings.Split(cfg.Section(mode).Key("kafka.host").String(),",")
	//端口
	port := cfg.Section(mode).Key("kafka.port").String()
	//消费者组
	group := cfg.Section(mode).Key("kafka.consumer.group").String()
	if consumerGroup == "" {
		consumerGroup = group
	}
	//机器列表
	var brokerList []string
	for i:=0;i<len(host) ;i++  {
		brokerList = append(brokerList,host[i]+":"+port)
	}
	config := cluster.NewConfig()
	config.Consumer.Return.Errors = true
	config.Group.Return.Notifications = true
	consumer ,err := cluster.NewConsumer(brokerList,consumerGroup, topics, config)
	if err != nil {
		panic(err)
	}
	defer func() {
		if err:=recover();err!=nil {
			ginlog.LogPrint("创建消费者者失败",err)
		}
	}()
	go func() {
		for ntf := range consumer.Notifications(){
			ginlog.LogPrint("ntf:",ntf)
		}
	}()
	return consumer
}