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
	u := &user{
		Id:   10,
		Name: "name",
		Age:  10,
	}
	goconfig.InitConf("test.conf")
	goconfig.GetSetString("key.name", "cander")
	goconfig.GetSetFloat("key.weigth", 0.64)
	goconfig.GetSetString("listen", ":5000")
	goconfig.GetSetString("password", ":98895000")
	goconfig.GetSetJson("user", u)
	fmt.Println(goconfig.GetString("key.name"))
	fmt.Println(goconfig.GetString("listen"))
}
