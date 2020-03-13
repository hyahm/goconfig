package goconfig

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"sigs.k8s.io/yaml"
)

//const middle = "========="
var (
	SEP         = "="  // key 和 value 分隔符
	NOTE        = "#;" // #开头的为注释
	MODEL_START = "["  // 模块开头符号
	MODEL_END   = "]"  // 模块结尾符号
	WELL        = "#"  // 写入的注释
)

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

// 这个是ini的config
type config struct {
	Groups   []groupLine       // 组
	Lines    []node            // 单key
	Read     []byte            // 文件读出来的所有内容
	Write    []byte            // 文件写的所有内容
	KeyValue map[string]string // 键值缓存， key的值  key or group.key
	Filepath string            // 配置文件路径
	sjson    map[string]interface{}
}

var Config *config
var tp string = "ini"

// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
func Reload() error {
	Config = nil

	file := Config.Filepath
	if file == "" {
		return errors.New("not found config")
	}
	//判断文件目录是否存在
	_, err := os.Stat(Config.Filepath)
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
	if err := tmp.readIni(); err != nil {
		return err
	}
	// 更新值
	Config = nil
	Config = tmp
	return nil
}

func InitConf(path string, typ ...string) error {
	fptmp := filepath.Clean(path)

	//判断文件目录是否存在
	_, err := os.Stat(filepath.Dir(fptmp))
	if err != nil {
		// 不存在就先创建目录
		if err := os.MkdirAll(filepath.Dir(fptmp), 0755); err != nil {
			return err
		}

	}
	// 创建文件
	os.OpenFile(fptmp, os.O_CREATE, 0644)
	Config = &config{
		Filepath: fptmp,
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}
	Config.Read, err = ioutil.ReadFile(fptmp)
	if err != nil {
		return err
	}
	if len(typ) > 0 {
		switch typ[0] {
		case "json":
			tp = "json"
			if err := Config.readJson(); err != nil {
				return err
			}
		case "yaml":
			tp = "yaml"
			j, err := yaml.YAMLToJSON(Config.Read)
			if err != nil {
				return err
			}
			fmt.Println("convent to json")
			Config.Read = j
			fmt.Println(string(j))
			if err := Config.readJson(); err != nil {
				return err
			}
		default:
			if err := Config.readIni(); err != nil {
				return err
			}
		}
	}
	return nil
}

// 从bytes 解析， Reload方法
func InitFromBytes(data []byte, typ ...string) error {

	Config = &config{
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}
	Config.Read = data
	if len(typ) > 0 {
		switch typ[0] {
		case "json":
			tp = "json"
		case "yaml":
			tp = "yaml"
			j, err := yaml.JSONToYAML(Config.Read)
			if err != nil {
				return err
			}
			Config.Read = j

		default:
			if err := Config.readIni(); err != nil {
				return err
			}
		}
	}
	return nil
}

// 初始化写入文件的方法， 会清空内容
func InitWriteConf(configpath string, typ ...string) {

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

	if len(typ) > 0 {
		switch typ[0] {
		case "json":
			tp = "json"
		case "yaml":
			tp = "json"
		default:
		}
	}

	Config = &config{
		Filepath: configpath,
		Lines:    make([]node, 0),
		KeyValue: make(map[string]string),
	}

}

//
// func InitBytes() {
// 	Config = &config{
// 		Filepath: "",
// 		Lines:    make([]node, 0),
// 		KeyValue: make(map[string]string),
// 	}
// }
