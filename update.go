package goconfig

import (
	"fmt"
	"strings"
)

// 修改配置文件的值

func UpdateValue(key string, value string, notes ...string) bool {
	// 修改位置文件的值， 如果修改成功返回true， 否则 false

	defer FlushWrite()

	kvconfig.KeyValue[key] = value

	if strings.Contains(key, ".") {
		// 判断是否存在组
		haveGroup := false
		// 组索引
		groupIndex := 0
		// 组成员
		index := strings.LastIndex(key, ".")
		module := key[:index]
		key := key[index+1:]
		for i := range kvconfig.Groups {
			if kvconfig.Groups[i].name == module {
				haveGroup = true
				groupIndex = i
				for j := range kvconfig.Groups[i].group {

					if key == kvconfig.Groups[i].group[j].key {
						kvconfig.Groups[i].group[j].value = value
						kvconfig.Groups[i].group[j].note = append(kvconfig.Groups[i].group[j].note, notes...)
						return true
					}
				}

			}
		}
		if haveGroup {
			// 存在组， 就添加行就可以了
			kvconfig.addGroupKeyValue(groupIndex, key, value, notes...)
		} else {
			// 不存在组， 添加新的
			kvconfig.newGroupKeyValue(module, key, value, notes...)
		}
	} else {
		// 普通成员
		for i := range kvconfig.Lines {
			if kvconfig.Lines[i].key == key {
				kvconfig.Lines[i].value = value
				kvconfig.Lines[i].note = append(kvconfig.Lines[i].note, notes...)
				return true
			}
		}
		// 不存在的话
		kvconfig.newKeyValue(key, value, notes...)
		fmt.Println(kvconfig.Lines)

	}
	return false

}
