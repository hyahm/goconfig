package goconfig

import (
	"strings"
)

// 修改配置文件的值

func UpdateValue(key string, value string) bool {
	// 修改位置文件的值， 如果修改成功返回true， 否则 false
	if _, ok := kvconfig.KeyValue[key]; !ok {
		// 不存在这个key， 就不修改
		return false
	}
	kvconfig.KeyValue[key] = value
	if strings.Contains(key, ".") {
		index := strings.LastIndex(key, ".")
		module := key[:index]
		key := key[index+1:]
		for i := range kvconfig.Groups {
			if kvconfig.Groups[i].name == module {
				for j := range kvconfig.Groups[i].group {

					if key == kvconfig.Groups[i].group[j].key {
						kvconfig.Groups[i].group[j].value = value
						FlushWrite()
						return true
					}
				}

			}
		}
	} else {
		for _, v := range kvconfig.Lines {
			if v.key == key {
				v.value = value
				FlushWrite()
				return true
			}
		}
	}
	return false
}
