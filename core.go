package goconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"reflect"
)

// 通过key获取json数据的核心文件

type jsonStore struct {
	node string // 当前节点的key
	v    []byte
	m    map[string]*jsonStore
	s    []*jsonStore // 不知道需不需要排序，暂时不排序

}

var store *jsonStore

func init() {
	store = &jsonStore{}
}

func Start() (err error) {

	store.v, err = ioutil.ReadFile("aaa.json")
	if err != nil {
		return
	}
	// 判断数据是slice 还是 struct
	var jsonConfig interface{}
	err = json.Unmarshal(store.v, &jsonConfig)
	if err != nil {
		return
	}
	// store := &jsonStore{}
	jt := reflect.TypeOf(jsonConfig)
	jv := reflect.ValueOf(jsonConfig)
	switch jt.Kind() {
	case reflect.Map:
		fmt.Println(jv)
		store.m = make(map[string]*jsonStore)
		jr := jv.MapRange()
		for jr.Next() {
			// var vb []byte
			// vb, err = json.Marshal(jr.Value().Interface)
			// if err != nil {
			// 	return
			// }
			// store.m[jr.Key().String()] = vb
		}
		for i := 0; i < jt.NumField(); i++ {
			// store.m[]
			// fmt.Println(jv.Field(i))
			fmt.Println(jt.Field(i))
		}
	case reflect.Slice:
		store.s = make([]*jsonStore, 0)
		for i := 0; i < jv.Len(); i++ {

			// 判断数据类型， 只有整型字符串
			store.s = append(store.s,
				sliceHandle(jv.Index(i).Interface()))

			// fj := &jsonStore{}
		}
	default:
		return errors.New("not support")
	}
	if reflect.TypeOf(jsonConfig).Kind() == reflect.Map {

	}
	fmt.Println(reflect.TypeOf(jsonConfig).Kind())
	return
}

func sliceHandle(value interface{}) *jsonStore {
	js := &jsonStore{}
	st := reflect.TypeOf(value)
	switch st.Kind() {
	// 数组里面可以是map， slice， string， int， bool, float
	case reflect.Bool:
	case reflect.Map:
	}
	return js
}

func (js *jsonStore) mapHandle(value interface{}) {
	st := reflect.TypeOf(value)
	switch st.Kind() {
	// 数组里面可以是map， slice， string， int， bool, float
	case reflect.Bool:
	case reflect.Map:
	}
}
