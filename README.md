# goconfig
read go config file

### 安装
```
go get github.com/hyahm/goconfig
```
### 使用
指定配置文件路径
goconfig.InitConf(path string) 指定配置文件路径, 如果没有配置文件会生成空的配置文件, 读取的配置文件读取至缓存中
```
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

```
get(或者数值)
set（设置数值， 可以添加注释）
getset（设置获取）
