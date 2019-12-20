package goconfig

import (
	"fmt"
	"io/ioutil"
	"strconv"
)

// 设置的默认值只会改变缓存， 不会改变文件
// 如果不存在key， 或者value 是错的才选用默认值

func ReadFloat64(key string, value ...float64) float64 {
	if configKeyValue == nil {
		panic("init first")
	}
	// key 不能包含多个.
	var this float64
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		str := strconv.FormatFloat(this, 'E', -1, 64)
		configKeyValue[key] = str
		return this
	}
	f64, err := strconv.ParseFloat(configKeyValue[key], 64)
	if err != nil {
		return this
	}
	return f64
}

func ReadFile(key string, value ...string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		configKeyValue[key] = this
	}
	// 读取文件
	bs, err := ioutil.ReadFile(configKeyValue[key])
	if err != nil {
		return ""
	}
	return string(bs)
}

func ReadString(key string, value ...string) string {
	if configKeyValue == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		configKeyValue[key] = this
		return this
	}
	return configKeyValue[key]
}

// 返回int
func ReadInt(key string, value ...int) int {
	if configKeyValue == nil {
		panic("init first")
	}
	var this int
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		str := strconv.Itoa(this)
		configKeyValue[key] = str
		return this
	}
	i, err := strconv.Atoi(configKeyValue[key])
	if err != nil {
		return this
	}
	return i
}

func ReadUint64(key string, value ...uint64) uint64 {
	if configKeyValue == nil {
		panic("init first")
	}
	var this uint64
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		str := strconv.FormatUint(this, 10)
		configKeyValue[key] = str
		return this
	}
	i, err := strconv.ParseUint(configKeyValue[key], 10, 64)
	if err != nil {
		return this
	}
	return i
}

// 2边需要用到引号
func ReadPassword(key string, value ...string) string {
	fmt.Println(configKeyValue)
	if configKeyValue == nil {
		panic("init first")
	}
	var this string
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		configKeyValue[key] = fmt.Sprintf(`"%s"`, this)
		return this
	}
	v := configKeyValue[key]
	fmt.Println(v)
	// 如果头尾不是"
	l := len(v)
	if v[0:1] != "\"" || v[l-1:l] != "\"" {
		return this
	}
	fmt.Println(v[1 : l-1])
	return v[1 : l-1]
}

func ReadBool(key string, value ...bool) bool {
	if configKeyValue == nil {
		panic("init first")
	}
	var this bool
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		if this {
			configKeyValue[key] = "true"
		} else {
			configKeyValue[key] = "false"
		}
		return this
	}
	if configKeyValue[key] == "true" {
		return true
	} else if configKeyValue[key] == "false" {
		return false
	} else {
		return this
	}
}

func ReadInt64(key string, value ...int64) int64 {
	if configKeyValue == nil {
		panic("init first")
	}
	var this int64
	if len(value) > 0 {
		this = value[0]
	}
	if _, ok := configKeyValue[key]; !ok {
		str := strconv.FormatInt(this, 10)
		configKeyValue[key] = str
		return this
	}
	i, err := strconv.ParseInt(configKeyValue[key], 10, 64)
	if err != nil {
		return this
	}
	return i
}

func ReadBytes(key string, value ...[]byte) []byte {
	if configKeyValue == nil {
		panic("init first")
	}
	if _, ok := configKeyValue[key]; !ok {
		return nil
	}
	//i, err := strconv.ParseInt(configKeyValue[key], 10, 64)
	//if err != nil {
	//	return nil
	//}
	return value[0]
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
