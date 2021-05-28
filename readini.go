package goconfig

import (
	"bytes"
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

// 模块去重的
var module_filter map[string]bool

// 注释变量
var notes []string

func (c *config) readIni() error {
	// 去掉windows换行的\r
	c.Read = bytes.ReplaceAll(c.Read, []byte("\r"), []byte(""))
	lines := bytes.Split(c.Read, []byte("\n"))
	for i := 0; i < len(lines); i++ {
		// 处理每一行的信息
		if err := classification(lines[i], i); err != nil {
			return err
		}
	}
	return nil
}

func classification(line []byte, i int) error {
	// 分类 注释还是组名还是键值对
	//去掉2边的空格
	module_filter = make(map[string]bool)
	line_byte_no_space := bytes.Trim(line, " ")
	// 忽略空行
	if string(line_byte_no_space) == "" {
		return nil
	}

	//判断是否是组
	line_lenth := len(line_byte_no_space)
	if string(line_byte_no_space[0:1]) == MODEL_START && string(line_byte_no_space[line_lenth-1:line_lenth]) == MODEL_END {
		// 模块名
		module_name = string(bytes.Trim(line_byte_no_space[1:line_lenth-1], " "))
		if _, ok := module_filter[string(module_name)]; ok {
			return fmt.Errorf("group %s Repetition, line: %d", string(module_name), i)
		}
		kvconfig.newGroup(string(module_name), notes...)
		notes = nil
		module_filter[string(module_name)] = true
		return nil

	}
	// 判断是否是注释
	if strings.ContainsAny(NOTE, string(line_byte_no_space[0:1])) {
		// 注释
		notes = append(notes, string(line_byte_no_space))
		return nil
	}
	k, v, err := getKeyValue(line_byte_no_space, i)
	if err != nil {
		return err
	}
	if string(module_name) == "" {
		kvconfig.newKeyValue(k, string(v), notes...)
		notes = nil
	} else {
		// 组
		for i, g := range kvconfig.Groups {
			//在组里面就添加,
			if string(g.name) == string(module_name) {
				kvconfig.addGroupKeyValue(i, k, string(v), notes...)
				notes = nil
				return nil
			}
		}

	}
	return nil
}

// 获取kv并存入内存
func getKeyValue(line []byte, i int) (string, []byte, error) {
	// 存入值到 configKeyValue， 更新 fis
	index := bytes.Index(line, []byte(SEP))
	if index == -1 {
		return "", nil, fmt.Errorf("key %s error, not found =, line: %d", string(line), i)
	}
	// 左边的是key， 右边的是值
	key := bytes.Trim(line[:index], " ")
	// 不能包含. 和空格
	if bytes.Contains(key, []byte(" ")) {
		return "", nil, fmt.Errorf("key error, not allow contain space, key: %s, line: %d", string(key), i)
	}
	if bytes.Contains(key, []byte(".")) {
		return "", nil, fmt.Errorf("key error, not allow contain point, key: %s, line: %d", string(key), i)
	}
	// value 去掉2边空格
	value := bytes.Trim(line[index+1:], " ")
	k := string(key)

	// if string(module_name) != "" {
	// 	k = string(module_name) + "." + k
	// }
	if _, ok := kvconfig.KeyValue[k]; ok {
		// 去掉重复项
		return "", nil, fmt.Errorf("key duplicate, key: %s line: %d", k, i)
	}

	return k, value, nil
}
