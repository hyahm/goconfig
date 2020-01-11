package main

import (
	"github.com/hyahm/goconfig"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// 初始化配置文件
	goconfig.InitConf("write.conf") // 与InitConf的区别是， 这个会清空里面原有数据
	goconfig.PrintKeyValue()
}
