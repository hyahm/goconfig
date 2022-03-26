package main

import (
	"fmt"

	"github.com/hyahm/goconfig"
)

type p struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type ftp struct {
	On      bool   `json:"on"`
	Key     string `json:"key"`
	FtpRoot string `json:"ftp_root"`
}

type user struct {
	On            bool     `json:"on"`
	Key           string   `json:"key"`
	RedisHost     string   `json:"redis_host"`
	RedisPassword string   `json:"redis_password"`
	RedisDb       int      `json:"redis_db"`
	People        *p       `json:"people"`
	Kps           []string `json:"kps"`
}

func main() {
	// 初始化配置文件
	goconfig.InitConf("client.ini")
	m := goconfig.ReadBool("u5.redis_download.key")
	goconfig.Reload()
	fmt.Println(m)
	fmt.Println(goconfig.ReadInt("bbb"))
}
