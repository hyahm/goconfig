package goconfig

import (
	"fmt"
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

	if _, ok := fl.KeyValue[key]; !ok {
		// 更新文件
		writeFile(subkey, value, module, notes...)
		// 更新map
		fl.KeyValue[key] = value
	}
}

func WriteFloat(key string, value float64, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	s1 := strconv.FormatFloat(value, 'E', -1, 64)
	update(key, s1, notes...)

}

func WriteString(key string, value string, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	update(key, value, notes...)

}

// 返回int
func WriteInt(key string, value int, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	s := strconv.Itoa(value)
	update(key, s, notes...)

}

func WriteUint64(key string, value uint64, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	s := strconv.FormatUint(value, 10)
	update(key, s, notes...)
}

// 2边需要用到引号
func WritePassword(key string, value string, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	s := fmt.Sprintf(`"%s"`, value)
	update(key, s, notes...)
}

func WriteBool(key string, value bool, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	s := "false"
	if value {
		s = "true"
	}
	update(key, s, notes...)
}

func WriteInt64(key string, value int64, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	s := strconv.FormatInt(value, 10)
	update(key, s, notes...)

}

func WriteBytes(key string, value []byte, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	update(key, string(value), notes...)

}

func WriteNotesForModule(name string, notes ...string) {
	if fl == nil {
		panic("init first")
	}
	for _, v := range fl.Groups {
		if string(v.name) == name {
			v.note = append(v.note, notes...)
			return
		}
	}
	fl.newGroup(name, notes...)
}
