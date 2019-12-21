package main

import (
	"fmt"
	"github.com/hyahm/goconfig"
)

type user struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {

	goconfig.InitConf("test.conf")

	fmt.Println(goconfig.WriteString("key.name", "liubiao", "姓名"))
	fmt.Println(goconfig.ReadString("listen"))
}
