package main

import (
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
	if err := goconfig.InitConf("test.yaml", "yaml"); err != nil {
		log.Fatal(err)
	}

	// goconfig.InitConf("write.conf") // 与InitConf的区别是， 这个会清空里面原有数据
	goconfig.PrintKeyValue()
}
