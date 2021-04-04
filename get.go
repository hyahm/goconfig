package goconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"time"
)

// 设置的默认值只会改变缓存， 不会改变文件
// 如果不存在key， 或者value 是错的才选用默认值

func ReadFloat64(key string, value ...float64) float64 {
	if kvconfig == nil {
		panic("init config file first")
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

	if _, ok := kvconfig.KeyValue[key]; !ok {
		str := strconv.FormatFloat(this, 'E', -1, 64)
		kvconfig.KeyValue[key] = str
		return this
	}
	f64, err := strconv.ParseFloat(kvconfig.KeyValue[key], 64)
	if err != nil {
		return this
	}
	return f64
}

func ReadFile(key string, value ...string) string {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		kvconfig.KeyValue[key] = this
	}
	// 读取文件
	bs, err := ioutil.ReadFile(kvconfig.KeyValue[key])
	if err != nil {
		return ""
	}
	return string(bs)
}

func ReadPath(key string, value ...string) string {
	val := ReadString(key, value...)
	return filepath.Clean(val)
}

func ReadString(key string, value ...string) string {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		kvconfig.KeyValue[key] = this
		return this
	}
	return kvconfig.KeyValue[key]
}

// 返回int
func ReadInt(key string, value ...int) int {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		str := strconv.Itoa(this)
		kvconfig.KeyValue[key] = str
		return this
	}
	i, err := strconv.Atoi(string(kvconfig.KeyValue[key]))
	if err != nil {
		return this
	}
	return i
}

func ReadUint64(key string, value ...uint64) uint64 {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		str := strconv.FormatUint(this, 10)
		kvconfig.KeyValue[key] = str
		return this
	}
	i, err := strconv.ParseUint(kvconfig.KeyValue[key], 10, 64)
	if err != nil {
		return this
	}
	return i
}

// 2边需要用到引号
func ReadPassword(key string, value ...string) string {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		kvconfig.KeyValue[key] = fmt.Sprintf(`"%s"`, this)
		return this
	}
	v := kvconfig.KeyValue[key]
	// 如果头尾不是"
	l := len(v)
	if string(v[0:1]) != "\"" || string(v[l-1:l]) != "\"" {
		return ReadString(key, this)
	}
	return string(v[1 : l-1])
}

func ReadBool(key string, value ...bool) bool {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		if this {
			kvconfig.KeyValue[key] = "true"
		} else {
			kvconfig.KeyValue[key] = "false"
		}
		return this
	}
	if string(kvconfig.KeyValue[key]) == "true" {
		return true
	} else if string(kvconfig.KeyValue[key]) == "false" {
		return false
	} else {
		return this
	}
}

func ReadInt64(key string, value ...int64) int64 {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		str := strconv.FormatInt(this, 10)
		kvconfig.KeyValue[key] = str
		return this
	}
	i, err := strconv.ParseInt(string(kvconfig.KeyValue[key]), 10, 64)
	if err != nil {
		return this
	}
	return i
}

func ReadBytes(key string, value ...[]byte) []byte {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		return this
	}
	return []byte(kvconfig.KeyValue[key])
}

func ReadEnv(key string, value ...string) (s string) {
	s = os.Getenv(key)
	if s == "" && len(value) > 0 {
		s = value[0]
	}
	return
}

func ReadDuration(key string, value ...time.Duration) time.Duration {
	if kvconfig == nil {
		panic("init config file first")
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

	if _, ok := kvconfig.KeyValue[key]; !ok {
		return this
	}
	this, _ = dr(kvconfig.KeyValue[key]).Duration()

	return this
}

// 末尾不带/
func ReadWithoutEndSlash(key string, value ...string) string {
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		return this
	}
	this = kvconfig.KeyValue[key]
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
	if kvconfig == nil {
		panic("init config file first")
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
	if _, ok := kvconfig.KeyValue[key]; !ok {
		return this
	}
	this = kvconfig.KeyValue[key]
	if this != "" {
		l := len(this)
		if this[l-1:l] != "/" {
			this = this + "/"
		}
	}
	return this
}

func ReadStructFromNode(key string, value interface{}) error {
	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		return errors.New("use ReadArrayBytes Function, the value must be a point")
	}
	arrayValue := reflect.TypeOf(value).Elem()

	if arrayValue.Kind() != reflect.Struct {
		return errors.New("not a struct point ")
	}

	keys := make(map[string]*asset, 0)
	for i := 0; i < arrayValue.NumField(); i++ {
		// 这个key 就是ini中的key
		key := strings.Split(arrayValue.Field(i).Tag.Get("json"), ",")[0]
		a := &asset{
			index: i,
			kind:  arrayValue.Field(i).Type.Kind(),
			root:  arrayValue,
		}
		keys[key] = a
	}
	if len(keys) == 0 {
		return errors.New("use ReadArrayBytes Function, the array value must be a struct point or struct ")
	}
	// 得到一定是个数组类似： [ {"name": "a"}, {"name":"b"}]
	if kvconfig == nil {
		panic("init config file first")
	}

	return walkNode(key, keys, value)
}

type asset struct {
	index int
	kind  reflect.Kind
	root  interface{}
}

func ReadArrayFromNode(key string, value interface{}) error {
	if reflect.TypeOf(value).Kind() != reflect.Ptr {
		return errors.New("use ReadArrayBytes Function, the value must be a point")
	}
	arrayValue := reflect.TypeOf(value).Elem()

	if arrayValue.Kind() != reflect.Slice {
		return errors.New("use ReadArrayBytes Function, the value must be a array")
	}
	var s reflect.Type
	if arrayValue.Elem().Kind() == reflect.Ptr {
		s = arrayValue.Elem().Elem()
	} else if arrayValue.Elem().Kind() == reflect.Struct {
		s = arrayValue.Elem()
	} else {
		return errors.New("use ReadArrayBytes Function, the array value must be a struct point or struct ")
	}
	keys := make(map[string]*asset, 0)
	for i := 0; i < s.NumField(); i++ {
		// 这个key 就是ini中的key
		key := strings.Split(s.Field(i).Tag.Get("json"), ",")[0]
		a := &asset{
			index: i,
			kind:  s.Field(i).Type.Kind(),
			root:  reflect.ValueOf(value).Field(i).Interface,
		}
		keys[key] = a
		fmt.Println(key, s.Field(i).Type.Kind())
	}
	if len(keys) == 0 {
		return errors.New("use ReadArrayBytes Function, the array value must be a struct point or struct ")
	}
	// 得到一定是个数组类似： [ {"name": "a"}, {"name":"b"}]
	if kvconfig == nil {
		panic("init config file first")
	}

	return walkNode(key, keys, value)
}

func walkNode(key string, keys map[string]*asset, value interface{}) error {
	l := len(key)
	// 获取key的前面都等于key.的key, 并且获取后缀
	// eg:   u5.aaa.bbb.ccc   key=u5 key. =u5.  后缀: aaa.bbb.ccc
	// key和 keys 中间一级是array的分隔符
	// 最后拼成的字符串应该是这个 [{"xxxx":"yyyy", "aaaa": "name"},{"xxxx":"zzzz", "aaaa": "nnnn"}]
	fmt.Println("---------------")
	// 先获取所有数组的key
	skey := make(map[string]int)
	for k := range kvconfig.KeyValue {
		if k[:l+1] == key+"." {
			endkey := k[l+1:]
			index := strings.Index(endkey, ".")
			gk := endkey[:index]
			if _, ok := skey[gk]; !ok {
				skey[gk] = 0
			}

		}
	}
	fmt.Println(skey)
	jsonStr := "["
	// 根据key 获取值拼接字符串
	for k := range skey {
		jsonStr += "{"
		for lk, ass := range keys {
			nodeKey := fmt.Sprintf("%s.%s.%s", key, k, lk)
			switch ass.kind {
			case reflect.String:
				jsonStr += fmt.Sprintf(`"%s":"%s",`, lk, ReadString(nodeKey))
			case reflect.Int:
				jsonStr += fmt.Sprintf(`"%s":%d,`, lk, ReadInt(nodeKey))
			case reflect.Bool:
				jsonStr += fmt.Sprintf(`"%s":%t,`, lk, ReadBool(nodeKey))
			case reflect.Int64:
				jsonStr += fmt.Sprintf(`"%s":%d,`, lk, ReadInt64(nodeKey))
			case reflect.Ptr:
				if reflect.ValueOf(value).CanSet() {
					ReadStructFromNode(nodeKey, ass.root)
				}
			case reflect.Slice:
				if ass.kind == reflect.Ptr {
					if reflect.ValueOf(value).CanSet() {
						if err := ReadArrayFromNode(nodeKey, &ass.root); err != nil {
							return err
						}
					}
				} else if ass.kind == reflect.Struct {
					if reflect.ValueOf(value).CanSet() {
						if err := ReadStructFromNode(nodeKey, &ass.root); err != nil {
							return err
						}
					}
				}
			case reflect.Struct:
				if reflect.ValueOf(value).CanSet() {
					if err := ReadStructFromNode(nodeKey, &ass.root); err != nil {
						return err
					}
				}
			default:
			}

		}

		if fl := len(jsonStr); jsonStr[fl-1:] == "," {
			jsonStr = jsonStr[:len(jsonStr)-1]
		}

		jsonStr += "},"
	}
	if fl := len(jsonStr); jsonStr[fl-1:] == "," {
		jsonStr = jsonStr[:len(jsonStr)-1]
	}

	jsonStr += "]"
	fmt.Println(jsonStr)
	return json.Unmarshal([]byte(jsonStr), value)
}
