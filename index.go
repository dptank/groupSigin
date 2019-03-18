package main
//
//import (
//	"fmt"
//	"github.com/go-ini/ini"
//	"os"
//	"strings"
//)
//
//func main() {
//	cfg, err := ini.Load("conf/kafka.ini", "conf/app.ini")
//	if err != nil {
//		fmt.Printf("%v", err)
//		os.Exit(1)
//	}
//
//	// 运行模式
//	mode := cfg.Section("").Key("app_mode").String()
//	host := strings.Split(cfg.Section(mode).Key("kafka.host").String(),",")
//	port := cfg.Section(mode).Key("kafka.port").String()
//	//version := cfg.Section(mode).Key("kafka.version").String()
//	//timeout := cfg.Section(mode).Key("kafka.timeout").String()
//	//timeout := cfg.Section(mode).Key("kafka.timeout").String()
//	var hosts []string
//	for i:=0;i<len(host) ;i++  {
//		hosts = append(hosts,host[i]+":"+port)
//	}
//	fmt.Println(hosts)
//	//brokerList := "127.0.0.1:9092,127.0.0.1:9092"
//	//a := strings.Split(brokerList,",")
//	//fmt.Println(reflect.TypeOf(a))
//	//defer func() {
//	//	fmt.Println("c")
//	//	if err:=recover();err!=nil {
//	//		fmt.Println(err)
//	//	}
//	//	fmt.Println("d")
//	//}()
//	//f()
//}
////func f() {
////	fmt.Println("a")
////	panic(55)
////	fmt.Println("b")
////	fmt.Println("f")
////}