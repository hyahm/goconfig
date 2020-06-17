package goconfig

import (
	"bytes"
	"errors"
	"fmt"
	"strings"
)

// 全部统一 utf-8格式
// 文件操作
// 用于写文件
// 读取ini文件

//
//var fis []*fileinfo
var module_name string
var module_filter map[string]bool
var notes []string

func (c *config) readIni() error {
	// 去掉windows换行的\r
	c.Read = bytes.ReplaceAll(c.Read, []byte("\r"), []byte(""))

	lines := bytes.Split(c.Read, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		// 分类
		if err := classification(lines[i]); err != nil {
			return err
		}
	}
	return nil
}

func classification(line []byte) error {
	// 分大类
	//去掉2边的空格
	module_filter = make(map[string]bool)
	line_byte_no_space := bytes.Trim(line, " ")
	// 忽略注释和空行
	if string(line_byte_no_space) == "" {
		return nil
	}

	//判断是否是组

	line_lenth := len(line_byte_no_space)
	if string(line_byte_no_space[0:1]) == MODEL_START && string(line_byte_no_space[line_lenth-1:line_lenth]) == MODEL_END {
		// 模块
		module_name = string(bytes.Trim(line_byte_no_space[1:line_lenth-1], " "))
		if _, ok := module_filter[string(module_name)]; ok {
			return errors.New(fmt.Sprintf("group %s Repetition", string(module_name)))
		}
		kvconfig.newGroup(string(module_name), notes...)
		notes = nil
		module_filter[string(module_name)] = true
		return nil

	}
	if string(module_name) == "" {
		//判断是否是注释

		if strings.ContainsAny(NOTE, string(line_byte_no_space[0:1])) {
			// 注释
			notes = append(notes, string(line_byte_no_space))
			return nil
		}
		// 添加kv
		k, v, err := getKeyValue(line_byte_no_space)
		if err != nil {
			return err
		}
		kvconfig.newKeyValue(k, string(v), notes...)
		notes = nil
	} else {
		// 组
		if strings.ContainsAny(NOTE, string(line_byte_no_space[0:1])) {
			// 注释
			notes = append(notes, string(line_byte_no_space))
			return nil
		}
		k, v, err := getKeyValue(line_byte_no_space)
		if err != nil {
			return err
		}
		for i, g := range kvconfig.Groups {
			//在组里面就添加
			if string(g.name) == string(module_name) {
				kvconfig.addGroupKeyValue(i, k, string(v), notes...)
				notes = nil
				return nil
			}
		}

	}
	return nil
}

// 写入到
func getKeyValue(line []byte) (string, []byte, error) {
	// 存入值到 configKeyValue， 更新 fis
	index := bytes.Index(line, []byte(SEP))
	if index == -1 {
		return "", nil, errors.New(fmt.Sprintf("key error, not found =, line: %s", string(line)))
	}
	// 左边的是key， 右边的是值
	key := bytes.Trim(line[:index], " ")
	// 不能包含. 和空格
	if bytes.Contains(key, []byte(" ")) {
		return "", nil, errors.New(fmt.Sprintf("key error, not allow contain space, key: %s", string(key)))
	}
	if bytes.Contains(key, []byte(".")) {
		return "", nil, errors.New(fmt.Sprintf("key error, not allow contain point, key: %s", string(key)))
	}
	// value 去掉2边空格
	value := bytes.Trim(line[index+1:], " ")
	k := string(key)
	if string(module_name) != "" {
		k = string(module_name) + "." + k
	}
	if _, ok := kvconfig.KeyValue[k]; ok {
		// 去掉重复项
		return "", nil, errors.New(fmt.Sprintf("key duplicate, key: %s", string(key)))
	}
	kvconfig.KeyValue[k] = string(value)
	return string(key), value, nil
}
