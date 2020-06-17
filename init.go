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
	Groups   []*groupLine      // 组
	Lines    []*node           // 单key
	Read     []byte            // 文件读出来的所有内容
	Write    []byte            // 文件写的所有内容
	KeyValue map[string]string // 键值缓存， key value的值  key or group.key
	Filepath string            // 配置文件路径
	sjson    map[string]interface{}
}

var kvconfig *config
var Deep = 3
var tp typ

type typ int

const (
	INI typ = iota
	JSON
	YAML
)

func (t typ) String() string {
	switch t {
	case 0:
		return "ini"
	case 1:
		return "json"
	case 2:
		return "yaml"
	default:
		return "ini"
	}
}

// 读取配置文件到全局变量，并检查重复项, 重载配置文件执行这个函数
func Reload() error {
	kvconfig = nil

	file := kvconfig.Filepath
	if file == "" {
		return errors.New("not found config")
	}
	//判断文件目录是否存在
	_, err := os.Stat(kvconfig.Filepath)
	if err != nil {
		// 不存在就先创建目录
		return err

	}

	// 清空数据
	// 检查是否有错
	tmp := &config{
		Filepath: file,
		Lines:    make([]*node, 0),
		KeyValue: make(map[string]string),
	}
	tmp.Read, err = ioutil.ReadFile(file)
	if err != nil {
		return err
	}
	switch tp {
	case JSON:
		if err := tmp.readJson(); err != nil {
			return err
		}
	case YAML:
		if err := tmp.readYaml(); err != nil {
			return err
		}
	default:
		if err := tmp.readIni(); err != nil {
			return err
		}
	}
	kvconfig = tmp
	// 更新值
	return nil
}

func InitConf(path string, t typ) error {
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
	kvconfig = &config{
		Filepath: fptmp,
		Lines:    make([]*node, 0),
		KeyValue: make(map[string]string),
	}
	kvconfig.Read, err = ioutil.ReadFile(fptmp)
	if err != nil {
		return err
	}
	switch t {
	case JSON:
		tp = t
		if err := kvconfig.readJson(); err != nil {
			return err
		}
	case YAML:
		tp = t
		if err := kvconfig.readYaml(); err != nil {
			return err
		}
	default:
		if err := kvconfig.readIni(); err != nil {
			return err
		}
	}

	return nil
}

// 从bytes 解析， Reload方法
func InitFromBytes(data []byte, t typ) error {

	kvconfig = &config{
		Lines:    make([]*node, 0),
		KeyValue: make(map[string]string),
	}
	kvconfig.Read = data
	switch t {
	case JSON:
		tp = t
		if err := kvconfig.readJson(); err != nil {
			return err
		}
	case YAML:
		tp = t
		if err := kvconfig.readYaml(); err != nil {
			return err
		}
	default:
		if err := kvconfig.readIni(); err != nil {
			return err
		}
	}
	return nil
}

// 初始化写入文件的方法， 会清空内容
func InitWriteConf(configpath string, t typ) {

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

	switch t {
	case JSON:
		tp = t
	case YAML:
		tp = t
	default:

	}
	kvconfig = &config{
		Filepath: configpath,
		Lines:    make([]*node, 0),
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
