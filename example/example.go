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

	fmt.Println(goconfig.ReadString("key.crt"))
	fmt.Println(goconfig.ReadString("key.key"))
	fmt.Println(goconfig.ReadString("listen"))
}
