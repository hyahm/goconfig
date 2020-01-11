package goconfig

import (
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
)

//const middle = "========="
var (
	SEP         = "="  // key 和 value 分隔符
	NOTE        = "#;" // #开头的为注释
	MODEL_START = "["  // 模块开头符号
	MODEL_END   = "]"  // 模块结尾符号
	WELL        = "#"  // 写入的注释
)

// 读取配置文件

type node struct {
	key   string
	value string
	note  []string
}

type groupLine struct {
	group []node   // 组的行
	note  []string // 组注释
	name  string   // 组名
}

type config struct {
	Groups   []groupLine       // 组
	Lines    []node            // 单key
	Read     []byte            // 文件读出来的所有内容
	Write    []byte            // 文件写的所有内容
	KeyValue map[string]string // 键值缓存， key的值  key or group.key
	Filepath string            // 配置文件路径
}

var fl *config

// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
func Reload() error {
	notes = nil

	file := fl.Filepath
	if file == "" {
		return errors.New("not found config")
	}
	//判断文件目录是否存在
	_, err := os.Stat(fl.Filepath)
	if err != nil {
		// 不存在就先创建目录
		return err

	}
	// 清空数据
	// 检查是否有错
	tmp := &config{
		Filepath: file,
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}
	tmp.Read, err = ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	if err := tmp.readlines(); err != nil {
		return err
	}
	// 更新值
	fl = nil
	fl = tmp
	return nil
}

func InitConf(path string) {
	notes = make([]string, 0)
	fptmp := filepath.Clean(path)
	//判断文件目录是否存在
	_, err := os.Stat(filepath.Dir(fptmp))
	if err != nil {
		// 不存在就先创建目录
		if err := os.MkdirAll(filepath.Dir(fptmp), 0755); err != nil {
			panic(err)
		}

	}
	// 创建文件
	os.OpenFile(fptmp, os.O_CREATE, 0644)
	fl = &config{
		Filepath: fptmp,
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}
	fl.Read, err = ioutil.ReadFile(fptmp)
	if err != nil {
		panic(err)
	}

	if err := fl.readlines(); err != nil {
		panic(err)
	}
}

// 从bytes 解析， 不支持Reload方法
func InitFromBytes(data []byte) {
	notes = make([]string, 0)

	fl = &config{
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}
	fl.Read = data

	if err := fl.readlines(); err != nil {
		panic(err)
	}
}

func InitWriteConf(configpath string) {
	notes = make([]string, 0)
	fptmp := filepath.Clean(configpath)
	//判断文件目录是否存在
	_, err := os.Stat(filepath.Dir(fptmp))
	if err != nil {
		// 不存在就先创建目录
		if err := os.MkdirAll(filepath.Dir(fptmp), 0755); err != nil {
			panic(err)
		}

	}
	os.Remove(fptmp)

	fl = &config{
		Filepath: configpath,
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}

}

func InitBytes() {
	notes = make([]string, 0)
	fl = &config{
		Filepath: "",
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}
}
