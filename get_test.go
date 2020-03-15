package goconfig

import (
	"testing"
	"time"
)

type getvalue struct {
	title       string
	key         string
	value       interface{}
	shouldMatch interface{}
}

var x = `
uint64 =9223372036854775808
int64=123456223
int=1234556223
dd=2

[one]
float=0.25

[two]
bool=true
password= "adlfjlskdf "
`

func TestINI(t *testing.T) {
	// 方便测试，
	InitFromBytes([]byte(x), INI)
	tests := []getvalue{
		{
			title: "float64",
			key:   "one.float",
			value: 0.25,
		},
		{
			title: "uint64",
			key:   "uint64",
			value: uint64(1 << 63),
		},
		{
			title: "int64",
			key:   "int64",
			value: int64(123456223),
		},
		{
			title: "int",
			key:   "int",
			value: 1234556223,
		},
		{
			title: "bool",
			key:   "two.bool",
			value: true,
		},
		{
			title: "Duration",
			key:   "dd",
			value: 2 * time.Second,
		},
		{
			title: "password",
			key:   "two.password",
			value: "adlfjlskdf ",
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			testGet(t, test)
		})
	}
}

var y = `
{
	"uint64": 9223372036854775808,
	"int64": 123456223,
	"int": 1234556223,
	"dd": 2,
	"one": {
		"float": 0.25
	},
	"two": {
		"bool": true,
		"password": "adlfjlskdf "
	}

}

`

func TestJSON(t *testing.T) {
	// 方便测试，
	InitFromBytes([]byte(y), JSON)
	tests := []getvalue{
		{
			title: "float64",
			key:   "one.float",
			value: 0.25,
		},
		{
			title: "uint64",
			key:   "uint64",
			value: uint64(1 << 63),
		},
		{
			title: "int64",
			key:   "int64",
			value: int64(123456223),
		},
		{
			title: "int",
			key:   "int",
			value: 1234556223,
		},
		{
			title: "bool",
			key:   "two.bool",
			value: true,
		},
		{
			title: "Duration",
			key:   "dd",
			value: 2 * time.Second,
		},
		{
			title: "password",
			key:   "two.password",
			value: "adlfjlskdf ",
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			testGet(t, test)
		})
	}
}

var z = `
uint64: 9223372036854775808
int64: 123456223
int: 1234556223
dd: 2

one:
  float: 0.25
two:
  bool: true
  password: "adlfjlskdf "
`

func TestYAML(t *testing.T) {
	// 方便测试，
	InitFromBytes([]byte(z), YAML)
	tests := []getvalue{
		{
			title: "float64",
			key:   "one.float",
			value: 0.25,
		},
		{
			title: "uint64",
			key:   "uint64",
			value: uint64(1 << 63),
		},
		{
			title: "int64",
			key:   "int64",
			value: int64(123456223),
		},
		{
			title: "int",
			key:   "int",
			value: 1234556223,
		},
		{
			title: "bool",
			key:   "two.bool",
			value: true,
		},
		{
			title: "Duration",
			key:   "dd",
			value: 2 * time.Second,
		},
		{
			title: "password",
			key:   "two.password",
			value: "adlfjlskdf ",
		},
	}
	for _, test := range tests {
		t.Run(test.title, func(t *testing.T) {
			testGet(t, test)
		})
	}
}

func testGet(t *testing.T, test getvalue) {
	switch test.value.(type) {
	case string:
		t.Log(test.title)
		if test.key == "string" {
			if test.value != ReadString(test.key) {
				t.Error("fail string")
			}
		}
		if test.key == "file" {
			if test.value != ReadFile(test.key) {
				t.Error("fail file")
			}
		}
		if test.key == "password" {
			if test.value != ReadPassword(test.key) {
				t.Logf("--%s---", ReadPassword(test.key))
				t.Error("fail password")
			}
		}
	case int:
		t.Log(test.title)
		if test.value != ReadInt(test.key) {
			t.Error("fail int")
		}
	case uint64:
		t.Log(test.title)
		if test.value != ReadUint64(test.key) {
			t.Error("fail uint64")
		}
	case int64:
		t.Log(test.title)
		if test.value != ReadInt64(test.key) {
			t.Error("fail int64")
		}
	case float64:
		t.Log(test.title)
		t.Log("test value: ", ReadFloat64(test.key))
		t.Log("value: ", test.value)
		if test.value != ReadFloat64(test.key) {
			t.Error("fail float64")
		}
	case time.Duration:
		t.Log(test.title)
		if test.value != ReadDuration(test.key)*time.Second {
			t.Error("fail")
		}
	default:

	}
}
