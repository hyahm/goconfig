package goconfig
//
//import (
//	"fmt"
//	"io/ioutil"
//)
//
//// 格式文件, sit 0, 1 ,
//func writeFile(key, value, module string, notes ...string) {
//	if len(fis) == 0 {
//		// 如果有注释， 先添加注释
//		for _,v := range notes {
//			fis = append(fis, &fileinfo{
//				Key:  []byte(key),
//				Note:  []byte(v),
//				Data:   nil,
//			})
//		}
//
//		// 如果文件是空的， 直接插入
//		fis = append(fis, &fileinfo{
//			Key:    []byte(key),
//			Value:  []byte(value),
//			Data:   nil,
//			Module: module,
//		})
//	} else {
//		// 不是第一行
//		tmp := make([]*fileinfo, 0)
//		for i, v := range fis {
//			if module == "" {
//				if i == 0 &&  v.Module != "" {
//					// 添加的不是模块， 第一行是模块的
//					for _,v := range notes {
//						tmp = append(tmp, &fileinfo{
//							Key:  []byte(key),
//							Note:  []byte(v),
//							Data:   nil,
//						})
//					}
//
//					// 添加key
//					tmp = append(tmp, &fileinfo{
//						Key:  []byte(key),
//						Value: []byte(value),
//						Data:   nil,
//						Module: "",
//					})
//					// 添加后面的
//					tmp = append(tmp, fis...)
//					fis = nil
//					fis = tmp
//					break
//				}
//				if v.Module != "" {
//					tmp = append(tmp, fis[:i]...)
//					// 如果有注释， 先添加注释
//					for _,v := range notes {
//						tmp = append(tmp, &fileinfo{
//							Key:  []byte(key),
//							Note:  []byte(v),
//							Data:   nil,
//						})
//					}
//
//					// 添加key
//					tmp = append(tmp, &fileinfo{
//						Key:  []byte(key),
//						Value: []byte(value),
//						Data:   nil,
//						Module: "",
//					})
//					// 添加后面的
//					tmp = append(tmp, fis[i:]...)
//					fis = nil
//					fis = tmp
//					break
//				} else	if i == len(fis) -1 && v.Module == "" {
//					// 最后一行插入
//
//					for _,v := range notes {
//						fis = append(fis,&fileinfo{
//							Key:  []byte(key),
//							Note:  []byte(v),
//							Data:   nil,
//						})
//					}
//
//					// 到最后一行都没找到就插入到末尾
//					fis = append(fis,&fileinfo{
//						Key:    []byte(key),
//						Value:  []byte(value),
//						Data:   nil,
//						Module: "",
//					})
//					break
//				}
//			} else {
//				// 插入模块
//				//第一行可以直接过
//				if i == 0 {
//					continue
//				}
//				if  fis[i-1].Module == module && v.Module != module {
//					// 上个是同一个模块， 这次
//					tmp = append(tmp, fis[:i]...)
//
//					//插入注释
//					for _,v := range notes {
//						tmp = append(tmp,&fileinfo{
//							Key:  []byte(key),
//							Note:  []byte(v),
//							Data:   nil,
//						})
//					}
//
//					tmp = append(tmp, &fileinfo{
//						Key:    []byte(key),
//						Value:  []byte(value),
//						Data:   nil,
//						Module: module,
//					})
//					tmp = append(tmp, fis[i:]...)
//					fis = nil
//					fis = tmp
//					break
//				} else if i == len(fis) -1  {
//					//如果读取到最后，还是没变， 末尾添加
//
//					//插入注释
//					for _,v := range notes {
//						fis = append(fis,&fileinfo{
//							Key:  []byte(key),
//							Note:  []byte(v),
//							Data:   nil,
//						})
//					}
//
//					fis = append(fis, &fileinfo{
//						Key:    []byte(key),
//						Value:  []byte(value),
//						Data:   nil,
//						Module: module,
//					})
//					break
//				}
//
//			}
//
//		}
//
//	}
//	// 组合文件内容
//	var fd []byte
//	var module_write string // 保留模块名
//	fmt.Println("----------------------------------")
//	for i, v := range fis {
//		fmt.Printf("module: %s, note: %s, key: %s, value: %s \n", v.Module, string(v.Note), string(v.Key), string(v.Value))
//		continue
//		// 有模块的
//		if v.Module == "" {
//			// 没有模块的，直接顺序添加注释和kv
//			if string(v.Note) != "" {
//				fd = insertNode(fd, v.Note)
//				continue
//			} else {
//				fd =insertKeyValue(fd,v.Key, v.Value)
//				continue
//			}
//		} else {
//			// 有模块的
//			if module_write != v.Module {
//
//				if v.Module != "" && string(v.Note) != "" {
//					// 如果带module的注释，先添加注释
//					fd = insertNode(fd, v.Note)
//				} else  {
//					// 如果不是注释
//					module_write = v.Module
//					// 如果是模块key的注释
//				}
//
//
//				// 插入注释
//				if string(v.Note) != "" {
//					fd = insertNode(fd, v.Note)
//					fd =insertModule(fd,v.Module)
//					continue
//				}
//				// 模块开头空行
//				if i != 0 {
//					fd =insertSpace(fd)
//				}
//				fd =insertModule(fd,v.Module)
//				fd =insertKeyValue(fd,v.Key, v.Value)
//				continue
//			}
//			// 插入注释
//			if string(v.Note) != "" {
//				fd = insertNode(fd, v.Note)
//				continue
//			}
//			fd =insertKeyValue(fd,v.Key, v.Value)
//		}
//
//
//	}
//	if err := ioutil.WriteFile(configPath, fd, 0644); err != nil {
//		panic(err)
//	}
//}
//
//
//func insertSpace(data []byte) (result []byte) {
//	// 插入空格
//	result = append(data, []byte("\n")...)
//	return result
//}
//
//func insertModule(data []byte,module string) (result []byte) {
//	// 插入模块
//	result = append(data, []byte(MODEL_START)...)
//	result = append(result, []byte(module)...)
//	result = append(result, []byte(MODEL_END)...)
//	result = append(result, []byte("\n")...)
//	return result
//}
//
//func insertNode(data []byte, note []byte)  (result []byte) {
//	// 插入注释
//	result = append(data, []byte(NOTE + " ")...)
//	result = append(result, note...)
//	result = append(result, []byte("\n")...)
//	return result
//}
//
//func insertKeyValue(data []byte, key , value []byte)  (result []byte) {
//	// 插入kv
//	result = append(data, key...)
//	result = append(result, []byte(" ")...)
//	result = append(result, []byte(SEP)...)
//	result = append(result, []byte(" ")...)
//	result = append(result, value...)
//	result = append(result, []byte("\n")...)
//	return result
//}