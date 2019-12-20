package goconfig

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func update(key, value string, notes ...string) {
	module := ""
	subkey := ""
	// 是否包含空格
	if strings.Contains(key, " ") {
		panic("key error , not allow contain space")
	}
	c := strings.Count(key, ".")
	if c >= 2 {
		panic("key error , not allow contain point more than one ")
	}
	if c == 1 {
		kv := strings.Split(key, ".")
		module = kv[0]
		subkey = kv[1]
	} else {
		subkey = key
	}

	if _, ok := configKeyValue[key]; !ok {
		// 更新map
		configKeyValue[key] = value
		// 更新文件
		writeFile(subkey, value, module, notes...)
	}
}

func WriteFloat(key string, value float64, notes ...string) float64 {
	if configKeyValue == nil {
		panic("init first")
	}
	s1 := strconv.FormatFloat(value, 'E', -1, 64)
	update(key, s1, notes...)
	return value
}

func WriteFile(key string, value string, notes ...string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	// 读取文件
	update(key, value)
	bs, err := ioutil.ReadFile(value)
	if err != nil {
		return ""
	}
	return string(bs)
}

func WriteString(key string, value string, notes ...string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	update(key, value)
	return value
}

// 返回int
func WriteInt(key string, value int, notes ...string) int {
	if configKeyValue == nil {
		panic("init first")
	}
	s := strconv.Itoa(value)
	update(key, s)
	return value
}

func WriteUint64(key string, value uint64, notes ...string) uint64 {
	if configKeyValue == nil {
		panic("init first")
	}
	s := strconv.FormatUint(value, 10)
	update(key, s)
	return value
}

// 2边需要用到引号
func WritePassword(key string, value string, notes ...string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	s := fmt.Sprintf(`"%s"`, value)
	update(key, s)
	return value
}

func WriteBool(key string, value bool, notes ...string) bool {
	if configKeyValue == nil {
		panic("init first")
	}
	s := "false"
	if value {
		s = "true"
	}
	update(key, s)
	return value
}

func WriteInt64(key string, value int64, notes ...string) int64 {
	if configKeyValue == nil {
		panic("init first")
	}
	s := strconv.FormatInt(value, 10)
	update(key, s)
	return value
}

func WriteJson(key string, value interface{}, notes ...string) interface{} {
	if configKeyValue == nil {
		panic("init first")
	}
	s, err := json.Marshal(value)
	if err != nil {
		panic(err)
	}
	update(key, string(s))
	return value
}
