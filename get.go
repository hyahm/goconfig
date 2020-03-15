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
	if Config == nil {
		panic("init first")
	}
	// key 不能包含多个.
	var this float64
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if len(value) > 0 {
		this = value[0]
	}

	if _, ok := Config.KeyValue[key]; !ok {
		str := strconv.FormatFloat(this, 'E', -1, 64)
		Config.KeyValue[key] = str
		return this
	}
	f64, err := strconv.ParseFloat(string(Config.KeyValue[key]), 64)
	if err != nil {
		return this
	}
	return f64
}

func ReadFile(key string, value ...string) string {
	if Config == nil {
		panic("init first")
	}
	var this string
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		Config.KeyValue[key] = this
	}
	// 读取文件
	bs, err := ioutil.ReadFile(string(Config.KeyValue[key]))
	if err != nil {
		return ""
	}
	return string(bs)
}

func ReadString(key string, value ...string) string {
	if Config == nil {
		panic("init first")
	}
	var this string
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		Config.KeyValue[key] = this
		return this
	}
	return string(Config.KeyValue[key])
}

// 返回int
func ReadInt(key string, value ...int) int {
	if Config == nil {
		panic("init first")
	}
	var this int
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		str := strconv.Itoa(this)
		Config.KeyValue[key] = str
		return this
	}
	i, err := strconv.Atoi(string(Config.KeyValue[key]))
	if err != nil {
		return this
	}
	return i
}

func ReadUint64(key string, value ...uint64) uint64 {
	if Config == nil {
		panic("init first")
	}
	var this uint64
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		str := strconv.FormatUint(this, 10)
		Config.KeyValue[key] = str
		return this
	}
	i, err := strconv.ParseUint(string(Config.KeyValue[key]), 10, 64)
	if err != nil {
		return this
	}
	return i
}

// 2边需要用到引号
func ReadPassword(key string, value ...string) string {
	if Config == nil {
		panic("init first")
	}
	var this string
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		Config.KeyValue[key] = fmt.Sprintf(`"%s"`, this)
		return this
	}
	v := Config.KeyValue[key]
	// 如果头尾不是"
	l := len(v)
	if string(v[0:1]) != "\"" || string(v[l-1:l]) != "\"" {
		return this
	}
	return string(v[1 : l-1])
}

func ReadBool(key string, value ...bool) bool {
	if Config == nil {
		panic("init first")
	}
	var this bool
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		if this {
			Config.KeyValue[key] = "true"
		} else {
			Config.KeyValue[key] = "false"
		}
		return this
	}
	if string(Config.KeyValue[key]) == "true" {
		return true
	} else if string(Config.KeyValue[key]) == "false" {
		return false
	} else {
		return this
	}
}

func ReadInt64(key string, value ...int64) int64 {
	if Config == nil {
		panic("init first")
	}
	var this int64
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		str := strconv.FormatInt(this, 10)
		Config.KeyValue[key] = str
		return this
	}
	i, err := strconv.ParseInt(string(Config.KeyValue[key]), 10, 64)
	if err != nil {
		return this
	}
	return i
}

func ReadBytes(key string, value ...[]byte) []byte {
	if Config == nil {
		panic("init first")
	}
	var this []byte
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if string(v) != string(this) {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		return this
	}
	return []byte(Config.KeyValue[key])
}

func ReadEnv(key string, value ...string) (s string) {
	s = os.Getenv(key)
	if s == "" && len(value) > 0 {
		s = value[0]
	}
	return
}

func ReadDuration(key string, value ...time.Duration) time.Duration {
	if Config == nil {
		panic("init first")
	}
	var this time.Duration
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		return this
	}
	i, err := strconv.ParseInt(string(Config.KeyValue[key]), 10, 64)
	if err != nil {
		return this
	}
	return time.Duration(i)
}

// 末尾不带/
func ReadWithoutEndSlash(key string, value ...string) string {
	if Config == nil {
		panic("init first")
	}
	var this string
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if this != "" {
		l := len(this)
		if this[l-1:l] == "/" {
			this = this[:l-1]
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		return this
	}
	this = string(Config.KeyValue[key])
	if this != "" {
		l := len(this)
		if this[l-1:l] == "/" {
			this = this[:l-1]
		}
	}
	return this
}

// 末尾带/
func ReadWithEndSlash(key string, value ...string) string {
	if Config == nil {
		panic("init first")
	}
	var this string
	for i, v := range value {
		// 最多3个， 超过了就使用默认值
		if i == Deep {
			break
		}
		if v != this {
			this = v
			break
		}
	}
	if this != "" {
		l := len(this)
		if this[l-1:l] != "/" {
			this = this + "/"
		}
	}
	if _, ok := Config.KeyValue[key]; !ok {
		return this
	}
	this = string(Config.KeyValue[key])
	if this != "" {
		l := len(this)
		if this[l-1:l] != "/" {
			this = this + "/"
		}
	}
	return this
}
