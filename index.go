package main

import (
	"encoding/json"
	"fmt"
	"bytes"
	"io"
	"net/http"
	"io/ioutil"
)

func main() {
	url := "http://jiapeng.express.com/api/service/glove/list"
	song := make(map[string]interface{})
	song["mailNo"] = "251602300981"
	song["reqNo"] = "SF5BF7F60A39A2E"
	bytesData, err := json.Marshal(song)
	//fmt.Println(bytesData)
	if err != nil {
		fmt.Println(err.Error() )
		return
	}

	reader := bytes.NewReader(bytesData)
	res,err := HttpPost(url,reader)
	var user map[string]interface{}
	json.Unmarshal(res, &user)
	data := user["data"].(map[string]interface{})
	//list  := (data["list"]).(map[string]interface{})
	fmt.Printf("%T %+v", data, data)
	//for key,value := range list{
	//	fmt.Println(value)
	//	fmt.Println(key)
	//}
	//str := (*string)(unsafe.Pointer(&res))
	//s := make(map[string]interface{})
	//fmt.Println(data)
}
func HttpPost(url string,data io.Reader) (res []byte, e error) {
	request, err := http.NewRequest("POST", url, data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	request.Header.Set("Content-Type", "application/json;charset=UTF-8")
	client := http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	respBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//byte数组直接转成string，优化内存
	//str := (*string)(unsafe.Pointer(&respBytes))
	//fmt.Println(*str)
	return respBytes ,nil
	//resp ,err := http.Post(url,"application/json;charset=UTF-8",data)
	//fmt.Println(err)
	//if err != nil {
	//	fmt.Println(err)
	//	panic(err)
	//}
	//defer resp.Body.Close()
	//body, _ := ioutil.ReadAll(resp.Body)
	//fmt.Println(body)
	//return body, nil
}
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