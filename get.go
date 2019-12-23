package goconfig

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// 设置的默认值只会改变缓存， 不会改变文件
// 如果不存在key， 或者value 是错的才选用默认值

func ReadFloat64(key string, value ...float64) float64 {
	if fl == nil {
		panic("init first")
	}
	// key 不能包含多个.
	var this float64
	if len(value) > 0 {
		this = value[0]
	}

	if _, ok := fl.KeyValue[key]; !ok {
		str := strconv.FormatFloat(this, 'E', -1, 64)
		fl.KeyValue[key] = []byte(str)
		return this
	}
	f64, err := strconv.ParseFloat(string(fl.KeyValue[key]), 64)
	if err != nil {
		return this
	}
	return f64
}

func ReadFile(key string, value ...string) string {
	if fl == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		fl.KeyValue[key] = []byte(this)
	}
	// 读取文件
	bs, err := ioutil.ReadFile(string(fl.KeyValue[key]))
	if err != nil {
		return ""
	}
	return string(bs)
}

func ReadString(key string, value ...string) string {
	if fl == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		fl.KeyValue[key] = []byte(this)
		return this
	}
	return string(fl.KeyValue[key])
}

// 返回int
func ReadInt(key string, value ...int) int {
	if fl == nil {
		panic("init first")
	}
	var this int
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		str := strconv.Itoa(this)
		fl.KeyValue[key] = []byte(str)
		return this
	}
	i, err := strconv.Atoi(string(fl.KeyValue[key]))
	if err != nil {
		return this
	}
	return i
}

func ReadUint64(key string, value ...uint64) uint64 {
	if fl == nil {
		panic("init first")
	}
	var this uint64
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		str := strconv.FormatUint(this, 10)
		fl.KeyValue[key] = []byte(str)
		return this
	}
	i, err := strconv.ParseUint(string(fl.KeyValue[key]), 10, 64)
	if err != nil {
		return this
	}
	return i
}

// 2边需要用到引号
func ReadPassword(key string, value ...string) string {
	if fl == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		fl.KeyValue[key] = []byte(fmt.Sprintf(`"%s"`, this))
		return this
	}
	v := fl.KeyValue[key]
	// 如果头尾不是"
	l := len(v)
	if string(v[0:1]) != "\"" || string(v[l-1:l]) != "\"" {
		return this
	}
	return string(v[1 : l-1])
}

func ReadBool(key string, value ...bool) bool {
	if fl == nil {
		panic("init first")
	}
	var this bool
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		if this {
			fl.KeyValue[key] = []byte("true")
		} else {
			fl.KeyValue[key] = []byte("false")
		}
		return this
	}
	if string(fl.KeyValue[key]) == "true" {
		return true
	} else if string(fl.KeyValue[key]) == "false" {
		return false
	} else {
		return this
	}
}

func ReadInt64(key string, value ...int64) int64 {
	if fl == nil {
		panic("init first")
	}
	var this int64
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		str := strconv.FormatInt(this, 10)
		fl.KeyValue[key] = []byte(str)
		return this
	}
	i, err := strconv.ParseInt(string(fl.KeyValue[key]), 10, 64)
	if err != nil {
		return this
	}
	return i
}

func ReadBytes(key string, value ...[]byte) []byte {
	if fl == nil {
		panic("init first")
	}
	var this []byte
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		return this
	}
	return fl.KeyValue[key]
}

func ReadEnv(key string, value ...string) (s string) {
	s = os.Getenv(key)
	if s == "" && len(value) > 0 {
		s = value[0]
	}
	return
}

func ReadDuration(key string, value ...time.Duration) time.Duration {
	if fl == nil {
		panic("init first")
	}
	var this time.Duration
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := fl.KeyValue[key]; !ok {
		return this
	}
	i, err := strconv.Atoi(string(fl.KeyValue[key]))
	if err != nil {
		return this
	}
	return time.Duration(i)
}

// 末尾不带/
func ReadPath(key string, value ...string) string {
	if fl == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		l := len(value[0])
		if value[0][l-1:l] == "/" {
			this = value[0][:l-1]
		} else {
			this = value[0]
		}
	}
	if _, ok := fl.KeyValue[key]; !ok {
		return this
	}
	this = string(fl.KeyValue[key])
	l := len(this)
	if this[l:l] == "/" {
		this = this[:l-1]
	}
	return this
}

// 末尾带/
func ReadUrl(key string, value ...string) string {
	if fl == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		l := len(value[0])
		if value[0][l-1:l] != "/" {
			this = value[0] + "/"
		} else {
			this = value[0]
		}
	}
	if _, ok := fl.KeyValue[key]; !ok {
		return this
	}
	this = string(fl.KeyValue[key])
	l := len(this)
	if this[l:l] != "/" {
		this = this + "/"
	}
	return this
}
