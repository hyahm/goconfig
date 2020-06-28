package goconfig

import (
	"encoding/json"
	"errors"
	"reflect"
	"strconv"
	"strings"
)

// json 只有2种通用结构

// map[string]interface{}  和  []interface{}
func (c *config) readJson() error {

	err := json.Unmarshal(c.Read, &c.sjson)
	if err != nil {
		return err
	}
	if err := c.parseJson("", c.sjson); err != nil {
		return err
	}

	return nil
}

func (c *config) parseJson(module string, value map[string]interface{}) error {
	for k, v := range value {
		// refType := reflect.TypeOf(v)
		refValue := reflect.ValueOf(v)
		switch refValue.Kind() {
		case reflect.Float64, reflect.Float32:
			// 如果是整数， 去掉后面的
			f := strconv.FormatFloat(refValue.Float(), 'f', 6, 64)
			fl := strings.Split(f, ".")
			if fl[1] == "000000" {
				f = fl[0]
			}

			if module == "" {
				c.KeyValue[k] = f
			} else {
				c.KeyValue[module+"."+k] = f
			}
		case reflect.Bool:
			//
			bs := "false"
			b := refValue.Bool()
			if b {
				bs = "true"
			}
			if module == "" {
				c.KeyValue[k] = bs
			} else {
				c.KeyValue[module+"."+k] = bs
			}

		case reflect.Int64, reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			//
			if module == "" {
				c.KeyValue[k] = strconv.FormatInt(refValue.Int(), 10)
			} else {
				c.KeyValue[module+"."+k] = strconv.FormatInt(refValue.Int(), 10)
			}

		case reflect.Uint64, reflect.Uint16, reflect.Uint32, reflect.Uint8:
			//

			if module == "" {
				c.KeyValue[k] = strconv.FormatUint(refValue.Uint(), 10)
			} else {
				c.KeyValue[module+"."+k] = strconv.FormatUint(refValue.Uint(), 10)
			}

		case reflect.Map:
			//
			var next map[string]interface{}
			n, _ := json.Marshal(v)
			err := json.Unmarshal(n, &next)
			if err != nil {
				return err
			}
			if module == "" {
				if err := c.parseJson(k, next); err != nil {
					return err
				}
			} else {
				if err := c.parseJson(module+"."+k, next); err != nil {
					return err
				}
			}

		case reflect.Slice:
			//
			b, _ := json.Marshal(refValue.Interface())
			if module == "" {
				c.KeyValue[k] = string(b)
			} else {
				c.KeyValue[module+"."+k] = string(b)
			}

		case reflect.String:
			if module == "" {
				c.KeyValue[k] = refValue.String()
			} else {
				c.KeyValue[module+"."+k] = refValue.String()
			}
		default:
			return errors.New("not a valid config file")
		}

	}
	return nil
}
