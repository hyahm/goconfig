package goconfig

import (
	"bytes"
	"fmt"
)

// 全部统一 utf-8格式
// 文件操作
// 用于写文件



//
//var fis []*fileinfo
var module_name []byte
var module_filter map[string]bool
var is_module_note bool

const (
	N  = iota
	K
	M
)

func (fl *filelines)readlines() {
	module_filter = make(map[string]bool)
	// 去掉windows换行的\r
	fl.All = bytes.ReplaceAll(fl.All , []byte("\r"), []byte(""))

	lines := bytes.Split(fl.All, []byte("\n"))
	for i:= 0; i < len(lines); i++ {
		// 分类
		classification(lines[i])
	}
	//fl.getKey()
}

func classification(line []byte) {
	// 分大类
	//去掉2边的空格
	line_byte_no_space := bytes.Trim(line, " ")

	// 忽略注释和空行
	if string(line_byte_no_space) == "" {
		return
	}

	// 判断是否是模块的注释
	if string(line_byte_no_space[0:1]) == MODEL_START {
		// 模块
		line_lenth := len(line_byte_no_space)
		if string(line_byte_no_space[line_lenth-1:line_lenth]) == MODEL_END {
			// 模块
			module_name = bytes.Trim(line_byte_no_space[1:line_lenth-1], " ")
			if _, ok := module_filter[string(module_name)]; ok {
				panic(fmt.Sprintf("group %s Repetition", string(module_name)))
			}
			module_filter[string(module_name)] = true
			// 倒叙查找注释
			fl.line = append(fl.line, &filetype{
				Sign:   M,
				Key:    nil,
				Module: module_name,
				Note:   nil,
			})
			//模块直接跳过
			return
		}
	}
	if string(line_byte_no_space[0:1]) == NOTE {
		// nothing to do
		// 注释
		fl.line = append(fl.line, &filetype{
			Sign: N,
			Note: &note{
				Module: module_name,
				Key: nil,
				Value: line_byte_no_space,
			},
		})
		return
	}
	// 如果是key
	name, value := getKeyValue(line_byte_no_space)
	if name != nil {
		fl.line = append(fl.line, &filetype{
			Sign:   K,
			Key: &key{
				Module: module_name,
				Name: name,
				Value: value,
			},
		})
	}


}

// 写入到
func getKeyValue(line []byte) ([]byte, []byte) {
	// 存入值到 configKeyValue， 更新 fis
		index := bytes.Index(line, []byte(SEP))
		if index == -1 {
			panic(fmt.Sprintf("key error, not found =, line: %s", string(line)))
		}
		// 左边的是key， 右边的是值
		key := bytes.Trim(line[:index], " ")
		// 不能包含. 和空格
		if bytes.Contains(key, []byte(" ")) {
			panic(fmt.Sprintf("key error, not allow contain space, key: %s", string(key)))
		}
		if bytes.Contains(key, []byte(".")) {
			panic(fmt.Sprintf("key error, not allow contain point, key: %s", string(key)))
		}
		value := bytes.Trim(line[index+1:], " ")

		if _, ok := fl.configKeyValue[string(key)]; ok {
			// 去掉重复项
			return nil, nil
		}

		fl.configKeyValue[string(key)] = value
		return key, value
}
//
//const BEGIN = 0
//const END = 1


