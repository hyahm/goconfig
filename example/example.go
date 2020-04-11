package main

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/hyahm/goconfig"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 初始化配置文件

	// goconfig.InitWriteConf("test.conf")
	// goconfig.WriteInt("aaa.bbb.ccc", 5)
	// goconfig.FlushWrite()
	if err := goconfig.InitConf("client.ini", goconfig.INI); err != nil {
		log.Fatal(err)
	}

	// goconfig.InitConf("write.conf") // 与InitConf的区别是， 这个会清空里面原有数据
	// goconfig.PrintKeyValue()
	fmt.Println(goconfig.ReadBool("u5.redis_download.on"))
	fmt.Println(goconfig.ReadDuration("server.timeout"))
	fmt.Println(goconfig.ReadString("u5.redis_handle.redis_password"))
	l := make([]string, 0)
	json.Unmarshal(goconfig.ReadBytes("u5.redis_download.pks"), &l)
	goconfig.PrintKeyValue()
	fmt.Println(l)
}
