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
var notes [][]byte

func (fl *config) readlines() {
	module_filter = make(map[string]bool)
	// 去掉windows换行的\r
	fl.Read = bytes.ReplaceAll(fl.Read, []byte("\r"), []byte(""))

	lines := bytes.Split(fl.Read, []byte("\n"))
	for i := 0; i < len(lines); i++ {
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

	//判断是否是组
	if string(line_byte_no_space[0:1]) == MODEL_START {
		// 模块
		line_lenth := len(line_byte_no_space)
		if string(line_byte_no_space[line_lenth-1:line_lenth]) == MODEL_END {
			// 模块
			module_name = bytes.Trim(line_byte_no_space[1:line_lenth-1], " ")
			if _, ok := module_filter[string(module_name)]; ok {
				panic(fmt.Sprintf("group %s Repetition", string(module_name)))
			}
			fl.newGroup(module_name, notes...)
			notes = nil
			module_filter[string(module_name)] = true
			return
		}
	}
	if string(module_name) == "" {
		//判断是否是注释
		if string(line_byte_no_space[0:1]) == NOTE {
			// 注释
			notes = append(notes, line_byte_no_space)
			return
		}
		// 添加
		k, v := getKeyValue(line_byte_no_space)
		fl.newKeyValue(k, v, notes...)
		notes = nil
	} else {
		// 组
		if string(line_byte_no_space[0:1]) == NOTE {
			// 注释
			notes = append(notes, line_byte_no_space)
			return
		}
		k, v := getKeyValue(line_byte_no_space)
		for i, g := range fl.Groups {
			//在组里面就添加
			if string(g.name) == string(module_name) {
				fl.addGroupKeyValue(i, k, v, notes...)
				notes = nil
				return
			}
		}
		panic("something error")
	}
}

// 写入到
func getKeyValue(line []byte) (string, []byte) {
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

	if _, ok := fl.KeyValue[string(key)]; ok {
		// 去掉重复项
		return "", nil
	}

	fl.KeyValue[string(key)] = value
	return string(key), value
}
