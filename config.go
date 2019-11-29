package goconfig

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

//const middle = "========="
const SEP = "=" // key 和 value 分隔符

var ConfigPath string // 配置文件路径，保存后方便重新加载配置文件
var ConfigKeyValue map[string]string
var NOTE = "#" // #开头的为注释
var MODEL_START = "[" // [开头的为注释
var MODEL_END = "]" // [开头的为注释

// 读取配置文件

func InitConf(configpath string) {
	fptmp := configpath
	fptmp = filepath.Clean(fptmp)
	//判断是相对路径还是绝对路径

	//这是绝对路径
	_, err := os.Stat(fptmp)
	if err != nil {
		if err := os.MkdirAll(filepath.Dir(fptmp), 0744); err != nil {
			log.Fatal(err)
		}
		os.OpenFile(fptmp, os.O_CREATE, 0644)
	}
	ConfigPath = fptmp
	fmt.Println("configfile:", ConfigPath)

	LoopKey()
}

func Print() {
	if ConfigKeyValue == nil {
		panic("not init")
	}
	for k,v := range ConfigKeyValue {
		log.Printf("key: %s ---- value: %s", k, v)
	}
}

// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
func LoopKey() {
	var err error
	//获取文件字节
	cf, err := ioutil.ReadFile(ConfigPath)
	if err != nil {
		log.Fatal(err)
	}

	line := 0
	bs := make([]string, 0)
	if runtime.GOOS != "windows" {
		bs = strings.Split(string(cf), "\n")
	} else {
		bs = strings.Split(string(cf), "\r\n")
	}

	//换行符切割字符串
	ConfigKeyValue = make(map[string]string, 0)
	modelname := ""
	for i := 0; i < len(bs); i++ {
		//fmt.Println()

		line++

		//去掉2边的空格
		sbs := strings.Trim(bs[i], " ")
		//  #开头是注释， [ 开头是模块 , 空行
		if sbs == "" || sbs[0:1] == NOTE {
			continue
		}
		//模块的话
		if sbs[0:1] == MODEL_START && sbs[len(sbs)-1:] == MODEL_END {
			modelname = sbs[1:len(sbs)-1]
			continue
		}

		index := strings.Index(sbs, SEP)
		key := strings.Trim(sbs[:index], " ")
		if modelname != "" {
			key = fmt.Sprintf("%s.%s",modelname, key)
		}

		if _, ok := ConfigKeyValue[key]; ok {
			log.Fatal(fmt.Sprintf("key:%s duplicate，line:%d \n", key, line))
		} else {
			fmt.Printf("Key:%s --- Value: %s \n", key, strings.Trim(sbs[index+1:], " "))
			ConfigKeyValue[key] = strings.Trim(sbs[index+1:], " ")
		}

	}
}
