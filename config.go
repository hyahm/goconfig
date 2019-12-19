package goconfig

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

//const middle = "========="
const SEP = "=" // key 和 value 分隔符

var configPath string // 配置文件路径，保存后方便重新加载配置文件
var file *os.File
var configKeyValue map[string]string

const NOTE = "#"        // #开头的为注释
const MODEL_START = "[" // [开头的为注释
const MODEL_END = "]"   // [开头的为注释

// 读取配置文件

func InitConf(configpath string) {
	configKeyValue = make(map[string]string)
	fptmp := configpath
	fptmp = filepath.Clean(fptmp)
	//判断是相对路径还是绝对路径

	_, err := os.Stat(fptmp)
	if err != nil {
		if err := os.MkdirAll(filepath.Dir(fptmp), 0755); err != nil {
			panic(err)
		}
	}
	file, err = os.OpenFile(fptmp, os.O_CREATE|os.O_RDONLY, 0644)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	configPath = fptmp

	readlines()
}

func Print() {
	if configKeyValue == nil {
		fmt.Println()
	}
	for k, v := range configKeyValue {
		log.Printf("key: %s ---- value: %s \n", k, v)
	}
}

// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
