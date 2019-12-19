package goconfig

import (
	"io/ioutil"
	"strconv"
)

func GetFloat(key string) float64 {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	// key 不能包含多个.
	if _, ok := ConfigKeyValue[key]; !ok {
		return 0
	}
	f64, err := strconv.ParseFloat(ConfigKeyValue[key], 64)
	if err != nil {
		return 0
	}
	return f64
}

func GetFile(key string) string {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return ""
	}
	// 读取文件
	bs, err := ioutil.ReadFile(ConfigKeyValue[key])
	if err != nil {
		return ""
	}
	return string(bs)
}

func GetString(key string) string {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return ""
	}
	return ConfigKeyValue[key]
}

// 返回int
func GetInt(key string) int {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return 0
	}
	i, err := strconv.Atoi(ConfigKeyValue[key])
	if err != nil {
		return 0
	}
	return i
}

func GetInt16(key string) int16 {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return 0
	}
	i := GetInt(key)
	// 如果大于取值区间，返回0
	if i > ((1<<16)/2)-1 || i < -((1<<16)/2) {
		return 0
	}
	return int16(GetInt(key))
}

func GetUint64(key string) uint64 {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return 0
	}
	i, err := strconv.ParseUint(ConfigKeyValue[key], 10, 64)
	if err != nil {
		return 0
	}
	return i
}

// 2边需要用到引号
func GetPassword(key string) string {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return ""
	}
	v := ConfigKeyValue[key]
	// 如果头尾不是"
	l := len(v)
	if string(v[0]) != "\"" || string(v[l-1:]) != "\"" {
		return ""
	}
	return v[1 : l-1]
}

func GetBool(key string) bool {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return false
	}
	if _, ok := ConfigKeyValue[key]; ok {
		if ConfigKeyValue[key] == "true" {
			return true
		} else {
			return false
		}
	}
	return false
}

func GetInt64(key string) int64 {
	if ConfigKeyValue == nil {
		panic("init first")
	}
	if _, ok := ConfigKeyValue[key]; !ok {
		return 0
	}
	i, err := strconv.ParseInt(ConfigKeyValue[key], 10, 64)
	if err != nil {
		return 0
	}
	return i
}

//func GetMap(key string) map[string]interface{} {
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
//		return nil
//	}
//
//	return kv
//}
//
//func GetIntArray(key string) []int {
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
//		log.Fatalf("key:%s,not an int array format \n", key)
//	}
//	return il
//}
//
//func GetStringArray(key string) []string {
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
//				panic(fmt.Sprintf("key:%s,value must be has quote \n", key))
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
