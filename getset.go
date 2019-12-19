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

func GetSetFloat(key string, value float64) float64 {
	if configKeyValue == nil {
		panic("init first")
	}
	s1 := strconv.FormatFloat(value, 'E', -1, 64)
	update(key, s1)
	return value
}

func GetSetFile(key string, value string) string {
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

func GetSetString(key string, value string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	update(key, value)
	return value
}

// 返回int
func GetSetInt(key string, value int) int {
	if configKeyValue == nil {
		panic("init first")
	}
	s := strconv.Itoa(value)
	update(key, s)
	return value
}

func GetSetUint64(key string, value uint64) uint64 {
	if configKeyValue == nil {
		panic("init first")
	}
	s := strconv.FormatUint(value, 10)
	update(key, s)
	return value
}

// 2边需要用到引号
func GetSetPassword(key string, value string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	s := fmt.Sprintf(`"%s"`, value)
	update(key, s)
	return value
}

func GetSetBool(key string, value bool) bool {
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

func GetSetInt64(key string, value int64) int64 {
	if configKeyValue == nil {
		panic("init first")
	}
	s := strconv.FormatInt(value, 10)
	update(key, s)
	return value
}

func GetSetJson(key string, value interface{}) interface{} {
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

//func GetSetMap(key string) map[string]interface{} {
//	// value only three format
//	x := configKeyValue[key]
//	//x := `{"asdf":"ajsdfkl","type":6,"has":true}`
//
//	l := len(x)
//	kv := make(map[string]interface{}, 0)
//	if string(x[0]) == "{" && string(x[l-1]) == "}" {
//
//		// 去头和尾和空格
//		newstr := strings.Trim(x[1:l-1], " ")
//
//		// 逗号分开组
//		sl := strings.Split(newstr, ",")
//		if sl[0] == "" {
//			return kv
//		}
//		for _, v := range sl {
//			// 去头尾空格
//			var k string
//
//			mstr := strings.Trim(v, " ")
//			// 冒号分开,避免values存在:,以第一个冒号分割
//			index := strings.Index(mstr, ":")
//			//key是：,要去掉头尾空格
//			keyquote := strings.Trim(mstr[:index], " ")
//			//还要去掉2边的引号,如果没有冒号，格式不正确
//			kl := len(keyquote)
//			if string(keyquote[0]) == "\"" && string(keyquote[kl-1]) == "\"" && kl > 2 {
//				// key去掉2边的空格， 如果是空的，key值不能为空
//				k = strings.Trim(keyquote[1:kl-1], " ")
//
//				if k == "" {
//					panic("key 值不能为空")
//				}
//			} else {
//				panic("key 缺少双引号或者没有值")
//			}
//
//			// value是：去头尾空格
//			valuequote := strings.Trim(mstr[index+1:], " ")
//			//查看左右2边是否存在双引号
//			vl := len(valuequote)
//			if string(valuequote[0]) == "\"" && string(valuequote[vl-1]) == "\"" {
//				// 存在双引号，那就是字符串,获取里面的值，不去掉2边的空格
//				value := valuequote[1 : vl-1]
//				kv[k] = value
//				continue
//
//			} else {
//				// 否则是数字或者布尔值，先判断布尔值
//				if valuequote == "true" {
//					kv[k] = true
//					continue
//				} else if valuequote == "false" {
//					kv[k] = false
//					continue
//				} else if v, err := strconv.ParseInt(valuequote, 10, 64); err == nil {
//					//判断数字int64
//					kv[k] = v
//					continue
//				} else if v, err := strconv.ParseFloat(valuequote, 64); err == nil {
//					//判断数字float64
//					kv[k] = v
//					continue
//				} else {
//
//					panic("value 只支持string,int64,float64,bool")
//				}
//
//			}
//		}
//
//	} else {
//		panic("头尾缺少大括号")
//	}
//
//	return kv
//}
//
//func GetSetIntArray(key string) []int {
//	il := make([]int, 0)
//	vl := configKeyValue[key]
//	vlength := len(vl)
//	if vl[0:1] == "[" && vl[vlength-1:vlength] == "]" {
//		vlist := strings.Split(vl[1:vlength-1], ",")
//		//如果没值就返回空数组
//		if vlist[0] == "" {
//			return il
//		}
//		for _, v := range vlist {
//			//去掉2边的空格
//			i, err := strconv.Atoi(strings.Trim(v, " "))
//			if err != nil {
//				panic(fmt.Sprintf("key:%s,%v", key, err))
//			}
//			il = append(il, i)
//		}
//		return il
//	} else {
//		panic(fmt.Sprintf("key:%s,not an int array format \n", key))
//	}
//	return il
//}
//
//func GetSetStringArray(key string) []string {
//	sl := make([]string, 0)
//	vl := configKeyValue[key]
//	vlength := len(vl)
//	if vl[0:1] == "[" && vl[vlength-1:vlength] == "]" {
//		vlist := strings.Split(vl[1:vlength-1], ",")
//		//如果没值就返回空数组
//		if vlist[0] == "" {
//			return sl
//		}
//		for _, v := range vlist {
//			//去掉2边的空格
//			stringquote := strings.Trim(v, " ")
//			// 检查2边是否有双引号
//			ql := len(stringquote)
//			if stringquote[0:1] == "\"" && stringquote[ql-1:ql] == "\"" {
//				stringlist := stringquote[1 : ql-1]
//				if stringlist == "" {
//					continue
//				}
//				sl = append(sl, stringlist)
//			} else {
//				log.Fatal(fmt.Sprintf("key:%s,value must be has quote \n", key))
//				return sl
//			}
//
//		}
//		return sl
//	} else {
//		log.Fatalf("key:%s,not an int array format \n", key)
//	}
//	return sl
//}
