package main

import (
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

type ttt struct {
	key   string
	value string
	note  []string
}

func main() {
	// 初始化配置文件

	goconfig.InitConf("client.ini", goconfig.INI)

	// goconfig.WriteInt("aaa.bbb.ccc", 5)

	// // goconfig.FlushWrite()
	// if err := goconfig.InitConf("client.ini", goconfig.INI); err != nil {
	// 	log.Fatal(err)
	// }
	// u5 := make([]*user, 0)
	// err := goconfig.ReadArrayFromNode("u5", &u5)
	// if err != nil {
	// 	golog.Error(err)
	// }
	// for _, v := range u5 {
	// 	fmt.Printf("%+v", v)
	// }
	// fmt.Println(u5)
	// // goconfig.InitConf("write.conf") // 与InitConf的区别是， 这个会清空里面原有数据
	// // goconfig.PrintKeyValue()
	// fmt.Println(goconfig.ReadBool("u5.redis_download.on"))
	// fmt.Println(goconfig.ReadDuration("server.timeout"))
	// fmt.Println(goconfig.ReadString("u5.redis_handle.redis_password"))
	// l := make([]string, 0)
	// json.Unmarshal(goconfig.ReadBytes("u5.redis_download.pks"), &l)
	// goconfig.PrintKeyValue()
	// fmt.Println(l)
	// log.Fatal(goconfig.Start())
}
