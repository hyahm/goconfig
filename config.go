package goconfig

import (
	"io/ioutil"
	"os"
	"path/filepath"
)

//const middle = "========="
const SEP = "=" // key 和 value 分隔符
const NOTE = "#"        // #开头的为注释
const MODEL_START = "[" // [开头的为注释
const MODEL_END = "]"   // [开头的为注释
var fl *filelines
// 读取配置文件

type key struct {
	Module []byte
	Name []byte
	Value []byte
}

type note struct {
	Module []byte // module的注释
	Key []byte   // 某个key的注释
	Value []byte // 注释的值
}

type filetype struct {
	// 每一行存在这三种类型， 空行去掉了
	Sign   int   // 1是 key， 2， 是module   3， note
	Key    *key   // key
	Module []byte   // Module名
	Note   *note   // 注释
}

type filelines struct {
	line []*filetype // 每一行
	filepath string    // 配置文件
	read  *os.File
	write  *os.File
	All []byte  // 文件所有数据
	configKeyValue  map[string][]byte // key, value
}

func InitConf(configpath string) {

	fptmp := filepath.Clean(configpath)
	//判断文件目录是否存在
	_, err := os.Stat(filepath.Dir(fptmp))
	if err != nil {
		// 不存在就先创建目录
		if err := os.MkdirAll(filepath.Dir(fptmp), 0755); err != nil {
			panic(err)
		}
	}
	fl = &filelines{
		filepath: configpath,
		line: make([]*filetype,0),
		configKeyValue: make(map[string][]byte),
	}
	fl.All, err = ioutil.ReadFile(fptmp)
	if err != nil {
		panic(err)
	}

	fl.readlines()
}


// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
