package goconfig

import (
	"fmt"
	"io/ioutil"
)

// 格式文件, sit 0, 1 ,
func writeFile(key, value, module string, notes ...[]byte) {
	// 判断是不是组
	if module == "" {
		//先添加key
		fl.newKeyValue(key, []byte(value), notes...)
	} else {
		// 组
		for i, g := range fl.Groups {
			if string(g.name) == module {
				fl.addGroupKeyValue(i, key, []byte(value), notes...)
				return
			}
		}
		// 不存在就新建
		fl.newGroupKeyValue([]byte(module), key, []byte(value), notes...)

	}
}

func FlushWrite() {

	for _, v := range fl.Lines {
		// 打印注释
		for _, n := range v.note {
			line := fmt.Sprintf("%s %s\n", NOTE, string(n))
			fl.Write = append(fl.Write, []byte(line)...)
		}
		// 打印kv
		kv := fmt.Sprintf("%s = %s\n", v.key, string(v.value))
		fl.Write = append(fl.Write, []byte(kv)...)
	}
	for _, v := range fl.Groups {
		// 模块添加换行
		fl.Write = append(fl.Write, []byte("\n")...)
		// 打印组注释
		for _, n := range v.note {
			line := fmt.Sprintf("%s %s\n", NOTE, string(n))
			fl.Write = append(fl.Write, []byte(line)...)
		}
		// 打印组
		g := fmt.Sprintf("%s%s%s\n", MODEL_START, string(v.name), MODEL_END)
		fl.Write = append(fl.Write, []byte(g)...)
		for _, gn := range v.group {
			// 组key 注释
			for _, nn := range gn.note {
				line := fmt.Sprintf("%s %s\n", NOTE, string(nn))
				fl.Write = append(fl.Write, []byte(line)...)
			}
			// 打印kv
			kv := fmt.Sprintf("%s = %s\n", gn.key, string(gn.value))
			fl.Write = append(fl.Write, []byte(kv)...)
		}
	}
	if err := ioutil.WriteFile(fl.Filepath, fl.Write, 0644); err != nil {
		panic(err)
	}
}
