package goconfig

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
)

// 全部统一 utf-8格式
// 文件操作
// 用于写文件

type fileinfo struct {
	Key    []byte
	Value  []byte
	Data   []byte
	Module string
	Note   []byte
}

var fis []*fileinfo
var mu sync.RWMutex

func init() {
	mu = sync.RWMutex{}
}

func readlines() {
	fis = make([]*fileinfo, 0)

	buf := bufio.NewReader(file)
	module_name := ""
	line_num := 1
	for {
		line_byte, err := buf.ReadBytes('\n')
		if err != nil {
			if err != io.EOF {
				file.Close()
				panic(err)
			} else {
				format(line_byte, module_name, line_num)
				break
			}
		}
		format(line_byte, module_name, line_num)
		line_num++
	}

	getKey()
}

func format(line_byte []byte, module_name string, line_num int) {
	// 去掉windows换行的\r
	line_byte = bytes.ReplaceAll(line_byte, []byte("\r"), []byte(""))
	// 去掉换行 \n
	line_byte = bytes.ReplaceAll(line_byte, []byte("\n"), []byte(""))
	// 读取完成后会一次写入配置文件

	//去掉2边的空格
	line_byte_no_space := bytes.Trim(line_byte, " ")

	// 忽略注释和空行
	if string(line_byte_no_space) == "" {
		return
	}
	if string(line_byte_no_space[0:1]) == NOTE {
		// nothing to do
		// 注释
		fis = append(fis, &fileinfo{
			Data: []byte(""),
			Note: line_byte_no_space,
		})
		return
	}

	if string(line_byte_no_space[0:1]) == MODEL_START {
		// 模块
		line_lenth := len(line_byte_no_space)
		if string(line_byte_no_space[line_lenth-1:line_lenth]) == MODEL_END {
			// 模块
			module_name = strings.Trim(string(line_byte_no_space[1:line_lenth-1]), " ")
			//模块直接跳过
			return
		} else {
			file.Close()
			panic(fmt.Sprintf("格式错误， 行号：%d", line_num))
		}
	}

	// 如果是模块
	if module_name != "" {
		fis = append(fis, &fileinfo{
			Data:   line_byte_no_space,
			Module: module_name,
		})
	} else {
		fis = append(fis, &fileinfo{
			Data: line_byte_no_space,
		})
	}
}

// 写入到
func getKey() {
	// 存入值到 configKeyValue， 更新 fis
	tmp := make([]*fileinfo, 0)
	for i, v := range fis {
		//先找到第一个 SEP
		line := string(v.Data)
		index := strings.Index(line, SEP)
		if index == -1 {
			panic(fmt.Sprintf("key error, not found =, line: %d", i+1))
		}
		// 左边的是key， 右边的是值
		key := strings.Trim(line[:index], " ")
		// 不能包含. 和空格
		if strings.Contains(key, " ") {
			panic(fmt.Sprintf("key error, not allow contain space, key: %s", key))
		}
		if strings.Contains(key, ".") {
			panic(fmt.Sprintf("key error, not allow contain point, key: %s", key))
		}
		v.Key = []byte(key)
		value := strings.Trim(line[index+1:], " ")
		if v.Module != "" {
			key = v.Module + "." + key
		}
		if _, ok := configKeyValue[key]; ok {
			// 去掉重复项
			continue
		}

		configKeyValue[key] = value
		v.Value = []byte(value)
		tmp = append(tmp, v)
	}
	fis = nil
	fis = tmp
}

const BEGIN = 0
const END = 1

// 格式文件, sit 0, 1 ,
func writeFile(key, value, module string, notes ...string) {
	if len(fis) == 0 {
		fis = append(fis, &fileinfo{
			Key:    []byte(key),
			Value:  []byte(value),
			Data:   nil,
			Module: module,
		})

	} else {
		tmp := make([]*fileinfo, 0)
		module_read := module
		for i, v := range fis {
			if module == "" {
				//没有模块，插入到开头
				tmp = append(tmp, &fileinfo{
					Key:    []byte(key),
					Value:  []byte(value),
					Data:   nil,
					Module: "",
				})
				tmp = append(tmp, fis...)
				fis = nil
				fis = tmp
				break
			} else {
				if module_read == "" && v.Module == module {
					// 存在模块， 模块开始行
					module_read = module
					continue
				} else if module_read != "" && v.Module != module {
					// 说明已经改变了模块
					tmp = append(tmp, fis[:i]...)
					tmp = append(tmp, &fileinfo{
						Key:    []byte(key),
						Value:  []byte(value),
						Data:   nil,
						Module: module,
					})
					tmp = append(tmp, fis[i:]...)
					break
				} else {
					//如果读取到最后，还是没变， 末尾添加
					tmp = fis
					tmp = append(tmp, &fileinfo{
						Key:    []byte(key),
						Value:  []byte(value),
						Data:   nil,
						Module: module,
					})

				}

			}

		}
		fis = nil
		fis = tmp
	}
	// 组合文件内容

	var fileData []byte
	var module_write string // 保留模块名
	for i, v := range fis {
		if module_write != v.Module {
			module_write = v.Module
			// 插入模块, 模块开头空行
			if i != 0 {
				fileData = append(fileData, []byte("\n")...)
			}

			fileData = append(fileData, []byte(MODEL_START)...)
			fileData = append(fileData, []byte(module_write)...)
			fileData = append(fileData, []byte(MODEL_END)...)
			fileData = append(fileData, []byte("\n")...)
		}
		if string(v.Note) != "" {
			fileData = append(fileData, v.Note...)
			fileData = append(fileData, []byte("\n")...)
		} else {
			fileData = append(fileData, v.Key...)
			fileData = append(fileData, []byte(" ")...)
			fileData = append(fileData, []byte(SEP)...)
			fileData = append(fileData, []byte(" ")...)
			fileData = append(fileData, v.Value...)
			fileData = append(fileData, []byte("\n")...)
		}

	}
	if err := ioutil.WriteFile(configPath, fileData, 0644); err != nil {
		panic(err)
	}
}
