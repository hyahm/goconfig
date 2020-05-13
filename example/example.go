package main

import (
	"fmt"
	"log"

	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

type user struct {
	On             bool   `json:"on"`
	Key            string `json:"key"`
	redis_host     string `json:"redis_host"`
	redis_password string `json:"redis_password"`
	redis_db       int    `json:"redis_db"`
}

func main() {
	// 初始化配置文件

	// goconfig.InitWriteConf("test.conf")
	// goconfig.WriteInt("aaa.bbb.ccc", 5)
	// goconfig.FlushWrite()
	if err := goconfig.InitConf("client.ini", goconfig.INI); err != nil {
		log.Fatal(err)
	}
	u5 := make([]*user, 0)
	err := goconfig.ReadArrayBytes("u5", &u5)
	if err != nil {
		golog.Error(err)
	}
	for _, v := range u5 {
		fmt.Printf("%+v", v)
	}
	fmt.Println(u5)
	// // goconfig.InitConf("write.conf") // 与InitConf的区别是， 这个会清空里面原有数据
	// // goconfig.PrintKeyValue()
	// fmt.Println(goconfig.ReadBool("u5.redis_download.on"))
	// fmt.Println(goconfig.ReadDuration("server.timeout"))
	// fmt.Println(goconfig.ReadString("u5.redis_handle.redis_password"))
	// l := make([]string, 0)
	// json.Unmarshal(goconfig.ReadBytes("u5.redis_download.pks"), &l)
	// goconfig.PrintKeyValue()
	// fmt.Println(l)
}
