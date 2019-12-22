package main

import (
	"encoding/json"
	"github.com/hyahm/goconfig"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 初始化配置文件
	//goconfig.InitConf("test.conf")
	goconfig.InitConf("write.conf")
	//uint64 = 9223372036854775808
	//int64 = 123456223
	//int = 1234556223
	//bool = true
	//password = "adlfjlskdf "
	//[one]
	//float = 0.25
	//string = goconfig
	//print(goconfig.ReadString("uint64"))
	// 写入 模块：key, key: name, value: cander, 备注：姓名     的配置文件
	goconfig.WriteString("key.name", "cander", "姓名")
	//// 写入 模块：key, key: name, value: cander, 备注：用户表       的配置文件
	send, _ := json.Marshal(&user{
		Id:1,
		Name: "cander",
		Age: 20,
	})
	goconfig.WriteBytes("user", send, "用户表")
	goconfig.WriteBytes("one.user", send, "用户表")
	goconfig.FlushWrite()
	//goconfig.PrintKeyValue()
	//goconfig.PrintKeyValue()

}
