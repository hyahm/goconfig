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
	"app"
	"github.com/hyahm/goconfig"
	"github.com/hyahm/golog"
)

func main() {
	golog.InitLogger("/data/log", 0, false)
	
  goconfig.InitConf("showgo.conf")
	app.ShowApp()
}

```
### 加载配置文件不启动服务, 只要调用就会更新缓存中的配置
```
goconfig.LoopKey()
```
### 配置文件说明
文件是key = value 的格式, 必须写一行
- [ 和 # 开头的都是注释
- key 不能有 = 符号, 这个是区分key和value的, 会使用第一个 = 号做区分
- map类型 {"key": value} 格式
- key 有重复会报错提示 哪个key 重复了

### 方法
调用了InitConf方法, 任何地方都可以调用下面的方法从配置文件读取, 因为是从缓存中读取, 速度不会慢就是
```
# 返回一个float64类型, 默认为0
func ReadFloat(key string) float64 {   # eg:   a = 433.65,  return 433.65
读取文件的内容 返回一个string, 默认为""
func ReadFile(key string) string {     # eg:   a = /root/.bashrc 
返回一个string,会去掉2边的空格 默认为""
func ReadString(key string) string {   # eg: a = _asdf*^%&    return _asdf*^%&
// 返回int, 默认为0
func ReadInt(key string) int  {        # eg: a = 5
返回int16, 默认为 0
func ReadInt16(key string) int16 {       # eg: a = 5
// 2边需要用到引号, 这个与ReadString的区别是不会去掉2边的空格, 并且要用双引号引起来
func ReadPassword(key string) string {     # eg: a = "5   "  return "5   "    双引号起来是为了标记有几个空格
# 读取bool类型, 默认返回false
func ReadBool(key string) bool {           # eg: a = true
// 返回int64, 默认为0
func ReadInt64(key string) int64 {           # eg: a = 5
// 返回map
func ReadMap(key string) map[string]interface{} {      # eg: a = {"name":"hyahm", "age": 5}
// 返回数字, 类型必须是int, 默认返回[]
func ReadIntArray(key string) []int {           # eg: a = [1,5,23,5,6,34,23,786]
// 返回数字, 类型必须是string, 默认返回[]
func ReadStringArray(key string) []string {     # eg: a = ["aaa","bbb","ccc"]
```
